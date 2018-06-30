package token

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
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
	token, err := t.GetToken(account.TokenSubset{ID: id})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve token")
		return err
	}

	// #Close token listener
	go func() {
		if err := t.SendListener(event.Listener{ID: token.ID, Action: event.Close}, token.CorePool); err != nil {
			logger.Error().Err(err).Msg("failed to close listener")
		}
	}()

	// #Close token recurrer
	go func() {
		if err := t.SendRecurrer(event.Recurrer{ID: token.ID, Action: event.Close}, token.SyncPool); err != nil {
			logger.Error().Err(err).Msg("failed to close recurrer")
		}
	}()

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

	// #For each entity permission
	permissions, err := t.ListPermission(entity.PermissionSubset{Source: token.ID.String()})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve permissions")
		return err
	}
	for _, permission := range permissions {

		targetID := ulid.MustParse(permission.Target)

		// #Close entity listener
		go func() {
			if err := t.SendListener(event.Listener{ID: targetID, Action: event.Close}, token.CorePool); err != nil {
				logger.Error().Err(err).Msg("failed to close listener")
			}
		}()

		// #Close entity recurrer
		go func() {
			if err := t.SendRecurrer(event.Recurrer{ID: targetID, Action: event.Close}, token.SyncPool); err != nil {
				logger.Error().Err(err).Msg("failed to close recurrer")
			}
		}()

		// #Delete token permission on entity.
		go func() {
			if err := t.DelPermission(entity.PermissionSubset{
				Source: token.ID.String(),
				Target: targetID.String(),
			}); err != nil {
				logger.Error().Err(err).Msg("failed to delete entity")
			}
		}()

		// #Delete pc entity
		go func() {
			if err := t.EntityMapper.DelEntity(entity.Subset{ID: targetID}); err != nil {
				logger.Error().Err(err).Msg("failed to delete entity")
			}
		}()
	}

	return nil
}
