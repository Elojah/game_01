package srg

import (
	"strconv"

	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/errors"
)

const (
	entityKey = "entity:"
)

// SetEntity implemented with redis.
func (s *Store) SetEntity(e entity.E, ts int64) error {
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

// GetEntity retrieves entity in Redis using ZRangeWithScores.
func (s *Store) GetEntity(subset entity.Subset) (entity.E, error) {
	cmd := s.ZRevRangeByScore(
		entityKey+subset.ID.String(),
		redis.ZRangeBy{
			Count: 1,
			Min:   "-inf",
			Max:   strconv.FormatInt(subset.MaxTS, 10),
		},
	)
	vals, err := cmd.Result()
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

// DelEntity deletes entity in redis.
func (s *Store) DelEntity(subset entity.Subset) error {
	if subset.MinTS == 0 {
		return s.Del(entityKey + subset.ID.String()).Err()
	}
	return s.ZRemRangeByScore(
		entityKey+subset.ID.String(),
		strconv.FormatInt(subset.MinTS, 10),
		"+inf",
	).Err()
}