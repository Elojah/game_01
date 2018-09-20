package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
)

type tick chan ulid.ID

// Sequencer is an ordering/event extractor layer between two consumers.
type Sequencer struct {
	id ulid.ID

	EventStore  event.Store
	EntityStore entity.Store

	logger zerolog.Logger

	input   tick
	fetch   tick
	process chan event.E

	min       tick
	last      tick
	interrupt chan struct{}

	callback func(ulid.ID, event.E)
}

// Close kills both fetch/input goroutines.
func (s *Sequencer) Close() {
	s.logger.Info().Msg("close sequencer")
	close(s.input)
	close(s.fetch)
	close(s.process)
}

// NewSequencer returns a new sequencer with two listening goroutines to fetch/order events.
func NewSequencer(id ulid.ID, limit int, callback func(ulid.ID, event.E)) *Sequencer {
	return &Sequencer{
		id:     id,
		logger: log.With().Str("sequencer", id.String()).Logger(),

		input:   make(tick, limit),
		fetch:   make(tick, limit),
		process: make(chan event.E, limit),

		min:       make(tick, limit),
		last:      make(tick, limit),
		interrupt: make(chan struct{}, 1),

		callback: callback,
	}
}

func (s *Sequencer) listenInput() {
	var last ulid.ID
	for {
		select {
		case id, ok := <-s.input:
			if !ok {
				return
			}
			if id.Compare(last) < 0 {
				s.logger.Info().Str("event", id.String()).Str("last", last.String()).Msg("interrupt")
				s.interrupt <- struct{}{}
			}
			s.logger.Info().Str("event", id.String()).Msg("fetch post events")
			s.min <- id
			s.fetch <- id
		case id := <-s.last:
			last = id
		}
	}
}

func (s *Sequencer) listenFetch() {
	var min ulid.ID
	for id := range s.fetch {
		if err := s.EntityStore.DelEntity(entity.Subset{ID: s.id, MinTS: id.Time()}); err != nil {
			s.logger.Error().Err(err).Msg("failed to clear entities")
			continue
		}
		events, err := s.EventStore.ListEvent(event.Subset{
			Key: s.id.String(),
			Min: id,
		})
		if err != nil {
			s.logger.Error().Err(err).Msg("failed to fetch events")
			continue
		}
	Event:
		for i, event := range events {
			select {
			case _ = <-s.interrupt:
				// Case where interrupt ticks at previous last run but not consumed. e.g: on last iteration
				if i != 0 {
					s.last <- ulid.Zero()
					break Event
				}
			case m := <-s.min:
				// min is the currently consumed event so we reset min value.
				s.logger.Info().Str("m", m.String()).Str("min", min.String()).Msg("m equal min ?")
				if m.Equal(min) {
					min = ulid.Zero()
				}
				// if min is not set yet or new value is inferior to min.
				if min.IsZero() || m.Compare(min) < 0 {
					min = m
				}
			default:
			}
			if !min.IsZero() && min.Compare(event.ID) < 0 {
				s.logger.Info().Msg("skip for earlier value in queue")
				break
			}
			s.last <- event.ID
			s.process <- event
		}
		s.last <- ulid.Zero()
	}
}

func (s *Sequencer) listenProcess() {
	for event := range s.process {
		s.logger.Info().Str("event", event.ID.String()).Msg("apply")
		s.callback(s.id, event)
	}
}

// Run starts the 3 goroutines to follow up events.
func (s *Sequencer) Run() {
	go s.listenInput()
	go s.listenFetch()
	go s.listenProcess()
}

// Handler is the consumer function to subscribe for event ordering.
func (s *Sequencer) Handler(msg *infra.Message) {
	var e event.E
	if err := e.Unmarshal([]byte(msg.Payload)); err != nil {
		s.logger.Error().Err(err).Msg("error unmarshaling event")
		return
	}
	if err := s.EventStore.SetEvent(e, s.id); err != nil {
		s.logger.Error().Err(err).Msg("error creating event")
		return
	}
	s.logger.Info().Str("event", e.ID.String()).Msg("event received")
	s.input <- e.ID
}
