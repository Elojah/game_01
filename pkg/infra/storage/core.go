package storage

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
func (s *Service) GetRandomCore(subset infra.CoreSubset) (infra.Core, error) {
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
func (s *Service) SetCore(core infra.Core) error {
	return s.SAdd(
		coreKey,
		core.ID.String(),
	).Err()
}

// DelCore redis implementation.
func (s *Service) DelCore(subset infra.CoreSubset) error {
	return s.SRem(
		coreKey,
		subset.ID.String(),
	).Err()
}
