package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/storage"
	"github.com/elojah/game_01/pkg/ulid"
)

type tick chan int64

// Sequencer is an ordering/event extractor layer between two consumers.
type Sequencer struct {
	id ulid.ID
	event.Mapper

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
func NewSequencer(id ulid.ID, limit int, em event.Mapper, callback func(ulid.ID, event.E)) *Sequencer {
	return &Sequencer{
		id:     id,
		logger: log.With().Str("sequencer", id.String()).Logger(),
		Mapper: em,

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
}

func (s *Sequencer) listenFetch() {
	var min int64
	for t := range s.fetch {
		events, err := s.ListEvent(event.Subset{
			Key: s.id.String(),
			Min: t,
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
}

func (s *Sequencer) listenProcess() {
	for event := range s.process {
		s.logger.Info().Str("event", event.ID.String()).Int64("ts", event.TS.UnixNano()).Msg("run")
		s.callback(s.id, event)
	}
}

// Start starts the 3 goroutines to follow up events.
func (s *Sequencer) Run() {
	go s.listenInput()
	go s.listenFetch()
	go s.listenProcess()
}

// Handler is the consumer function to subscribe for event ordering.
func (s *Sequencer) Handler(msg *event.Message) {
	var eventS storage.Event
	if _, err := eventS.Unmarshal([]byte(msg.Payload)); err != nil {
		s.logger.Error().Err(err).Msg("error unmarshaling event")
		return
	}
	event := eventS.Domain()
	if err := s.SetEvent(event, s.id); err != nil {
		s.logger.Error().Err(err).Msg("error creating event")
		return
	}
	s.logger.Info().Str("event", event.ID.String()).Msg("event received")
	s.input <- event.TS.UnixNano()
}
