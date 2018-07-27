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
		return errors.Wrapf(err, "get entity %s", id.String())
	}

	// #Close entity listener
	if err := e.L.Delete(id); err != nil {
		return errors.Wrapf(err, "delete listener %s", id.String())
	}

	// #Delete token permission on entity.
	if err := e.DelPermission(entity.PermissionSubset{
		Source: tok.ID.String(),
		Target: id.String(),
	}); err != nil {
		return errors.Wrapf(err, "delete permission %s %s", tok.ID.String(), id.String())
	}

	// #Delete pc entity position
	if err := e.RemoveEntityToSector(id, ent.Position.SectorID); err != nil {
		return errors.Wrapf(err, "remove entity %s to sector %s", id.String(), ent.Position.SectorID.String())
	}

	// #Delete pc entity
	if err := e.EntityMapper.DelEntity(entity.Subset{ID: id}); err != nil {
		return errors.Wrapf(err, "delete entity %s", id.String())
	}

	return nil
}
