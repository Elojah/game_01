package redis

import (
	"github.com/oklog/ulid"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/pkg/sector"
)

const (
	sectorEntitiesKey = "sector_entities:"
)

// GetEntities implemented with redis.
func (s *Service) GetEntities(subset sector.EntitiesSubset) (sector.Entities, error) {
	sectorEntities := sector.Entities{SectorID: subset.SectorID}
	cmd := s.SMembers(sectorEntitiesKey + subset.SectorID.String())
	vals, err := cmd.Result()
	if err != nil {
		return sector.Entities{}, err
	}
	sectorEntities.EntityIDs = make([]game.ID, len(vals))
	for i, val := range vals {
		sectorEntities.EntityIDs[i] = ulid.MustParse(val)
	}
	return sectorEntities, nil
}

// AddEntityToSector implemented with redis.
func (s *Service) AddEntityToSector(entityID game.ID, sectorID game.ID) error {
	return s.SAdd(sectorEntitiesKey+sectorID.String(), entityID.String()).Err()
}

// RemoveEntityToSector implemented with redis.
func (s *Service) RemoveEntityToSector(entityID game.ID, sectorID game.ID) error {
	return s.SRem(sectorEntitiesKey+sectorID.String(), entityID.String()).Err()
}
