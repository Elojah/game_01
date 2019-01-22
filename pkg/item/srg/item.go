package srg

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
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
			return item.I{}, errors.Wrapf(err, "get item %s", id.String())
		}
		return item.I{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: itemKey, Index: id.String()},
			"get item %s",
			id.String(),
		)
	}

	var it item.I
	if err := it.Unmarshal([]byte(val)); err != nil {
		return item.I{}, errors.Wrapf(err, "get item %s", id.String())
	}
	return it, nil
}

// SetItem implemented with redis.
func (s *Store) SetItem(it item.I) error {
	raw, err := it.Marshal()
	if err != nil {
		return errors.Wrapf(err, "set item %s", it.ID.String())
	}
	return errors.Wrapf(
		s.Set(itemKey+it.ID.String(), raw, 0).Err(),
		"set item %s",
		it.ID.String(),
	)
}

// DelItem implemented with redis.
func (s *Store) DelItem(id ulid.ID) error {
	return errors.Wrapf(s.Del(itemKey+id.String()).Err(), "del item %s", id.String())
}
