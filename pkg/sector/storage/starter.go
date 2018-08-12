package storage

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	starterKey = "starter:"
)

// GetRandomStarter redis implementation.
func (s *Service) GetRandomStarter(subset sector.StarterSubset) (sector.Starter, error) {
	val, err := s.SRandMember(starterKey).Result()
	if err != nil {
		if err != redis.Nil {
			return sector.Starter{}, err
		}
		return sector.Starter{}, errors.ErrNotFound
	}

	return sector.Starter{SectorID: ulid.MustParse(val)}, nil
}

// SetStarter redis implementation.
func (s *Service) SetStarter(starter sector.Starter) error {
	return s.SAdd(
		starterKey,
		starter.SectorID.String(),
	).Err()
}

// DelStarter redis implementation.
func (s *Service) DelStarter(subset sector.StarterSubset) error {
	return s.SRem(
		starterKey,
		subset.ID.String(),
	).Err()
}
