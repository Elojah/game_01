package main

import (
	"net"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
)

// Recurrer retrieves entity data associated to pc id and send it at regular ticks.
type Recurrer struct {
	game.EntityMapper

	logger zerolog.Logger
	id     game.ID
	addr   *net.UDPAddr

	ticker *time.Ticker
	done   chan struct{}
}

// NewRecurrer returns a new recurrer which sends entity data associated to pcID to addr, tick times per second.
func NewRecurrer(id game.ID, addr *net.UDPAddr, tick uint32) *Recurrer {
	return &Recurrer{
		logger: log.With().Str("recurrer", id.String()).Logger(),
		id:     id,
		addr:   addr,

		ticker: time.NewTicker(time.Second / time.Duration(tick)),
		done:   make(chan struct{}),
	}
}

// Close close the tick sender.
func (r *Recurrer) Close() {
	r.done <- struct{}{}
	r.ticker.Stop()
}

func (r *Recurrer) tick() {
	for {
		select {
		case _ = <-r.done:
			return
		case _ = <-r.ticker.C:
			// TODO Retrieve around entities
		}
	}
}

// Start init a new goroutine to tick regularly.
func (r *Recurrer) Start() {
	go r.tick()
}
