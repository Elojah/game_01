package tile38

import (
	"errors"

	"github.com/elojah/game_01"
	"github.com/go-redis/redis"
)

// CreateActor is the tile38 implementation for ActorService.
func (s *Service) CreateActor(actors []game.Actor) error {
	for _, actor := range actors {
		if err := s.Process(redis.NewStringCmd(
			"SET",
			"actor",
			actor.ID.String(),
			"point",
			actor.Position.X,
			actor.Position.Y,
		)); err != nil {
			return err
		}
	}
	return nil
}

// UpdateActor is the tile38 implementation for ActorService.
func (s *Service) UpdateActor(subset game.ActorSubset, patch game.ActorPatch) error {
	for _, id := range subset.IDs {
		if err := s.Process(redis.NewStringCmd(
			"SET",
			"actor",
			id.String(),
			"point",
			patch.Position.X,
			patch.Position.Y,
		)); err != nil {
			return err
		}
	}
	return nil
}

// DeleteActor is the tile38 implementation for ActorService.
func (s *Service) DeleteActor(subset game.ActorSubset) error {
	for _, id := range subset.IDs {
		if err := s.Process(redis.NewStringCmd("DEL", "actor", id)); err != nil {
			return err
		}
	}
	return nil
}

// ListActor is the tile38 implementation for ActorService.
func (s *Service) ListActor(subset game.ActorSubset) ([]game.Actor, error) {
	// var actors []game.Actor
	for _, id := range subset.IDs {
		cmd := redis.NewStringCmd("GET", "actor", id.String(), "POINT")
		if err := s.Process(cmd); err != nil {
			return nil, err
		}
		_, err := cmd.Result()
		if err != nil {
			return nil, err
		}
		// var actor game.Actor
		// actors = append(actors, actor)
	}
	return nil, errors.New("TO DEBUG")
}
