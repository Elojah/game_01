package storage

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/errors"
)

const (
	templateKey = "atpl:"
)

// GetTemplate implemented with redis.
func (s *Service) GetTemplate(subset ability.TemplateSubset) (ability.Template, error) {
	val, err := s.Get(templateKey + subset.Type.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return ability.Template{}, err
		}
		return ability.Template{}, errors.ErrNotFound
	}

	var a ability.A
	if err := a.Unmarshal([]byte(val)); err != nil {
		return ability.Template{}, err
	}
	return ability.Template(a), nil
}

// SetTemplate implemented with redis.
func (s *Service) SetTemplate(t ability.Template) error {
	a := ability.A(t)
	raw, err := a.Marshal()
	if err != nil {
		return err
	}
	return s.Set(templateKey+t.ID.String(), raw, 0).Err()
}

// ListTemplate implemented with redis.
func (s *Service) ListTemplate() ([]ability.Template, error) {
	keys, err := s.Keys(templateKey + "*").Result()
	if err != nil {
		return nil, err
	}

	abilities := make([]ability.Template, len(keys))
	for i, key := range keys {
		val, err := s.Get(key).Result()
		if err != nil {
			return nil, err
		}

		if err := abilities[i].Unmarshal([]byte(val)); err != nil {
			return nil, err
		}
	}
	return abilities, nil
}
