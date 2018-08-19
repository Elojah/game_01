package srg

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/sector"
)

const (
	sectorKey = "sector:"
)

// GetSector implemented with redis.
func (s *Store) GetSector(subset sector.Subset) (sector.S, error) {
	val, err := s.Get(sectorKey + subset.ID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return sector.S{}, err
		}
		return sector.S{}, errors.ErrNotFound
	}

	var sec sector.S
	if err := sec.Unmarshal([]byte(val)); err != nil {
		return sector.S{}, err
	}
	return sec, nil
}

// SetSector implemented with redis.
func (s *Store) SetSector(sec sector.S) error {
	raw, err := sec.Marshal()
	if err != nil {
		return err
	}
	return s.Set(sectorKey+sec.ID.String(), raw, 0).Err()
}
