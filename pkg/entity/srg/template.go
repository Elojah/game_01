package srg

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	templateKey = "etpl:"
)

// FetchTemplate implemented with redis.
func (s *Store) FetchTemplate(id ulid.ID) (entity.Template, error) {
	val, err := s.Get(templateKey + id.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.Template{}, errors.Wrapf(err, "fetch template %s", id.String())
		}
		return entity.Template{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: templateKey, Index: id.String()},
			"fetch template %s",
			id.String(),
		)
	}

	var e entity.E
	if err := e.Unmarshal([]byte(val)); err != nil {
		return entity.Template{}, errors.Wrapf(err, "fetch template %s", id.String())
	}
	return e, nil
}

// InsertTemplate implemented with redis.
func (s *Store) InsertTemplate(t entity.Template) error {
	raw, err := t.Marshal()
	if err != nil {
		return errors.Wrapf(err, "insert template %s", t.ID.String())
	}
	return errors.Wrapf(
		s.Set(templateKey+t.ID.String(), raw, 0).Err(),
		"insert template %s",
		t.ID.String(),
	)
}

// ListTemplate implemented with redis.
func (s *Store) ListTemplate() ([]entity.Template, error) {
	keys, err := s.Keys(templateKey + "*").Result()
	if err != nil {
		return nil, errors.Wrap(err, "list templates")
	}

	entities := make([]entity.Template, len(keys))
	for i, key := range keys {
		val, err := s.Get(key).Result()
		if err != nil {
			return nil, errors.Wrap(err, "list templates")
		}

		if err := entities[i].Unmarshal([]byte(val)); err != nil {
			return nil, errors.Wrap(err, "list templates")
		}
	}
	return entities, nil
}
