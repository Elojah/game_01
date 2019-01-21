package srg

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	templateKey = "etpl:"
)

// GetTemplate implemented with redis.
func (s *Store) GetTemplate(id ulid.ID) (entity.Template, error) {
	val, err := s.Get(templateKey + id.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.Template{}, err
		}
		return entity.Template{}, errors.ErrNotFound{Store: templateKey, Index: id.String()}
	}

	var e entity.E
	if err := e.Unmarshal([]byte(val)); err != nil {
		return entity.Template{}, err
	}
	return entity.Template(e), nil
}

// SetTemplate implemented with redis.
func (s *Store) SetTemplate(t entity.Template) error {
	raw, err := t.Marshal()
	if err != nil {
		return err
	}
	return s.Set(templateKey+t.ID.String(), raw, 0).Err()
}

// ListTemplate implemented with redis.
func (s *Store) ListTemplate() ([]entity.Template, error) {
	keys, err := s.Keys(templateKey + "*").Result()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Template, len(keys))
	for i, key := range keys {
		val, err := s.Get(key).Result()
		if err != nil {
			return nil, err
		}

		if err := entities[i].Unmarshal([]byte(val)); err != nil {
			return nil, err
		}
	}
	return entities, nil
}
