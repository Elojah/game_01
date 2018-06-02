package redis

import (
	"github.com/oklog/ulid"

	"github.com/elojah/game_01"
)

const (
	sectorEntitiesKey = "sector_entities:"
)

// GetSectorEntities implemented with redis.
func (s *Service) GetSectorEntities(subset game.SectorEntitiesSubset) (game.SectorEntities, error) {
	sectorEntities := game.SectorEntities{SectorID: subset.SectorID}
	cmd := s.SMembers(sectorEntitiesKey + subset.SectorID.String())
	vals, err := cmd.Result()
	if err != nil {
		return game.SectorEntities{}, err
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
