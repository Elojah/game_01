package main

import (
	"time"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/game_01/pkg/usecase/listener"
	"github.com/elojah/game_01/pkg/usecase/recurrer"
	"github.com/elojah/game_01/pkg/usecase/token"
	"github.com/rs/zerolog/log"
)

type app struct {
	account.TokenHCService
	account.TokenService

	EntityService entity.Service
	entity.PCService
	entity.PermissionService

	infra.QRecurrerService
	infra.QListenerService

	listener.L
	recurrer.R

	sector.EntitiesService

	lifespan time.Duration
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

	tokenIDs, err := a.ListTokenHC(account.TokenHCSubset{MaxTS: time.Now().Add(-a.lifespan).Unix()})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve expired tokens")
	}
	t := token.T{
		TokenService:      a.TokenService,
		EntityService:     a.EntityService,
		PCService:         a.PCService,
		L:                a.L,
		R:                a.R,
		PermissionService: a.PermissionService,
		EntitiesService:   a.EntitiesService,
	}
	for _, tokenID := range tokenIDs {
		go func(tokenID ulid.ID) {
			t.Disconnect(tokenID)
		}(tokenID)
	}
}
