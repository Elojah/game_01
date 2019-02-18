package srg

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	gulid "github.com/elojah/game_01/pkg/ulid"
)

const (
	lootKey = "loot:"
)

// GetLoot implements LootStore method with redis.
func (s *Store) GetLoot(id gulid.ID) (bool, error) {
	if err := s.Get(lootKey + id.String()).Err(); err != nil {
		if err != redis.Nil {
			return false, errors.Wrapf(err, "get loot %s", id.String())
		}
		return false, nil
	}
	return true, nil
}

// SetLoot implements LootStore method with redis.
func (s *Store) SetLoot(id gulid.ID) error {
	return errors.Wrapf(
		s.Set(lootKey+id.String(), 1, 0).Err(),
		"set loot %s",
		id.String(),
	)
}

// DelLoot implements LootStore method with redis.
func (s *Store) DelLoot(id gulid.ID) error {
	return errors.Wrapf(
		s.Del(lootKey+id.String()).Err(),
		"del loot %s",
		id.String(),
	)
}
