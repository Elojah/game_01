package srg

import (
	"strconv"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	entityKey = "entity:"
)

// SetEntity implemented with redis.
func (s *Store) SetEntity(e entity.E, ts uint64) error {
	raw, err := e.Marshal()
	if err != nil {
		return errors.Wrapf(err, "set entity %s at %d", e.ID.String(), ts)
	}
	return errors.Wrapf(s.ZAddNX(
		entityKey+e.ID.String(),
		redis.Z{
			Score:  float64(ts),
			Member: raw,
		},
	).Err(), "set entity %s at %d", e.ID.String(), ts)
}

// GetEntity retrieves entity in Redis using ZRevRangeByScore.
func (s *Store) GetEntity(id ulid.ID, max uint64) (entity.E, error) {
	vals, err := s.ZRevRangeByScore(
		entityKey+id.String(),
		redis.ZRangeBy{
			Count: 1,
			Min:   "-inf",
			Max:   strconv.FormatUint(max, 10),
		},
	).Result()
	if err != nil {
		return entity.E{}, errors.Wrapf(err, "get entity %s at %d", id.String(), max)
	}
	if len(vals) == 0 {
		return entity.E{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: entityKey, Index: id.String()},
			"get entity %s at %d",
			id.String(),
			max,
		)
	}
	var e entity.E
	if err := e.Unmarshal([]byte(vals[0])); err != nil {
		return entity.E{}, errors.Wrapf(err, "get entity %s at %d", id.String(), max)
	}
	return e, nil
}

// DelEntity deletes all entity states in redis.
func (s *Store) DelEntity(id ulid.ID) error {
	return errors.Wrapf(s.Del(entityKey+id.String()).Err(), "del entity %s", id.String())
}

// DelEntityByTS deletes entity states in redis from minTS to +inf.
func (s *Store) DelEntityByTS(id ulid.ID, min uint64) error {
	return errors.Wrapf(s.ZRemRangeByScore(
		entityKey+id.String(),
		strconv.FormatUint(min, 10),
		"+inf",
	).Err(), "del entity %s with min TS %d", id.String(), min)
}
