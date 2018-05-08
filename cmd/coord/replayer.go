package main

import (
	"sync/atomic"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
)

var (
	maxTime = time.Unix(1<<63-62135596801, 999999999)
)

type tick chan time.Time

type replayer struct {
	Fetcher   tick
	Trigger   tick
	Pass      chan game.Event
	Interrupt int32
}

func (r *replayer) Close() {
	close(r.Fetcher)
	close(r.Pass)
}

func (a *app) newReplayer(id game.ID) replayer {
	r := replayer{
		Fetcher: make(tick, 0),
		Trigger: make(tick, 0),
		Pass:    make(chan game.Event, 0),
	}
	go func(id game.ID) {
		logger := log.With().Str("fetcher", id.String()).Logger()
		for {
			select {
			case t, ok := <-r.Fetcher:
				if !ok {
					return
				}
				events, err := a.ListEvent(game.EventBuilder{
					Key: id.String(),
					Min: int(t.UnixNano()),
				})
				if err != nil {
					logger.Error().Err(err).Msg("failed to fetch events")
					break
				}
				for _, event := range events {
					if atomic.CompareAndSwapInt32(&r.Interrupt, 1, 0) {
						break
					}
					r.Pass <- event
				}
			}
		}
	}(id)
	go func(id game.ID) {
		var current time.Time
		waiting := maxTime
		for {
			select {
			case t := <-r.Trigger:
				if t.Before(current) {
					atomic.CompareAndSwapInt32(&r.Interrupt, 0, 1)
				}
				waiting = t
				r.Fetcher <- t
			case event, ok := <-r.Pass:
				if !ok {
					return
				}
				if event.TS.After(waiting) {
					atomic.CompareAndSwapInt32(&r.Interrupt, 0, 1)
					waiting = maxTime
					break
				}
				current = event.TS
				a.Play(event)
			}
		}
	}(id)
	return r
}
