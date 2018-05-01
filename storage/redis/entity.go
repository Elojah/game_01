package redis

import (
	"github.com/elojah/game_01"
)

const (
	entityKey = "entity:"
)

// ListEntity implemented with redis.
func (s *Service) ListEntity(subset game.EntitySubset) ([]game.Entity, error) {
	return nil, nil
}

// CreateEntity implemented with redis.
func (s *Service) CreateEntity(entities []game.Entity) error {
	return nil
}

// UpdateEntity implemented with redis.
func (s *Service) UpdateEntity(subset game.EntitySubset, patch game.EntityPatch) error {
	return nil
}

// DeleteEntity implemented with redis.
func (s *Service) DeleteEntity(subset game.EntitySubset) error {
	return nil
}
