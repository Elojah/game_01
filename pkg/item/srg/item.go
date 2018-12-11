package srg

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/item"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	itemKey = "item:"
)

// GetItem implemented with redis.
func (s *Store) GetItem(id ulid.ID) (item.I, error) {
	val, err := s.Get(itemKey + id.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return item.I{}, err
		}
		return item.I{}, errors.ErrNotFound
	}

	var it item.I
	if err := it.Unmarshal([]byte(val)); err != nil {
		return item.I{}, err
	}
	return it, nil
}

// SetItem implemented with redis.
func (s *Store) SetItem(it item.I) error {
	raw, err := it.Marshal()
	if err != nil {
		return err
	}
	return s.Set(itemKey+it.ID.String(), raw, 0).Err()
}

// DelItem implemented with redis.
func (s *Store) DelItem(id ulid.ID) error {
	return s.Del(itemKey + id.String()).Err()
}
