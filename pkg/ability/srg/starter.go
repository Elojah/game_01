package srg

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/ability"
	gerrors "github.com/elojah/game_01/pkg/errors"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

const (
	starterKey = "abilityStarter:"
)

// SetStarter implements starter abilities store with redis.
func (s *Store) SetStarter(st ability.Starter) error {
	raw, err := st.Marshal()
	if err != nil {
		return errors.Wrapf(err, "set starter abilities for entity %s", st.EntityID.String())
	}
	return errors.Wrapf(
		s.Set(starterKey+st.EntityID.String(), raw, 0).Err(),
		"set starter abilities for entity %s",
		st.EntityID.String(),
	)
}

// GetStarter implements starter abilities store with redis.
func (s *Store) GetStarter(entityID gulid.ID) (ability.Starter, error) {
	val, err := s.Get(starterKey + entityID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return ability.Starter{}, errors.Wrapf(err, "get starter abilities for entity %s", entityID.String())
		}
		return ability.Starter{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: starterKey, Index: entityID.String()},
			"get starter abilities for entity %s",
			entityID.String(),
		)
	}

	var st ability.Starter
	if err := st.Unmarshal([]byte(val)); err != nil {
		return ability.Starter{}, errors.Wrapf(err, "get starter abilities for entity %s", entityID.String())
	}

	return st, nil
}
