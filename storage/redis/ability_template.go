package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

const (
	abilityTemplateKey = "a_template:"
)

// GetAbilityTemplate implemented with redis.
func (s *Service) GetAbilityTemplate(subset game.AbilityTemplateSubset) (game.AbilityTemplate, error) {
	val, err := s.Get(abilityTemplateKey + subset.Type.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return game.AbilityTemplate{}, err
		}
		return game.AbilityTemplate{}, storage.ErrNotFound
	}

	var ability storage.Ability
	if _, err := ability.Unmarshal([]byte(val)); err != nil {
		return game.AbilityTemplate{}, err
	}
	return game.AbilityTemplate(ability.Domain()), nil
}

// SetAbilityTemplate implemented with redis.
func (s *Service) SetAbilityTemplate(template game.AbilityTemplate) error {
	raw, err := storage.NewAbility(game.Ability(template)).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(abilityTemplateKey+template.ID.String(), raw, 0).Err()
}

// ListAbilityTemplate implemented with redis.
func (s *Service) ListAbilityTemplate() ([]game.AbilityTemplate, error) {
	keys, err := s.Keys(abilityTemplateKey + "*").Result()
	if err != nil {
		return nil, err
	}

	entities := make([]game.AbilityTemplate, len(keys))
	for i, key := range keys {
		val, err := s.Get(key).Result()
		if err != nil {
			return nil, err
		}

		var entity storage.Ability
		if _, err := entity.Unmarshal([]byte(val)); err != nil {
			return nil, err
		}
		entities[i] = game.AbilityTemplate(entity.Domain())
	}
	return entities, nil
}
