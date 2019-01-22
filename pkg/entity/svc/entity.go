package svc

import (
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/sector"
	gulid "github.com/elojah/game_01/pkg/ulid"
	"github.com/oklog/ulid"

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
func (s Service) Disconnect(id gulid.ID, t account.Token) error {

	e, err := s.Entity.GetEntity(id, ulid.Now())
	if err != nil {
		return errors.Wrap(err, "disconnect entity")
	}

	// #Close entity sequencer
	if err := s.SequencerService.Remove(id); err != nil {
		return errors.Wrap(err, "disconnect entity")
	}

	// #Delete token permission on entity
	if err := s.EntityPermission.DelPermission(t.ID.String(), id.String()); err != nil {
		return errors.Wrap(err, "disconnect entity")
	}

	// #Delete pc entity position
	if err := s.SectorEntities.RemoveEntityFromSector(id, e.Position.SectorID); err != nil {
		return errors.Wrap(err, "disconnect entity")
	}

	// #Delete pc entity
	if err := s.Entity.DelEntity(id); err != nil {
		return errors.Wrap(err, "disconnect entity")
	}

	return nil
}
