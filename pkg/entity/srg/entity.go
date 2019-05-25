package srg

import (
	"strconv"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

const (
	entityKey = "entity:"
)

// Upsert implemented with redis.
func (s *Store) Upsert(e entity.E, ts uint64) error {
	e.State = gulid.NewID()
	raw, err := e.Marshal()
	if err != nil {
		return errors.Wrapf(err, "upsert entity %s at %d", e.ID.String(), ts)
	}
	return errors.Wrapf(s.ZAddNX(
		entityKey+e.ID.String(),
		redis.Z{
			Score:  float64(ts),
			Member: raw,
		},
	).Err(), "upsert entity %s at %d", e.ID.String(), ts)
}

// Fetch retrieves entity in Redis using ZRevRangeByScore.
func (s *Store) Fetch(id gulid.ID, max uint64) (entity.E, error) {
	vals, err := s.ZRevRangeByScore(
		entityKey+id.String(),
		redis.ZRangeBy{
			Count: 1,
			Min:   "-inf",
			Max:   strconv.FormatUint(max, 10),
		},
	).Result()
	if err != nil {
		return entity.E{}, errors.Wrapf(err, "fetch entity %s at %d", id.String(), max)
	}
	if len(vals) == 0 {
		return entity.E{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: entityKey, Index: id.String()},
			"fetch entity %s at %d",
			id.String(),
			max,
		)
	}
	var e entity.E
	if err := e.Unmarshal([]byte(vals[0])); err != nil {
		return entity.E{}, errors.Wrapf(err, "fetch entity %s at %d", id.String(), max)
	}
	return e, nil
}

// Remove deletes all entity states in redis.
func (s *Store) Remove(id gulid.ID) error {
	return errors.Wrapf(s.Del(entityKey+id.String()).Err(), "remove entity %s", id.String())
}

// RemoveByTS deletes entity states in redis from minTS to +inf.
func (s *Store) RemoveByTS(id gulid.ID, min uint64) error {
	return errors.Wrapf(s.ZRemRangeByScore(
		entityKey+id.String(),
		strconv.FormatUint(min, 10),
		"+inf",
	).Err(), "remove entity %s with min TS %d", id.String(), min)
}
