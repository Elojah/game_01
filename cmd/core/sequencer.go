package main

import (
	"fmt"
	"math"
	"sync/atomic"
	"time"

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
	output  chan game.Event
	fetcher tick

	current   int64
	min       int64
	interrupt chan struct{}
}

// Close kills both fetch/input goroutines.
func (s *Sequencer) Close() {
	s.logger.Info().Msg("close sequencer")
	close(s.input)
	close(s.fetcher)
	close(s.output)
}

// NewSequencer returns a new sequencer with two listening goroutines to fetch/order events.
func NewSequencer(id game.ID, es game.EventService, callback func(game.Event)) *Sequencer {
	s := Sequencer{
		id:           id,
		logger:       log.With().Str("sequencer", id.String()).Logger(),
		EventService: es,
		fetcher:      make(tick, 32),
		input:        make(tick, 32),
		output:       make(chan game.Event, 32),
		min:          math.MaxInt64,
		current:      0,
	}
	go func() {
		for {
			select {
			case t, ok := <-s.input:
				fmt.Println("input:", t.String(), ok)
				if !ok {
					break
				}
				if t < atomic.LoadInt64(&s.current) {
					s.interrupt <- struct{}{}
				}
				if t < atomic.LoadInt64(&s.min) {
					atomic.StoreInt64(&s.min, t)
				}
				s.fetcher <- t
			}
		}
	}()
	go func() {
		for {
			select {
			case t, ok := <-s.fetcher:
				fmt.Println("fetcher:", t.String(), ok)
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
				loop := func() {
					defer atomic.StoreInt64(&s.current, 0)
					for _, event := range events {
						select {
						case s.interrupt:
							fmt.Println("break event loop by interrupt")
							return
						default:
							ts := event.TS.UnixNano()
							atomic.StoreInt64(&s.current, ts)
							if ts > atomic.LoadInt64(&s.min) {
								fmt.Println("break event loop by min value")
								return
							}
							s.output <- event
						}
					}
					atomic.StoreInt64(&s.min, math.MaxInt64)
				}()
			}
		}
	}()
	go func() {
		for {
			select {
			case event, ok := <-s.output:
				fmt.Println("output:", event.TS.String(), ok)
				if !ok {
					return
				}
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
