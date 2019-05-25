package srg

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	spawnKey = "spawn:"
)

// FetchSpawn implemented with redis.
func (s *Store) FetchSpawn(id ulid.ID) (entity.Spawn, error) {
	val, err := s.Get(spawnKey + id.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.Spawn{}, errors.Wrapf(err, "fetch spawn %s", id.String())
		}
		return entity.Spawn{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: spawnKey, Index: id.String()},
			"fetch spawn %s",
			id.String(),
		)
	}

	var sp entity.Spawn
	if err := sp.Unmarshal([]byte(val)); err != nil {
		return entity.Spawn{}, errors.Wrapf(err, "fetch spawn %s", id.String())
	}
	return sp, nil
}

// UpsertSpawn implemented with redis.
func (s *Store) UpsertSpawn(spawn entity.Spawn) error {
	raw, err := spawn.Marshal()
	if err != nil {
		return errors.Wrapf(err, "upsert spawn %s", spawn.ID.String())
	}
	return errors.Wrapf(
		s.Set(spawnKey+spawn.ID.String(), raw, 0).Err(),
		"upsert spawn %s",
		spawn.ID.String(),
	)
}

// RemoveSpawn implemented with redis.
func (s *Store) RemoveSpawn(id ulid.ID) error {
	return errors.Wrapf(
		s.Del(spawnKey+id.String()).Err(),
		"remove spawn %s",
		id.String(),
	)
}
