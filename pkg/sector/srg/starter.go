package srg

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	starterKey = "starter:"
)

// GetRandomStarter redis implementation.
func (s *Store) GetRandomStarter() (sector.Starter, error) {
	val, err := s.SRandMember(starterKey).Result()
	if err != nil {
		if err != redis.Nil {
			return sector.Starter{}, errors.Wrap(err, "get random starter")
		}
		return sector.Starter{}, errors.Wrap(
			gerrors.ErrNotFound{Store: starterKey, Index: "random"},
			"get random starter",
		)
	}

	return sector.Starter{SectorID: ulid.MustParse(val)}, nil
}

// SetStarter redis implementation.
func (s *Store) SetStarter(starter sector.Starter) error {
	return errors.Wrapf(s.SAdd(
		starterKey,
		starter.SectorID.String(),
	).Err(), "set starter %s", starter.SectorID.String())
}

// DelStarter redis implementation.
func (s *Store) DelStarter(id ulid.ID) error {
	return errors.Wrapf(s.SRem(
		starterKey,
		id.String(),
	).Err(), "del starter %s", id.String())
}
