package srg

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
func (s *Store) ListAbility(entityID ulid.ID) ([]ability.A, error) {
	keys, err := s.Keys(aKey + entityID.String() + "*").Result()
	if err != nil {
		if err != redis.Nil {
			return nil, err
		}
		return nil, errors.ErrNotFound{Store: aKey, Index: entityID.String()}
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
func (s *Store) GetAbility(entityID ulid.ID, id ulid.ID) (ability.A, error) {
	val, err := s.Get(aKey + entityID.String() + ":" + id.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return ability.A{}, err
		}
		return ability.A{}, errors.ErrNotFound{Store: aKey, Index: entityID.String() + ":" + id.String()}
	}

	var a ability.A
	if err := a.Unmarshal([]byte(val)); err != nil {
		return ability.A{}, err
	}
	return a, nil
}

// SetAbility implemented with redis.
func (s *Store) SetAbility(a ability.A, en ulid.ID) error {
	raw, err := a.Marshal()
	if err != nil {
		return err
	}
	return s.Set(aKey+en.String()+":"+a.ID.String(), raw, 0).Err()
}
