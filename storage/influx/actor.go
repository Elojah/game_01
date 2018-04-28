package influx

import (
	"github.com/elojah/game_01"
)

// CreateActor is the influxDB implementation of Actor.
func (s *Service) CreateActor([]game.Actor) error {
	return nil
}

// UpdateActor is the influxDB implementation of Actor.
func (s *Service) UpdateActor(game.ActorSubset, game.ActorPatch) error {
	return nil
}

// DeleteActor is the influxDB implementation of Actor.
func (s *Service) DeleteActor(game.ActorSubset) error {
	return nil
}

// ListActor is the influxDB implementation of Actor.
func (s *Service) ListActor(game.ActorSubset) ([]game.Actor, error) {
	return nil, nil
}
