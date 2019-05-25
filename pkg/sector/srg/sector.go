package srg

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	sectorKey = "sector:"
)

// Fetch implemented with redis.
func (s *Store) Fetch(id ulid.ID) (sector.S, error) {
	val, err := s.Get(sectorKey + id.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return sector.S{}, errors.Wrapf(err, "fetch sector %s", id.String())
		}
		return sector.S{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: sectorKey, Index: id.String()},
			"fetch sector %s",
			id.String(),
		)
	}

	var sec sector.S
	if err := sec.Unmarshal([]byte(val)); err != nil {
		return sector.S{}, errors.Wrapf(err, "fetch sector %s", id.String())
	}
	return sec, nil
}

// Upsert implemented with redis.
func (s *Store) Upsert(sec sector.S) error {
	raw, err := sec.Marshal()
	if err != nil {
		return errors.Wrapf(err, "upsert sector %s", sec.ID.String())
	}
	return errors.Wrapf(
		s.Set(sectorKey+sec.ID.String(), raw, 0).Err(),
		"upsert sector %s",
		sec.ID.String(),
	)
}
