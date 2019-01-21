package srg

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/ability"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	templateKey = "atpl:"
)

// GetTemplate implemented with redis.
func (s *Store) GetTemplate(id ulid.ID) (ability.Template, error) {
	val, err := s.Get(templateKey + id.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return ability.Template{}, errors.Wrapf(err, "get template %s", id.String())
		}
		return ability.Template{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: templateKey, Index: id.String()},
			"get template %s",
			id.String(),
		)
	}

	var a ability.A
	if err := a.Unmarshal([]byte(val)); err != nil {
		return ability.Template{}, errors.Wrapf(err, "get template %s", id.String())
	}
	return ability.Template(a), nil
}

// SetTemplate implemented with redis.
func (s *Store) SetTemplate(t ability.Template) error {
	a := ability.A(t)
	raw, err := a.Marshal()
	if err != nil {
		return errors.Wrapf(err, "set template %s", t.ID.String())
	}
	return errors.Wrapf(s.Set(templateKey+t.ID.String(), raw, 0).Err(), "set template %s", t.ID.String())
}

// ListTemplate implemented with redis.
func (s *Store) ListTemplate() ([]ability.Template, error) {
	keys, err := s.Keys(templateKey + "*").Result()
	if err != nil {
		return nil, errors.Wrap(err, "list templates")
	}

	abilities := make([]ability.Template, len(keys))
	for i, key := range keys {
		val, err := s.Get(key).Result()
		if err != nil {
			return nil, errors.Wrap(err, "list templates")
		}

		if err := abilities[i].Unmarshal([]byte(val)); err != nil {
			return nil, errors.Wrap(err, "list templates")
		}
	}
	return abilities, nil
}
