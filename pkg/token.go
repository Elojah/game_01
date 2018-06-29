package pkg

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/perm"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
)

// Token wraps use cases around token object.
type Token struct {
	account.TokenMapper

	EntityMapper entity.Mapper
	entity.PCMapper

	event.QRecurrerMapper
	event.QListenerMapper

	perm.Mapper

	sector.EntitiesMapper
}

// Disconnect closes a token and all entities/listener/sync associated.
func (t Token) Disconnect(id ulid.ID) error {
	logger := log.With().
		Str("token", id.String()).
		Str("action", "close").
		Logger()

	// #Retrieve token
	token, err := t.GetToken(account.TokenSubset{ID: id})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve token")
		return err
	}

	// #Retrieve entity
	e, err := t.EntityMapper.GetEntity(entity.Subset{
		ID:    token.Entity,
		MaxTS: time.Now().UnixNano(),
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve entity")
		return err
	}

	// #Save last entity state into PC
	pc := entity.PC(e)
	pc.ID = token.PC
	if err := t.SetPC(pc, token.Account); err != nil {
		logger.Error().Err(err).Msg("failed to save pc")
		return err
	}

	// #Close token listener
	go func() {
		if err := t.SendListener(event.Listener{Action: event.Close}, token.CorePool); err != nil {
			logger.Error().Err(err).Msg("failed to close listener")
		}
	}()

	// #Close token recurrer
	go func() {
		if err := t.SendRecurrer(event.Recurrer{Action: event.Close}, token.SyncPool); err != nil {
			logger.Error().Err(err).Msg("failed to close recurrer")
		}
	}()

	// #Delete token permission on entity.
	go func() {
		if err := t.DelPermission(perm.Subset{
			Source: token.ID.String(),
			Target: token.Entity.String(),
		}); err != nil {
			logger.Error().Err(err).Msg("failed to delete entity")
		}
	}()

	// #Delete pc entity
	go func() {
		if err := t.EntityMapper.DelEntity(entity.Subset{ID: token.Entity}); err != nil {
			logger.Error().Err(err).Msg("failed to delete entity")
		}
	}()

	return nil
}
