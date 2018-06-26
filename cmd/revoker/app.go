package main

import (
	"time"

	"github.com/elojah/game_01"
	"github.com/rs/zerolog/log"
)

type app struct {
	game.EntityMapper
	game.SectorEntitiesMapper
	game.TokenHCMapper
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
	return nil
}

// Start start the revoker
func (a *app) Start() {
	logger := log.With().Str("revoker", "").Logger()

	tokenIDs, err := a.ListTokenHC(game.TokenHCSubset{MaxTS: time.Now().Add(-a.lifespan).Unix()})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve expired tokens")
	}
	for _, tokenID := range tokenIDs {
		go func(tokenID game.ID) {

		}(tokenID)
	}
}
