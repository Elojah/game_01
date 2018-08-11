package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/storage"
)

const (
	sectorKey = "sector:"
)

// GetSector implemented with redis.
func (s *Service) GetSector(subset sector.Subset) (sector.S, error) {
	val, err := s.Get(sectorKey + subset.ID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return sector.S{}, err
		}
		return sector.S{}, storage.ErrNotFound
	}

	var sec sector.S
	if err := sec.Unmarshal([]byte(val)); err != nil {
		return sector.S{}, err
	}
	return sec, nil
}

// SetSector implemented with redis.
func (s *Service) SetSector(sec sector.S) error {
	raw, err := sec.Marshal()
	if err != nil {
		return err
	}
	return s.Set(sectorKey+sec.ID.String(), raw, 0).Err()
}
