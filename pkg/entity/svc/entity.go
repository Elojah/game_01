package svc

import (
	"time"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"

	"github.com/pkg/errors"
)

// Service represents entity usecases.
type Service struct {
	Entity           entity.Store
	EntityPermission entity.PermissionStore
	SectorEntities   sector.EntitiesStore

	SequencerService infra.SequencerService
}

// Disconnect disconnects an entity.
func (s Service) Disconnect(id ulid.ID, t account.Token) error {

	e, err := s.Entity.GetEntity(id, time.Now().Unix())
	if err != nil {
		return errors.Wrapf(err, "get entity %s", id.String())
	}

	// #Close entity sequencer
	if err := s.SequencerService.Remove(id); err != nil {
		return errors.Wrapf(err, "delete sequencer %s", id.String())
	}

	// #Delete token permission on entity
	if err := s.EntityPermission.DelPermission(entity.PermissionSubset{
		Source: t.ID.String(),
		Target: id.String(),
	}); err != nil {
		return errors.Wrapf(err, "delete permission %s %s", t.ID.String(), id.String())
	}

	// #Delete pc entity position
	if err := s.SectorEntities.RemoveEntityFromSector(id, e.Position.SectorID); err != nil {
		return errors.Wrapf(err, "remove entity %s from sector %s", id.String(), e.Position.SectorID.String())
	}

	// #Delete pc entity
	if err := s.Entity.DelEntity(id); err != nil {
		return errors.Wrapf(err, "delete entity %s", id.String())
	}

	return nil
}
