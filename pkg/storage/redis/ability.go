package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/storage"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	abilityKey = "ability:"
)

// ListAbility implemented with redis.
func (s *Service) ListAbility(subset ability.Subset) ([]ability.A, error) {
	keys, err := s.Keys(abilityKey + ulid.String(subset.EntityID) + "*").Result()
	if err != nil {
		if err != redis.Nil {
			return nil, err
		}
		return nil, storage.ErrNotFound
	}

	abilities := make([]ability.A, len(keys))
	for i, key := range keys {
		val, err := s.Get(key).Result()
		if err != nil {
			return nil, err
		}

		var a storage.Ability
		if _, err := a.Unmarshal([]byte(val)); err != nil {
			return nil, err
		}
		abilities[i] = a.Domain()
	}
	return abilities, nil
}

// GetAbility implemented with redis.
func (s *Service) GetAbility(subset ability.Subset) (ability.A, error) {
	val, err := s.Get(abilityKey + ulid.String(subset.EntityID) + ":" + ulid.String(subset.ID)).Result()
	if err != nil {
		if err != redis.Nil {
			return ability.A{}, err
		}
		return ability.A{}, storage.ErrNotFound
	}

	var a storage.Ability
	if _, err := a.Unmarshal([]byte(val)); err != nil {
		return ability.A{}, err
	}
	return a.Domain(), nil
}

// SetAbility implemented with redis.
func (s *Service) SetAbility(a ability.A, entity ulid.ID) error {
	raw, err := storage.NewAbility(a).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(abilityKey+ulid.String(entity)+":"+ulid.String(a.ID), raw, 0).Err()
}
