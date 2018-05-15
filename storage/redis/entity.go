package redis

import (
	"strconv"

	"github.com/go-redis/redis"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

const (
	entityKey = "entity:"
)

// CreateEntity implemented with redis.
func (s *Service) CreateEntity(entity game.Entity, ts int64) error {
	raw, err := storage.NewEntity(entity).Marshal(nil)
	if err != nil {
		return err
	}
	return s.ZAddNX(
		entityKey+entity.ID.String(),
		redis.Z{
			Score:  float64(ts),
			Member: raw,
		},
	).Err()
}

// GetEntity retrieves entity in Redis using ZRangeWithScores.
func (s *Service) GetEntity(builder game.EntityBuilder) (game.Entity, error) {
	cmd := s.ZRevRangeByScore(
		entityKey+builder.Key,
		redis.ZRangeBy{
			Count: 1,
			Min:   "-inf",
			Max:   strconv.Itoa(builder.Max),
		},
	)
	vals, err := cmd.Result()
	if err != nil {
		return game.Entity{}, err
	}
	if len(vals) == 0 {
		return game.Entity{}, storage.ErrNotFound
	}
	var entityS storage.Entity
	if _, err := entityS.Unmarshal([]byte(vals[0])); err != nil {
		return game.Entity{}, err
	}
	return entityS.Domain(), nil
}
