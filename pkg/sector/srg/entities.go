package srg

import (
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/pkg/errors"
)

const (
	sectorEntitiesKey = "sector_entities:"
)

// FetchEntities implemented with redis.
func (s *Store) FetchEntities(sectorID ulid.ID) (sector.Entities, error) {
	sectorEntities := sector.Entities{SectorID: sectorID}
	vals, err := s.SMembers(sectorEntitiesKey + sectorID.String()).Result()
	if err != nil {
		return sector.Entities{}, errors.Wrapf(err, "fetch entities for sector %s", sectorID.String())
	}
	sectorEntities.EntityIDs = make([]ulid.ID, len(vals))
	for i, val := range vals {
		sectorEntities.EntityIDs[i] = ulid.MustParse(val)
	}
	return sectorEntities, nil
}

// AddEntityToSector implemented with redis.
func (s *Store) AddEntityToSector(entityID ulid.ID, sectorID ulid.ID) error {
	return errors.Wrapf(
		s.SAdd(sectorEntitiesKey+sectorID.String(), entityID.String()).Err(),
		"add entity %s to sector %s",
		entityID.String(),
		sectorID.String(),
	)
}

// RemoveEntityFromSector implemented with redis.
func (s *Store) RemoveEntityFromSector(entityID ulid.ID, sectorID ulid.ID) error {
	return errors.Wrapf(
		s.SRem(sectorEntitiesKey+sectorID.String(), entityID.String()).Err(),
		"remove entity %s from sector %s",
		entityID.String(),
		sectorID.String(),
	)
}
