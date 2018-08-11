package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
)

// SectorEntitiesService mocks sector.EntitiesService.
type SectorEntitiesService struct {
	GetEntitiesFunc           func(sector.EntitiesSubset) (sector.Entities, error)
	GetEntitiesCount          int32
	AddEntityToSectorFunc     func(ulid.ID, ulid.ID) error
	AddEntityToSectorCount    int32
	RemoveEntityToSectorFunc  func(ulid.ID, ulid.ID) error
	RemoveEntityToSectorCount int32
}

// GetEntities mocks sector.EntitiesService.
func (m *SectorEntitiesService) GetEntities(subset sector.EntitiesSubset) (sector.Entities, error) {
	atomic.AddInt32(&m.GetEntitiesCount, 1)
	if m.GetEntitiesFunc == nil {
		return sector.Entities{}, nil
	}
	return m.GetEntitiesFunc(subset)
}

// AddEntityToSector mocks sector.EntitiesService.
func (m *SectorEntitiesService) AddEntityToSector(entityID ulid.ID, sectorID ulid.ID) error {
	atomic.AddInt32(&m.AddEntityToSectorCount, 1)
	if m.AddEntityToSectorFunc == nil {
		return nil
	}
	return m.AddEntityToSectorFunc(entityID, sectorID)
}

// RemoveEntityToSector mocks sector.EntitiesService.
func (m *SectorEntitiesService) RemoveEntityToSector(entityID ulid.ID, sectorID ulid.ID) error {
	atomic.AddInt32(&m.RemoveEntityToSectorCount, 1)
	if m.RemoveEntityToSectorFunc == nil {
		return nil
	}
	return m.RemoveEntityToSectorFunc(entityID, sectorID)
}

// NewSectorEntitiesService returns a event service mock ready for usage.
func NewSectorEntitiesService() *SectorEntitiesService {
	return &SectorEntitiesService{}
}
