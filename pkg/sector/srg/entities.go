package srg

import (
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	sectorEntitiesKey = "sector_entities:"
)

// GetEntities implemented with redis.
func (s *Store) GetEntities(sectorID ulid.ID) (sector.Entities, error) {
	sectorEntities := sector.Entities{SectorID: sectorID}
	vals, err := s.SMembers(sectorEntitiesKey + sectorID.String()).Result()
	if err != nil {
		return sector.Entities{}, err
	}
	sectorEntities.EntityIDs = make([]ulid.ID, len(vals))
	for i, val := range vals {
		sectorEntities.EntityIDs[i] = ulid.MustParse(val)
	}
	return sectorEntities, nil
}

// AddEntityToSector implemented with redis.
func (s *Store) AddEntityToSector(entityID ulid.ID, sectorID ulid.ID) error {
	return s.SAdd(sectorEntitiesKey+sectorID.String(), entityID.String()).Err()
}

// RemoveEntityFromSector implemented with redis.
func (s *Store) RemoveEntityFromSector(entityID ulid.ID, sectorID ulid.ID) error {
	return s.SRem(sectorEntitiesKey+sectorID.String(), entityID.String()).Err()
}
