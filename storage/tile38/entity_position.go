package tile38

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01"
)

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
	if err := client.Process(cmd); err != nil {
		return err
	}
	_, err := cmd.Result()
	return err
}

func (s *Service) ListEntityPosition(subset game.EntityPositionSubset) (game.Entity, error) {
	cmd := redis.NewStringCmd(
		"NEARBY",
		"entity",
		"POINT",
		subset.Position.X,
		subset.Position.Y,
	)
	if err := client.Process(cmd); err != nil {
		return err
	}
	_, err := cmd.Result()
	return err
}
