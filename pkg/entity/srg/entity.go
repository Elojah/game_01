package srg

import (
	"strconv"

	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	entityKey = "entity:"
)

// SetEntity implemented with redis.
func (s *Store) SetEntity(e entity.E, ts uint64) error {
	raw, err := e.Marshal()
	if err != nil {
		return err
	}
	return s.ZAddNX(
		entityKey+e.ID.String(),
		redis.Z{
			Score:  float64(ts),
			Member: raw,
		},
	).Err()
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
		return entity.E{}, err
	}
	if len(vals) == 0 {
		return entity.E{}, errors.ErrNotFound
	}
	var e entity.E
	if err := e.Unmarshal([]byte(vals[0])); err != nil {
		return entity.E{}, err
	}
	return e, nil
}

// DelEntity deletes all entity states in redis.
func (s *Store) DelEntity(id ulid.ID) error {
	return s.Del(entityKey + id.String()).Err()
}

// DelEntityByTS deletes entity states in redis from minTS to +inf.
func (s *Store) DelEntityByTS(id ulid.ID, min uint64) error {
	return s.ZRemRangeByScore(
		entityKey+id.String(),
		strconv.FormatUint(min, 10),
		"+inf",
	).Err()
}
