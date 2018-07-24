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
	account.TokenHCMapper
	account.TokenMapper

	EntityMapper entity.Mapper
	entity.PCMapper
	entity.PermissionMapper

	infra.QRecurrerMapper
	infra.QListenerMapper

	listener.L
	recurrer.R

	sector.EntitiesMapper

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
		TokenMapper:      a.TokenMapper,
		EntityMapper:     a.EntityMapper,
		PCMapper:         a.PCMapper,
		L:                a.L,
		R:                a.R,
		PermissionMapper: a.PermissionMapper,
		EntitiesMapper:   a.EntitiesMapper,
	}
	for _, tokenID := range tokenIDs {
		go func(tokenID ulid.ID) {
			t.Disconnect(tokenID)
		}(tokenID)
	}
}
