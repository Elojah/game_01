package main

import (
	"github.com/nats-io/go-nats"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

type tick chan int64

// Sequencer is an ordering/event extractor layer between two consumers.
type Sequencer struct {
	id game.ID
	game.EventService

	logger zerolog.Logger

	input   tick
	fetch   tick
	process chan game.Event

	min       tick
	last      tick
	interrupt chan struct{}
}

// Close kills both fetch/input goroutines.
func (s *Sequencer) Close() {
	s.logger.Info().Msg("close sequencer")
	close(s.input)
	close(s.fetch)
	close(s.process)
}

// NewSequencer returns a new sequencer with two listening goroutines to fetch/order events.
func NewSequencer(id game.ID, limit int, es game.EventService, callback func(game.Event)) *Sequencer {
	s := Sequencer{
		id:           id,
		logger:       log.With().Str("sequencer", id.String()).Logger(),
		EventService: es,

		input:   make(tick, limit),
		fetch:   make(tick, limit),
		process: make(chan game.Event, limit),

		min:       make(tick, limit),
		last:      make(tick, limit),
		interrupt: make(chan struct{}, 1),
	}

	go func() {
		var last int64
		for {
			select {
			case t, ok := <-s.input:
				if !ok {
					return
				}
				if t < last {
					s.logger.Info().Int64("current", t).Int64("last", last).Msg("interrupt")
					s.interrupt <- struct{}{}
				}
				s.logger.Info().Int64("current", t).Msg("fetch post events")
				s.min <- t
				s.fetch <- t
			case t := <-s.last:
				last = t
			}
		}
	}()

	go func() {
		var min int64
		for t := range s.fetch {
			events, err := s.ListEvent(game.EventBuilder{
				Key: s.id.String(),
				Min: int(t),
			})
			if err != nil {
				s.logger.Error().Err(err).Msg("failed to fetch events")
				continue
			}
			for i, event := range events {
				select {
				case _ = <-s.interrupt:
					// Case where interrupt ticks at previous last run.
					if i != 0 {
						s.last <- 0
						break
					}
				case m := <-s.min:
					// Case where min is the tick from same event.
					if m == t {
						m = 0
					}
					if min == 0 || m < min {
						min = m
					}
				default:
				}
				ts := event.TS.UnixNano()
				if min != 0 && ts > min {
					s.logger.Info().Int64("ts", ts).Int64("min", min).Msg("skip")
					s.last <- 0
					break
				}
				s.last <- ts
				s.process <- event
			}
			s.last <- 0
		}
	}()

	go func() {
		for event := range s.process {
			s.logger.Info().Str("event", event.ID.String()).Int64("ts", event.TS.UnixNano()).Msg("run")
			callback(event)
		}
	}()
	return &s
}

// MsgHandler is the consumer function to subscribe for event ordering.
func (s *Sequencer) MsgHandler(msg *nats.Msg) {
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
	s.input <- event.TS.UnixNano()
}
