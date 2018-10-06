package main

import (
	"github.com/oklog/ulid"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

type app struct {
	account.TokenHCStore

	TokenService account.TokenService

	lifespan uint64
}

// Dial starts the auth server.
func (a *app) Dial(c Config) error {
	a.lifespan = c.Lifespan
	go a.Run()
	return nil
}

// Close shutdowns the server listening.
func (a *app) Close() error {
	return nil
}

// Run start the revoker
func (a *app) Run() {
	logger := log.With().Str("revoker", "").Logger()

	tokenIDs, err := a.ListTokenHC(ulid.Now() - a.lifespan)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve expired tokens")
	}
	for _, tokenID := range tokenIDs {
		go func(tokenID gulid.ID) {
			a.TokenService.Disconnect(tokenID)
		}(tokenID)
	}
}
