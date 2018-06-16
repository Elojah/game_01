package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01"
)

// SectorEntitiesMapper mocks game.SectorEntitiesMapper.
type SectorEntitiesMapper struct {
	GetSectorEntitiesFunc     func(game.SectorEntitiesSubset) (game.SectorEntities, error)
	GetSectorEntitiesCount    int32
	AddEntityToSectorFunc     func(game.ID, game.ID) error
	AddEntityToSectorCount    int32
	RemoveEntityToSectorFunc  func(game.ID, game.ID) error
	RemoveEntityToSectorCount int32
}

// GetSectorEntities mocks game.SectorEntitiesMapper.
func (m *SectorEntitiesMapper) GetSectorEntities(subset game.SectorEntitiesSubset) (game.SectorEntities, error) {
	atomic.AddInt32(&m.GetSectorEntitiesCount, 1)
	if m.GetSectorEntitiesFunc == nil {
		return game.SectorEntities{}, nil
	}
	return m.GetSectorEntitiesFunc(subset)
}

// AddEntityToSector mocks game.SectorEntitiesMapper.
func (m *SectorEntitiesMapper) AddEntityToSector(entityID game.ID, sectorID game.ID) error {
	atomic.AddInt32(&m.AddEntityToSectorCount, 1)
	if m.AddEntityToSectorFunc == nil {
		return nil
	}
	return m.AddEntityToSectorFunc(entityID, sectorID)
}

// RemoveEntityToSector mocks game.SectorEntitiesMapper.
func (m *SectorEntitiesMapper) RemoveEntityToSector(entityID game.ID, sectorID game.ID) error {
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
