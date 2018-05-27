package tile38

import (
	"fmt"
	"github.com/go-redis/redis"

	"github.com/elojah/game_01"
)

// SetEntityPosition is the tile38 implementation of EntityPosition service,
func (s *Service) SetEntityPosition(entity game.Entity, ts int64) error {
	cmd := redis.NewStringCmd(
		"SET",
		"entity",
		entity.ID.String(),
		"FIELD",
		"ts",
		ts,
		"POINT",
		entity.Position.X,
		entity.Position.Y,
	)
	if err := s.Process(cmd); err != nil {
		return err
	}
	_, err := cmd.Result()
	return err
}

// ListEntityPosition is the tile38 implementation of EntityPosition service,
func (s *Service) ListEntityPosition(subset game.EntityPositionSubset) ([]game.Entity, error) {
	cmd := redis.NewStringCmd(
		"NEARBY",
		"entity",
		"POINT",
		subset.Position.X,
		subset.Position.Y,
		subset.Radius,
	)
	if err := s.Process(cmd); err != nil {
		return nil, err
	}
	val, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	fmt.Println(string(val))
	return []game.Entity{}, nil
}
