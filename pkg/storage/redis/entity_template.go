package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/storage"
)

const (
	entityTemplateKey = "e_template:"
)

// GetEntityTemplate implemented with redis.
func (s *Service) GetEntityTemplate(subset entity.TemplateSubset) (entity.Template, error) {
	val, err := s.Get(entityTemplateKey + subset.Type.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.Template{}, err
		}
		return entity.Template{}, storage.ErrNotFound
	}

	var e storage.Entity
	if _, err := e.Unmarshal([]byte(val)); err != nil {
		return entity.Template{}, err
	}
	return entity.Template(e.Domain()), nil
}

// SetEntityTemplate implemented with redis.
func (s *Service) SetEntityTemplate(template entity.Template) error {
	raw, err := storage.NewEntity(entity.E(template)).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(entityTemplateKey+template.ID.String(), raw, 0).Err()
}

// ListEntityTemplate implemented with redis.
func (s *Service) ListEntityTemplate() ([]entity.Template, error) {
	keys, err := s.Keys(entityTemplateKey + "*").Result()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Template, len(keys))
	for i, key := range keys {
		val, err := s.Get(key).Result()
		if err != nil {
			return nil, err
		}

		var e storage.Entity
		if _, err := e.Unmarshal([]byte(val)); err != nil {
			return nil, err
		}
		entities[i] = entity.Template(e.Domain())
	}
	return entities, nil
}
