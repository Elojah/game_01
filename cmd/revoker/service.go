package main

import (
	"github.com/oklog/ulid"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

type service struct {
	account.TokenHCStore

	TokenService account.TokenService

	lifespan uint64
}

// Dial starts the auth server.
func (service *service) Dial(c Config) error {
	service.lifespan = c.Lifespan
	go service.Run()
	return nil
}

// Close shutdowns the server listening.
func (service *service) Close() error {
	return nil
}

// Run start the revoker
func (service *service) Run() {
	logger := log.With().Str("revoker", "").Logger()

	tokenIDs, err := service.ListTokenHC(ulid.Now() - service.lifespan)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve expired tokens")
	}
	for _, tokenID := range tokenIDs {
		go func(tokenID gulid.ID) {
			if err := service.TokenService.Disconnect(tokenID); err != nil {
				logger.Error().Err(err).Str("token", tokenID.String()).Msg("failed to disconnect expired token")
			}
		}(tokenID)
	}
}
