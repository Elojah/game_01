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
func NewSequencer(id game.ID, es game.EventService, callback func(game.Event)) *Sequencer {
	s := Sequencer{
		id:           id,
		logger:       log.With().Str("sequencer", id.String()).Logger(),
		EventService: es,

		input:   make(tick, 32),
		fetch:   make(tick, 32),
		process: make(chan game.Event, 32),

		last:      make(tick, 32),
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
				s.fetch <- t
			case t, ok := <-s.last:
				if !ok {
					return
				}
				last = t
			}
		}
	}()

	go func() {
		var min int64
		for {
			select {
			case t, ok := <-s.fetch:
				if !ok {
					return
				}

				events, err := s.ListEvent(game.EventBuilder{
					Key: s.id.String(),
					Min: int(t),
				})
				if err != nil {
					s.logger.Error().Err(err).Msg("failed to fetch events")
					break
				}

				send := func(event game.Event) {
					s.last <- event.TS.UnixNano()
					s.process <- event
				}
				func() {
					for i, event := range events {
						select {
						case _ = <-s.interrupt:
							if i != 0 {
								return
							}
							// Happens when s.last has not been set at 0 yet but interrupt has been sent.
							send(event)
						default:
							send(event)
						}
					}
					s.last <- 0
				}()
			}
		}
	}()

	go func() {
		for {
			select {
			case event, ok := <-s.process:
				if !ok {
					return
				}
				s.logger.Info().Str("event", event.ID.String()).Msg("run")
				callback(event)
			}
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
