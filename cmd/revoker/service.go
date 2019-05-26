package main

import (
	"github.com/oklog/ulid"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

type service struct {
	account account.App

	lifespan uint64
}

// Dial starts the auth server.
func (svc *service) Dial(c Config) error {
	svc.lifespan = c.Lifespan
	go svc.Run()
	return nil
}

// Close shutdowns the server listening.
func (svc *service) Close() error {
	return nil
}

// Run start the revoker
func (svc *service) Run() {
	logger := log.With().Str("revoker", "").Logger()

	tokenIDs, err := svc.account.ListTokenHC(ulid.Now() - svc.lifespan)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve expired tokens")
	}
	for _, tokenID := range tokenIDs {
		go func(tokenID gulid.ID) {
			if err := svc.account.DisconnectToken(tokenID); err != nil {
				logger.Error().Err(err).Str("token", tokenID.String()).Msg("failed to disconnect expired token")
			}
		}(tokenID)
	}
}
