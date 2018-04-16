package tile38

import (
	"github.com/elojah/game_01"
	"github.com/go-redis/redis"
)

func (s *Service) CreateActor(actors []game.Actor) error {
	query := `
		SET actor %s point %f %f
	`
	cmd, err := redis.NewCmd()
	return nil
}

func (s *Service) UpdateActor(subset game.ActorSubset, patch game.ActorPatch) error {
	return nil
}

func (s *Service) DeleteActor(subset game.ActorSubset) error {
	return nil
}

func (s *Service) ListActor(subset game.ActorSubset) ([]game.Actor, error) {
	return nil, nil
}
