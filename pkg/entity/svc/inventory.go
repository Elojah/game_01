package svc

import (
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

// InventoryService wraps entity inventory operations.
type InventoryService struct {
	EntityInventoryStore   entity.InventoryStore
	EntityMRInventoryStore entity.MRInventoryStore
}

// Get retrieve an inventory and check most recent store if not found.
func (s *InventoryService) Get(id ulid.ID, entityID ulid.ID) (entity.Inventory, error) {
	inv, err := s.EntityInventoryStore.GetInventory(id)
	if err != nil {
		switch errors.Cause(err).(type) {
		case gerrors.ErrNotFound:
			inv, err := s.EntityMRInventoryStore.GetMRInventory(entityID)
			if err != nil {
				return entity.Inventory{}, errors.Wrap(err, "get inventory")
			}
			return inv, nil
		}
		return entity.Inventory{}, errors.Wrap(err, "get inventory")
	}
	return inv, nil
}

// SetMR set an inventory as most recent and traditional store.
func (s *InventoryService) SetMR(entityID ulid.ID, inv entity.Inventory) error {
	if err := s.EntityInventoryStore.SetInventory(inv); err != nil {
		return errors.Wrap(err, "set mr inventory")
	}
	if err := s.EntityMRInventoryStore.SetMRInventory(entityID, inv); err != nil {
		return errors.Wrap(err, "set mr inventory")
	}
	return nil
}
