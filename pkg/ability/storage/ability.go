package storage

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	aKey = "ability:"
)

// ListAbility implemented with redis.
func (s *Service) ListAbility(subset ability.Subset) ([]ability.A, error) {
	keys, err := s.Keys(aKey + subset.EntityID.String() + "*").Result()
	if err != nil {
		if err != redis.Nil {
			return nil, err
		}
		return nil, errors.ErrNotFound
	}

	as := make([]ability.A, len(keys))
	for i, key := range keys {
		val, err := s.Get(key).Result()
		if err != nil {
			return nil, err
		}

		if err := as[i].Unmarshal([]byte(val)); err != nil {
			return nil, err
		}
	}
	return as, nil
}

// GetAbility implemented with redis.
func (s *Service) GetAbility(subset ability.Subset) (ability.A, error) {
	val, err := s.Get(aKey + subset.EntityID.String() + ":" + subset.ID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return ability.A{}, err
		}
		return ability.A{}, errors.ErrNotFound
	}

	var a ability.A
	if err := a.Unmarshal([]byte(val)); err != nil {
		return ability.A{}, err
	}
	return a, nil
}

// SetAbility implemented with redis.
func (s *Service) SetAbility(a ability.A, en ulid.ID) error {
	raw, err := a.Marshal()
	if err != nil {
		return err
	}
	return s.Set(aKey+en.String()+":"+a.ID.String(), raw, 0).Err()
}
