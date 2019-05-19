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

// FetchInventory implemented with redis.
func (s *Store) FetchInventory(id ulid.ID) (entity.Inventory, error) {
	val, err := s.Get(inventoryKey + id.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.Inventory{}, errors.Wrapf(err, "fetch inventory %s", id.String())
		}
		return entity.Inventory{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: inventoryKey, Index: id.String()},
			"fetch inventory %s",
			id.String(),
		)
	}

	var inv entity.Inventory
	if err := inv.Unmarshal([]byte(val)); err != nil {
		return entity.Inventory{}, errors.Wrapf(err, "fetch inventory %s", id.String())
	}
	return inv, nil
}

// InsertInventory implemented with redis.
func (s *Store) InsertInventory(inv entity.Inventory) error {
	raw, err := inv.Marshal()
	if err != nil {
		return errors.Wrapf(err, "insert inventory %s", inv.ID.String())
	}
	return errors.Wrapf(s.Set(inventoryKey+inv.ID.String(), raw, 0).Err(), "insert inventory %s", inv.ID.String())
}

// RemoveInventory implemented with redis.
func (s *Store) RemoveInventory(id ulid.ID) error {
	return errors.Wrapf(s.Del(inventoryKey+id.String()).Err(), "remove inventory %s", id.String())
}

// FetchMRInventory returns the Most Recent inventory saved for entityID.
func (s *Store) FetchMRInventory(entityID ulid.ID) (entity.Inventory, error) {

	val, err := s.Get(inventoryMRKey + entityID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.Inventory{}, errors.Wrapf(err, "fetch mr inventory for entity %s", entityID.String())
		}
		return entity.Inventory{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: inventoryMRKey, Index: entityID.String()},
			"fetch mr inventory for entity %s",
			entityID.String(),
		)
	}

	var inv entity.Inventory
	if err := inv.Unmarshal([]byte(val)); err != nil {
		return entity.Inventory{}, errors.Wrapf(err, "fetch inventory %s", entityID.String())
	}
	return inv, nil
}

// InsertMRInventory insert mr inventory for entityID.
func (s *Store) InsertMRInventory(entityID ulid.ID, inv entity.Inventory) error {

	raw, err := inv.Marshal()
	if err != nil {
		return errors.Wrapf(err, "insert mr inventory for entity %s", entityID.String())
	}
	return errors.Wrapf(
		s.Set(inventoryMRKey+entityID.String(), raw, 0).Err(),
		"insert mr inventory for entity %s",
		entityID.String(),
	)
}

// RemoveMRInventory implemented with redis.
func (s *Store) RemoveMRInventory(entityID ulid.ID) error {
	return errors.Wrapf(s.Del(inventoryMRKey+entityID.String()).Err(), "remove inventory for entity %s", entityID.String())
}
