package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01/pkg/sector"
)

// SectorService mocks sector.Service.
type SectorService struct {
	SetSectorFunc  func(sector.S) error
	SetSectorCount int32
	GetSectorFunc  func(sector.Subset) (sector.S, error)
	GetSectorCount int32
}

// SetSector mocks sector.Service.
func (m *SectorService) SetSector(s sector.S) error {
	atomic.AddInt32(&m.SetSectorCount, 1)
	if m.SetSectorFunc == nil {
		return nil
	}
	return m.SetSectorFunc(s)
}

// GetSector mocks sector.Service.
func (m *SectorService) GetSector(subset sector.Subset) (sector.S, error) {
	atomic.AddInt32(&m.GetSectorCount, 1)
	if m.GetSectorFunc == nil {
		return sector.S{}, nil
	}
	return m.GetSectorFunc(subset)
}

// NewSectorService returns a s service mock ready for usage.
func NewSectorService() *SectorService {
	return &SectorService{}
}
