package entity

import (
	"time"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/game_01/pkg/usecase/listener"
	"github.com/pkg/errors"
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

	ent, err := e.EntityMapper.GetEntity(entity.Subset{ID: id, MaxTS: time.Now().UnixNano()})
	if err != nil {
		return errors.Wrapf(err, "get entity %s", ulid.String(id))
	}

	// #Close entity listener
	if err := e.L.Delete(id); err != nil {
		return errors.Wrapf(err, "delete listener %s", ulid.String(id))
	}

	// #Delete token permission on entity.
	if err := e.DelPermission(entity.PermissionSubset{
		Source: ulid.String(tok.ID),
		Target: ulid.String(id),
	}); err != nil {
		return errors.Wrapf(err, "delete permission %s %s", ulid.String(tok.ID), ulid.String(id))
	}

	// #Delete pc entity position
	if err := e.RemoveEntityToSector(id, ent.Position.SectorID); err != nil {
		return errors.Wrapf(err, "remove entity %s to sector %s", ulid.String(id), ulid.String(ent.Position.SectorID))
	}

	// #Delete pc entity
	if err := e.EntityMapper.DelEntity(entity.Subset{ID: id}); err != nil {
		return errors.Wrapf(err, "delete entity %s", ulid.String(id))
	}

	return nil
}
