package srg

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	inventoryKey = "inventory:"
)

// GetInventory implemented with redis.
func (s *Store) GetInventory(id ulid.ID) (entity.Inventory, error) {
	val, err := s.Get(inventoryKey + id.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.Inventory{}, err
		}
		return entity.Inventory{}, errors.ErrNotFound{Store: inventoryKey, Index: id.String()}
	}

	var inv entity.Inventory
	if err := inv.Unmarshal([]byte(val)); err != nil {
		return entity.Inventory{}, err
	}
	return inv, nil
}

// SetInventory implemented with redis.
func (s *Store) SetInventory(inventory entity.Inventory) error {
	raw, err := inventory.Marshal()
	if err != nil {
		return err
	}
	return s.Set(inventoryKey+inventory.ID.String(), raw, 0).Err()
}

// DelInventory implemented with redis.
func (s *Store) DelInventory(id ulid.ID) error {
	return s.Del(inventoryKey + id.String()).Err()
}
