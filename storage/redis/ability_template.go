package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/storage"
)

const (
	abilityTemplateKey = "a_template:"
)

// GetAbilityTemplate implemented with redis.
func (s *Service) GetAbilityTemplate(subset ability.TemplateSubset) (ability.Template, error) {
	val, err := s.Get(abilityTemplateKey + subset.Type.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return ability.Template{}, err
		}
		return ability.Template{}, storage.ErrNotFound
	}

	var a storage.Ability
	if _, err := a.Unmarshal([]byte(val)); err != nil {
		return ability.Template{}, err
	}
	return ability.Template(a.Domain()), nil
}

// SetAbilityTemplate implemented with redis.
func (s *Service) SetAbilityTemplate(template ability.Template) error {
	raw, err := storage.NewAbility(ability.A(template)).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(abilityTemplateKey+template.ID.String(), raw, 0).Err()
}

// ListAbilityTemplate implemented with redis.
func (s *Service) ListAbilityTemplate() ([]ability.Template, error) {
	keys, err := s.Keys(abilityTemplateKey + "*").Result()
	if err != nil {
		return nil, err
	}

	abilities := make([]ability.Template, len(keys))
	for i, key := range keys {
		val, err := s.Get(key).Result()
		if err != nil {
			return nil, err
		}

		var a storage.Ability
		if _, err := a.Unmarshal([]byte(val)); err != nil {
			return nil, err
		}
		abilities[i] = ability.Template(a.Domain())
	}
	return abilities, nil
}
