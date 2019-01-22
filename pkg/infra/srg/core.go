package srg

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	coreKey = "core:"
)

// GetRandomCore redis implementation.
func (s *Store) GetRandomCore() (infra.Core, error) {
	val, err := s.SRandMember(coreKey).Result()
	if err != nil {
		if err != redis.Nil {
			return infra.Core{}, errors.Wrap(err, "get random core")
		}
		return infra.Core{}, errors.Wrap(
			gerrors.ErrNotFound{Store: coreKey, Index: "random"},
			"get random core",
		)
	}

	return infra.Core{ID: ulid.MustParse(val)}, nil
}

// SetCore redis implementation.
func (s *Store) SetCore(core infra.Core) error {
	return errors.Wrapf(s.SAdd(
		coreKey,
		core.ID.String(),
	).Err(), "set core %s", core.ID.String())
}

// DelCore redis implementation.
func (s *Store) DelCore(id ulid.ID) error {
	return errors.Wrapf(s.SRem(coreKey, id.String()).Err(), "del core %s", id.String())
}
