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

	EventStore          event.Store
	EntityStore         entity.Store
	EventTriggerService event.TriggerService

	logger zerolog.Logger

	input   tick
	fetch   tick
	process chan event.E

	min tick

	callback func(ulid.ID, event.E)
}

// Close kills both fetch/input goroutines.
func (s *Sequencer) Close() error {
	s.logger.Info().Msg("close sequencer")
	close(s.input)
	close(s.fetch)
	close(s.process)
	return nil
}

// NewSequencer returns a new sequencer with two listening goroutines to fetch/order events.
func NewSequencer(id ulid.ID, limit int, callback func(ulid.ID, event.E)) *Sequencer {
	return &Sequencer{
		id:     id,
		logger: log.With().Str("sequencer", id.String()).Logger(),

		input:   make(tick, limit),
		fetch:   make(tick, limit),
		process: make(chan event.E, limit),

		min: make(tick, limit),

		callback: callback,
	}
}

func (s *Sequencer) listenInput() {
	for id := range s.input {
		s.logger.Info().Str("event", id.String()).Msg("fetch post events")
		s.min <- id
		s.fetch <- id
	}
}

func (s *Sequencer) listenFetch() {
	var min ulid.ID
	for id := range s.fetch {
		if !min.IsZero() && min.Compare(id) < 0 {
			s.logger.Info().Str("id", id.String()).Str("min", min.String()).Msg("skip for earlier value in queue")
			continue
		}
		if err := s.EntityStore.DelEntityByTS(s.id, id.Time()); err != nil {
			s.logger.Error().Err(err).Msg("failed to clear entities")
			continue
		}
		events, err := s.EventStore.ListEvent(s.id, id)
		if err != nil {
			s.logger.Error().Err(err).Msg("failed to fetch events")
			continue
		}
	Event:
		for _, event := range events {
			select {
			case m := <-s.min:
				// if min is not set yet or new value is inferior to min.
				if min.IsZero() || m.Compare(min) < 0 {
					min = m
				}
			default:
			}
			switch min.Compare(event.ID) {
			case 0:
				// min is the currently consumed event so we reset min value.
				min = ulid.Zero()
			case -1:
				if !min.IsZero() {
					s.logger.Info().Str("event", event.ID.String()).Str("min", min.String()).Msg("skip for earlier value in queue")
					break Event
				}
			}
			s.process <- event
		}
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
	if err := s.TriggerService.Set(e); err != nil {
		s.logger.Error().Err(err).Msg("error setting event")
		return
	}
	s.logger.Info().Str("event", e.ID.String()).Msg("event received")
	s.input <- e.ID
}
