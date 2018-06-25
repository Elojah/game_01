package main

import (
	"context"
	"time"

	"github.com/elojah/game_01"
)

type app struct {
	game.EntityMapper
	game.SectorEntitiesMapper
	game.TokenMapper

	lifespan time.Duration
}

// Dial starts the auth server.
func (a *app) Dial(c Config) error {
	a.lifespan = c.Lifespan
	return nil
}

// Close shutdowns the server listening.
func (a *app) Close() error {
	return a.srv.Shutdown(context.Background())
}

// Start start the revoker
func (a *app) Start() {
}
