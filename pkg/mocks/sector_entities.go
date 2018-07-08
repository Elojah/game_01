package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
)

// SectorEntitiesMapper mocks sector.EntitiesMapper.
type SectorEntitiesMapper struct {
	GetEntitiesFunc           func(sector.EntitiesSubset) (sector.Entities, error)
	GetEntitiesCount          int32
	AddEntityToSectorFunc     func(ulid.ID, ulid.ID) error
	AddEntityToSectorCount    int32
	RemoveEntityToSectorFunc  func(ulid.ID, ulid.ID) error
	RemoveEntityToSectorCount int32
}

// GetEntities mocks sector.EntitiesMapper.
func (m *SectorEntitiesMapper) GetEntities(subset sector.EntitiesSubset) (sector.Entities, error) {
	atomic.AddInt32(&m.GetEntitiesCount, 1)
	if m.GetEntitiesFunc == nil {
		return sector.Entities{}, nil
	}
	return m.GetEntitiesFunc(subset)
}

// AddEntityToSector mocks sector.EntitiesMapper.
func (m *SectorEntitiesMapper) AddEntityToSector(entityID ulid.ID, sectorID ulid.ID) error {
	atomic.AddInt32(&m.AddEntityToSectorCount, 1)
	if m.AddEntityToSectorFunc == nil {
		return nil
	}
	return m.AddEntityToSectorFunc(entityID, sectorID)
}

// RemoveEntityToSector mocks sector.EntitiesMapper.
func (m *SectorEntitiesMapper) RemoveEntityToSector(entityID ulid.ID, sectorID ulid.ID) error {
	atomic.AddInt32(&m.RemoveEntityToSectorCount, 1)
	if m.RemoveEntityToSectorFunc == nil {
		return nil
	}
	return m.RemoveEntityToSectorFunc(entityID, sectorID)
}

// NewSectorEntitiesMapper returns a event service mock ready for usage.
func NewSectorEntitiesMapper() *SectorEntitiesMapper {
	return &SectorEntitiesMapper{}
}
