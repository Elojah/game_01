package main

import (
	"sync/atomic"
	"time"

	"github.com/nats-io/go-nats"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

var (
	maxTime = time.Unix(1<<63-62135596801, 999999999)
)

type tick chan time.Time

// Sequencer is an ordering/event extractor layer between two consumers.
type Sequencer struct {
	id game.ID
	game.EventService

	logger zerolog.Logger

	input   tick
	output  chan game.Event
	fetcher tick

	interrupt int32
}

// Close kills both fetch/input goroutines.
func (s *Sequencer) Close() {
	close(s.fetcher)
	close(s.output)
}

// NewSequencer returns a new sequencer with two listening goroutines to fetch/order events.
func NewSequencer(id game.ID, es game.EventService, callback func(game.Event)) *Sequencer {
	s := Sequencer{
		id:           id,
		logger:       log.With().Str("seq", id.String()).Logger(),
		EventService: es,
		fetcher:      make(tick, 0),
		input:        make(tick, 0),
		output:       make(chan game.Event, 0),
	}
	go func() {
		for {
			select {
			case t, ok := <-s.fetcher:
				if !ok {
					return
				}
				events, err := s.ListEvent(game.EventBuilder{
					Key: s.id.String(),
					Min: int(t.UnixNano()),
				})
				if err != nil {
					s.logger.Error().Err(err).Msg("failed to fetch events")
					break
				}
				for _, event := range events {
					if atomic.CompareAndSwapInt32(&s.interrupt, 1, 0) {
						break
					}
					s.output <- event
				}
			}
		}
	}()
	go func() {
		var current time.Time
		waiting := maxTime
		for {
			select {
			case t := <-s.input:
				if t.Before(current) {
					atomic.CompareAndSwapInt32(&s.interrupt, 0, 1)
				}
				waiting = t
				s.fetcher <- t
			case event, ok := <-s.output:
				if !ok {
					return
				}
				if event.TS.After(waiting) {
					atomic.CompareAndSwapInt32(&s.interrupt, 0, 1)
					waiting = maxTime
					break
				}
				current = event.TS
				callback(event)
			}
		}
	}()
	return &s
}

// MsgHandler is the consumer function to subscribe for event ordering.
func (s Sequencer) MsgHandler(msg *nats.Msg) {
	var eventS storage.Event
	if _, err := eventS.Unmarshal(msg.Data); err != nil {
		s.logger.Error().Err(err).Msg("error unmarshaling event")
		return
	}
	event := eventS.Domain()
	if err := s.CreateEvent(event, s.id); err != nil {
		s.logger.Error().Err(err).Msg("error creating event")
		return
	}
	s.logger.Info().Str("event", event.ID.String()).Msg("event received")
	go func(ts time.Time) { s.input <- ts }(event.TS)
}
