package token

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
	uce "github.com/elojah/game_01/pkg/usecase/entity"
)

// T wraps use cases around token object.
type T struct {
	account.TokenMapper

	EntityMapper entity.Mapper
	entity.PCMapper

	event.QRecurrerMapper
	event.QListenerMapper

	entity.PermissionMapper

	sector.EntitiesMapper
}

// Disconnect closes a token and all entities/listener/sync associated.
func (t T) Disconnect(id ulid.ID) error {
	logger := log.With().
		Str("token", id.String()).
		Str("action", "close").
		Logger()

	// #Retrieve token
	tok, err := t.GetToken(account.TokenSubset{ID: id})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve token")
		return err
	}

	// #Close token listener
	go func() {
		if err := t.PublishListener(event.Listener{ID: tok.ID, Action: event.Close}, tok.CorePool); err != nil {
			logger.Error().Err(err).Msg("failed to close listener")
		}
	}()

	// #Close token recurrer
	go func() {
		if err := t.SendRecurrer(event.Recurrer{ID: tok.ID, Action: event.Close}, tok.SyncPool); err != nil {
			logger.Error().Err(err).Msg("failed to close recurrer")
		}
	}()

	// #Retrieve entity
	e, err := t.EntityMapper.GetEntity(entity.Subset{
		ID:    tok.Entity,
		MaxTS: time.Now().UnixNano(),
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve entity")
		return err
	}

	// #Save last entity state into PC
	pc := entity.PC(e)
	pc.ID = tok.PC
	if err := t.SetPC(pc, tok.Account); err != nil {
		logger.Error().Err(err).Msg("failed to save pc")
		return err
	}

	// #For each entity permission
	permissions, err := t.ListPermission(entity.PermissionSubset{Source: tok.ID.String()})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve permissions")
		return err
	}
	ucentity := uce.E{
		EntityMapper:     t.EntityMapper,
		QRecurrerMapper:  t.QRecurrerMapper,
		QListenerMapper:  t.QListenerMapper,
		PermissionMapper: t.PermissionMapper,
		EntitiesMapper:   t.EntitiesMapper,
	}
	for _, permission := range permissions {
		targetID := ulid.MustParse(permission.Target)
		if err := ucentity.Disconnect(targetID, tok); err != nil {
			logger.Error().Err(err).Str("entity", targetID.String()).Msg("failed to disconnect entity")
		}
	}

	return nil
}
