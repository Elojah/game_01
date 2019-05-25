package srg

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	gulid "github.com/elojah/game_01/pkg/ulid"
)

const (
	lootKey = "loot:"
)

// FetchLoot implements LootStore method with redis.
func (s *Store) FetchLoot(id gulid.ID) (bool, error) {
	if err := s.Get(lootKey + id.String()).Err(); err != nil {
		if err != redis.Nil {
			return false, errors.Wrapf(err, "fetch loot %s", id.String())
		}
		return false, nil
	}
	return true, nil
}

// UpsertLoot implements LootStore method with redis.
func (s *Store) UpsertLoot(id gulid.ID) error {
	return errors.Wrapf(
		s.Set(lootKey+id.String(), 1, 0).Err(),
		"upsert loot %s",
		id.String(),
	)
}

// RemoveLoot implements LootStore method with redis.
func (s *Store) RemoveLoot(id gulid.ID) error {
	return errors.Wrapf(
		s.Del(lootKey+id.String()).Err(),
		"del loot %s",
		id.String(),
	)
}
