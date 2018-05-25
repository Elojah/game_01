package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

const (
	abilityKey = "ability:"
)

// ListAbility implemented with redis.
func (s *Service) ListAbility(subset game.AbilitySubset) ([]game.Ability, error) {
	keys, err := s.Keys(abilityKey + subset.EntityID.String() + "*").Result()
	if err != nil {
		if err != redis.Nil {
			return nil, err
		}
		return nil, storage.ErrNotFound
	}

	abilities := make([]game.Ability, len(keys))
	for i, key := range keys {
		val, err := s.Get(key).Result()
		if err != nil {
			return nil, err
		}

		var ability storage.Ability
		if _, err := ability.Unmarshal([]byte(val)); err != nil {
			return nil, err
		}
		abilities[i] = ability.Domain()
	}
	return abilities, nil
}

// GetAbility implemented with redis.
func (s *Service) GetAbility(subset game.AbilitySubset) (game.Ability, error) {
	val, err := s.Get(abilityKey + subset.EntityID.String() + ":" + subset.ID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return game.Ability{}, err
		}
		return game.Ability{}, storage.ErrNotFound
	}

	var ability storage.Ability
	if _, err := ability.Unmarshal([]byte(val)); err != nil {
		return game.Ability{}, err
	}
	return ability.Domain(), nil
}

// SetAbility implemented with redis.
func (s *Service) SetAbility(ability game.Ability, entity game.ID) error {
	raw, err := storage.NewAbility(ability).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(abilityKey+entity.String()+":"+ability.ID.String(), raw, 0).Err()
}
