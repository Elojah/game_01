package svc

import (
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/sector"
	gulid "github.com/elojah/game_01/pkg/ulid"
	"github.com/oklog/ulid"

	"github.com/pkg/errors"
)

// Service represents entity usecases.
type Service struct {
	EntityStore           entity.Store
	EntityPermissionStore entity.PermissionStore
	SectorEntitiesStore   sector.EntitiesStore

	SequencerService infra.SequencerService
}

// Disconnect disconnects an entity.
func (s Service) Disconnect(id gulid.ID) error {

	e, err := s.EntityStore.GetEntity(id, ulid.Now())
	if err != nil {
		return errors.Wrap(err, "disconnect entity")
	}

	// #Close entity sequencer
	if err := s.SequencerService.Remove(id); err != nil {
		return errors.Wrap(err, "disconnect entity")
	}

	// #Delete pc entity position
	if err := s.SectorEntitiesStore.RemoveEntityFromSector(id, e.Position.SectorID); err != nil {
		return errors.Wrap(err, "disconnect entity")
	}

	// #Delete pc entity
	if err := s.EntityStore.DelEntity(id); err != nil {
		return errors.Wrap(err, "disconnect entity")
	}

	return nil
}
