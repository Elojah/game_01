package svc

import (
	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/entity"
	gulid "github.com/elojah/game_01/pkg/ulid"
	"github.com/pkg/errors"
)

// PCService provides a clean pc removal.
type PCService struct {
	AbilityStore      ability.Store
	EntityPCStore     entity.PCStore
	EntityPCLeftStore entity.PCLeftStore
}

// RemovePC remove a pc and clean associated abilities.
func (s PCService) RemovePC(accountID gulid.ID, id gulid.ID) error {

	// Delete pc abilities
	if err := s.AbilityStore.DelAbilities(id); err != nil {
		return errors.Wrap(err, "remove pc")
	}

	// #Delete pc
	if err := s.EntityPCStore.DelPC(accountID, id); err != nil {
		return errors.Wrap(err, "remove pc")
	}

	// #Add 1 to pc left
	pcLeft, err := s.EntityPCLeftStore.GetPCLeft(accountID)
	if err != nil {
		return errors.Wrap(err, "remove pc")
	}
	pcLeft = pcLeft - 1
	if err := s.EntityPCLeftStore.SetPCLeft(pcLeft, accountID); err != nil {
		return errors.Wrap(err, "remove pc")
	}

	return nil
}
