package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

const (
	templateKey = "template:"
)

// GetTemplate implemented with redis.
func (s *Service) GetTemplate(subset game.TemplateSubset) (game.Template, error) {
	val, err := s.Get(templateKey + subset.Type).Result()
	if err != nil {
		if err != redis.Nil {
			return game.Template{}, err
		}
		return game.Template{}, storage.ErrNotFound
	}

	var entity storage.Entity
	if _, err := entity.Unmarshal([]byte(val)); err != nil {
		return game.Template{}, err
	}
	return game.Template(entity.Domain()), nil
}

// SetTemplate implemented with redis.
func (s *Service) SetTemplate(template game.Template) error {
	raw, err := storage.NewEntity(game.Entity(template)).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(templateKey+template.Type.String(), raw, 0).Err()
}
