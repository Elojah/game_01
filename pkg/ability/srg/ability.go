package srg

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/ability"
	gerrors "github.com/elojah/game_01/pkg/errors"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

const (
	aKey = "ability:"
)

// ListAbility implemented with redis.
func (s *Store) ListAbility(entityID gulid.ID) ([]ability.A, error) {
	keys, err := s.Keys(aKey + entityID.String() + "*").Result()
	if err != nil {
		if err != redis.Nil {
			return nil, errors.Wrapf(err, "list abilities for entity %s", entityID.String())
		}
		return nil, errors.Wrapf(
			gerrors.ErrNotFound{Store: aKey, Index: entityID.String()},
			"list abilities for entity %s",
			entityID.String(),
		)
	}

	as := make([]ability.A, len(keys))
	for i, key := range keys {
		val, err := s.Get(key).Result()
		if err != nil {
			return nil, errors.Wrapf(err, "list abilities for entity %s", entityID.String())
		}

		if err := as[i].Unmarshal([]byte(val)); err != nil {
			return nil, errors.Wrapf(err, "list abilities for entity %s", entityID.String())
		}
	}
	return as, nil
}

// GetAbility implemented with redis.
func (s *Store) GetAbility(entityID gulid.ID, id gulid.ID) (ability.A, error) {
	val, err := s.Get(aKey + entityID.String() + ":" + id.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return ability.A{}, errors.Wrapf(err, "get ability %s for entity %s", id.String(), entityID.String())
		}
		return ability.A{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: aKey, Index: entityID.String() + ":" + id.String()},
			"get ability %s for entity %s",
			id.String(),
			entityID.String(),
		)
	}

	var a ability.A
	if err := a.Unmarshal([]byte(val)); err != nil {
		return ability.A{}, errors.Wrapf(err, "get ability %s for entity %s", id.String(), entityID.String())
	}
	return a, nil
}

// SetAbility implemented with redis.
func (s *Store) SetAbility(a ability.A, entityID gulid.ID) error {
	raw, err := a.Marshal()
	if err != nil {
		return errors.Wrapf(err, "set ability %s for entity %s", a.ID.String(), entityID.String())
	}
	return errors.Wrapf(
		s.Set(aKey+entityID.String()+":"+a.ID.String(), raw, 0).Err(),
		"set ability %s for entity %s",
		a.ID.String(),
		entityID.String(),
	)
}

// DelAbility deletes an ability from an entity in redis.
func (s *Store) DelAbility(entityID gulid.ID, abilityID gulid.ID) error {
	return errors.Wrapf(
		s.Del(aKey+entityID.String()+":"+abilityID.String()).Err(),
		"del ability %s for entity %s",
		abilityID.String(),
		entityID.String(),
	)
}
