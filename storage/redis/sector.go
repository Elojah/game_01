package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

const (
	sectorKey = "sector:"
)

// GetSector implemented with redis.
func (s *Service) GetSector(subset game.SectorSubset) (game.Sector, error) {
	val, err := s.Get(sectorKey + subset.ID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return game.Sector{}, err
		}
		return game.Sector{}, storage.ErrNotFound
	}

	var sector storage.Sector
	if _, err := sector.Unmarshal([]byte(val)); err != nil {
		return game.Sector{}, err
	}
	return sector.Domain(), nil
}

// SetSector implemented with redis.
func (s *Service) SetSector(sector game.Sector) error {
	raw, err := storage.NewSector(sector).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(sectorKey+sector.ID.String(), raw, 0).Err()
}
