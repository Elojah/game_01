package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

const (
	entityTemplateKey = "e_template:"
)

// GetEntityTemplate implemented with redis.
func (s *Service) GetEntityTemplate(subset game.EntityTemplateSubset) (game.EntityTemplate, error) {
	val, err := s.Get(entityTemplateKey + subset.Type.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return game.EntityTemplate{}, err
		}
		return game.EntityTemplate{}, storage.ErrNotFound
	}

	var entity storage.Entity
	if _, err := entity.Unmarshal([]byte(val)); err != nil {
		return game.EntityTemplate{}, err
	}
	return game.EntityTemplate(entity.Domain()), nil
}

// SetEntityTemplate implemented with redis.
func (s *Service) SetEntityTemplate(template game.EntityTemplate) error {
	raw, err := storage.NewEntity(game.Entity(template)).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(entityTemplateKey+template.Type.String(), raw, 0).Err()
}
