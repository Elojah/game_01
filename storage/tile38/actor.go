package tile38

import (
	"github.com/elojah/game_01"
	"github.com/go-redis/redis"
)

// CreateActor is the tile38 implementation for ActorService.
func (s *Service) CreateActor(actors []game.Actor) error {
	query := `

	`
	return s.Process(redis.NewStringCmd("SET", "actor", "%s", "point", %f, %f))
}

// UpdateActor is the tile38 implementation for ActorService.
func (s *Service) UpdateActor(subset game.ActorSubset, patch game.ActorPatch) error {
	query := `
		SET actor %s point %f %f
	`
	return s.Process(redis.NewCmd(query))
}

// DeleteActor is the tile38 implementation for ActorService.
func (s *Service) DeleteActor(subset game.ActorSubset) error {
	query := `
		DEL actor %s
	`
	return s.Process(redis.NewCmd(query))
}

// ListActor is the tile38 implementation for ActorService.
func (s *Service) ListActor(subset game.ActorSubset) ([]game.Actor, error) {
	query := `
		SET actor %s point %f %f
	`
	return s.Process(redis.NewCmd(query))
}
