package entity

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/game_01/pkg/usecase/listener"
)

// E represents use cases for entity.
type E struct {
	EntityMapper entity.Mapper

	entity.PermissionMapper

	listener.L

	sector.EntitiesMapper
}

// Disconnect disconnects an entity.
func (e E) Disconnect(id ulid.ID, tok account.Token) error {

	logger := log.With().
		Str("entity", id.String()).
		Logger()

	ent, err := e.EntityMapper.GetEntity(entity.Subset{ID: id, MaxTS: time.Now().UnixNano()})
	if err != nil {
		logger.Error().Err(err).Str("entity", id.String()).Msg("failed to retrieve entity")
		return err
	}

	// #Close entity listener
	if err := e.L.Delete(id); err != nil {
		logger.Error().Err(err).Msg("failed to close listener")
		return err
	}

	// #Delete token permission on entity.
	if err := e.DelPermission(entity.PermissionSubset{
		Source: tok.ID.String(),
		Target: id.String(),
	}); err != nil {
		logger.Error().Err(err).Msg("failed to delete entity")
		return err
	}

	// #Delete pc entity position
	if err := e.RemoveEntityToSector(id, ent.Position.SectorID); err != nil {
		logger.Error().Err(err).Msg("failed to delete entity")
		return err
	}

	// #Delete pc entity
	if err := e.EntityMapper.DelEntity(entity.Subset{ID: id}); err != nil {
		logger.Error().Err(err).Msg("failed to delete entity")
		return err
	}

	return nil
}
