package svc

import (
	"github.com/pkg/errors"

	"github.com/oklog/ulid"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/sector"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// Service represents entity usecases.
type Service struct {
	EntityStore           entity.Store
	EntityPermissionStore entity.PermissionStore
	AbilityStore          ability.Store
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

	// Delete entity abilities
	if err := s.AbilityStore.DelAbilities(id); err != nil {
		return errors.Wrap(err, "remove pc")
	}

	// #Delete pc entity
	if err := s.EntityStore.DelEntity(id); err != nil {
		return errors.Wrap(err, "disconnect entity")
	}

	return nil
}
