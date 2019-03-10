package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
)

// EntitiesStore mocks sector.EntitiesStore.
type EntitiesStore struct {
	GetEntitiesFunc             func(ulid.ID) (sector.Entities, error)
	AddEntityToSectorFunc       func(ulid.ID, ulid.ID) error
	RemoveEntityFromSectorFunc  func(ulid.ID, ulid.ID) error
	GetEntitiesCount            int32
	AddEntityToSectorCount      int32
	RemoveEntityFromSectorCount int32
}

// GetEntities mocks sector.EntitiesStore.
func (s *EntitiesStore) GetEntities(sectorID ulid.ID) (sector.Entities, error) {
	atomic.AddInt32(&s.GetEntitiesCount, 1)
	if s.GetEntitiesFunc == nil {
		return sector.Entities{}, nil
	}
	return s.GetEntitiesFunc(sectorID)
}

// AddEntityToSector mocks sector.EntitiesStore.
func (s *EntitiesStore) AddEntityToSector(entityID ulid.ID, sectorID ulid.ID) error {
	atomic.AddInt32(&s.AddEntityToSectorCount, 1)
	if s.AddEntityToSectorFunc == nil {
		return nil
	}
	return s.AddEntityToSectorFunc(entityID, sectorID)
}

// RemoveEntityFromSector mocks sector.EntitiesStore.
func (s *EntitiesStore) RemoveEntityFromSector(entityID ulid.ID, sectorID ulid.ID) error {
	atomic.AddInt32(&s.RemoveEntityFromSectorCount, 1)
	if s.RemoveEntityFromSectorFunc == nil {
		return nil
	}
	return s.RemoveEntityFromSectorFunc(entityID, sectorID)
}

// NewEntitiesStore returns a event service mock ready for usage.
func NewEntitiesStore() *EntitiesStore {
	return &EntitiesStore{}
}
