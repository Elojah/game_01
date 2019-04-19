package srg

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	inventoryKey   = "inventory:"
	inventoryMRKey = "inventory_mr:"
)

// GetInventory implemented with redis.
func (s *Store) GetInventory(id ulid.ID) (entity.Inventory, error) {
	val, err := s.Get(inventoryKey + id.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.Inventory{}, errors.Wrapf(err, "get inventory %s", id.String())
		}
		return entity.Inventory{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: inventoryKey, Index: id.String()},
			"get inventory %s",
			id.String(),
		)
	}

	var inv entity.Inventory
	if err := inv.Unmarshal([]byte(val)); err != nil {
		return entity.Inventory{}, errors.Wrapf(err, "get inventory %s", id.String())
	}
	return inv, nil
}

// SetInventory implemented with redis.
func (s *Store) SetInventory(inv entity.Inventory) error {
	raw, err := inv.Marshal()
	if err != nil {
		return errors.Wrapf(err, "set inventory %s", inv.ID.String())
	}
	return errors.Wrapf(s.Set(inventoryKey+inv.ID.String(), raw, 0).Err(), "set inventory %s", inv.ID.String())
}

// DelInventory implemented with redis.
func (s *Store) DelInventory(id ulid.ID) error {
	return errors.Wrapf(s.Del(inventoryKey+id.String()).Err(), "del inventory %s", id.String())
}

// GetMRInventory returns the Most Recent inventory saved for entityID.
func (s *Store) GetMRInventory(entityID ulid.ID) (entity.Inventory, error) {

	val, err := s.Get(inventoryMRKey + entityID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.Inventory{}, errors.Wrapf(err, "get most recent inventory for entity %s", entityID.String())
		}
		return entity.Inventory{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: inventoryMRKey, Index: entityID.String()},
			"get most recent inventory for entity %s",
			entityID.String(),
		)
	}

	var inv entity.Inventory
	if err := inv.Unmarshal([]byte(val)); err != nil {
		return entity.Inventory{}, errors.Wrapf(err, "get inventory %s", entityID.String())
	}
	return inv, nil
}

// SetMRInventory set most recent inventory for entityID.
func (s *Store) SetMRInventory(entityID ulid.ID, inv entity.Inventory) error {

	raw, err := inv.Marshal()
	if err != nil {
		return errors.Wrapf(err, "set most recent inventory for entity %s", entityID.String())
	}
	return errors.Wrapf(
		s.Set(inventoryMRKey+entityID.String(), raw, 0).Err(),
		"set most recent inventory for entity %s",
		entityID.String(),
	)
}
