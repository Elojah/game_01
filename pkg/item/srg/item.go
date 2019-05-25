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

// FetchItem implemented with redis.
func (s *Store) FetchItem(id ulid.ID) (item.I, error) {
	val, err := s.Get(itemKey + id.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return item.I{}, errors.Wrapf(err, "fetch item %s", id.String())
		}
		return item.I{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: itemKey, Index: id.String()},
			"fetch item %s",
			id.String(),
		)
	}

	var it item.I
	if err := it.Unmarshal([]byte(val)); err != nil {
		return item.I{}, errors.Wrapf(err, "fetch item %s", id.String())
	}
	return it, nil
}

// UpsertItem implemented with redis.
func (s *Store) UpsertItem(it item.I) error {
	raw, err := it.Marshal()
	if err != nil {
		return errors.Wrapf(err, "upsert item %s", it.ID.String())
	}
	return errors.Wrapf(
		s.Set(itemKey+it.ID.String(), raw, 0).Err(),
		"upsert item %s",
		it.ID.String(),
	)
}

// RemoveItem implemented with redis.
func (s *Store) RemoveItem(id ulid.ID) error {
	return errors.Wrapf(s.Del(itemKey+id.String()).Err(), "remove item %s", id.String())
}
