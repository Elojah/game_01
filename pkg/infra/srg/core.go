package srg

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/errors"
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
			return infra.Core{}, err
		}
		return infra.Core{}, errors.ErrNotFound
	}

	return infra.Core{ID: ulid.MustParse(val)}, nil
}

// SetCore redis implementation.
func (s *Store) SetCore(core infra.Core) error {
	return s.SAdd(
		coreKey,
		core.ID.String(),
	).Err()
}

// DelCore redis implementation.
func (s *Store) DelCore(id ulid.ID) error {
	return s.SRem(coreKey, id.String()).Err()
}
