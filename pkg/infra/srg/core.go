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

// FetchRandomCore redis implementation.
func (s *Store) FetchRandomCore() (infra.Core, error) {
	val, err := s.SRandMember(coreKey).Result()
	if err != nil {
		if err != redis.Nil {
			return infra.Core{}, errors.Wrap(err, "fetch random core")
		}
		return infra.Core{}, errors.Wrap(
			gerrors.ErrNotFound{Store: coreKey, Index: "random"},
			"fetch random core",
		)
	}

	return infra.Core{ID: ulid.MustParse(val)}, nil
}

// UpsertCore redis implementation.
func (s *Store) UpsertCore(core infra.Core) error {
	return errors.Wrapf(s.SAdd(
		coreKey,
		core.ID.String(),
	).Err(), "upsert core %s", core.ID.String())
}

// RemoveCore redis implementation.
func (s *Store) RemoveCore(id ulid.ID) error {
	return errors.Wrapf(s.SRem(coreKey, id.String()).Err(), "remove core %s", id.String())
}
