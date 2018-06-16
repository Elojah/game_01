package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01"
)

// SectorMapper mocks game.SectorMapper.
type SectorMapper struct {
	SetSectorFunc  func(game.Sector) error
	SetSectorCount int32
	GetSectorFunc  func(game.SectorSubset) (game.Sector, error)
	GetSectorCount int32
}

// SetSector mocks game.SectorMapper.
func (m *SectorMapper) SetSector(event game.Sector) error {
	atomic.AddInt32(&m.SetSectorCount, 1)
	if m.SetSectorFunc == nil {
		return nil
	}
	return m.SetSectorFunc(event)
}

// GetSector mocks game.SectorMapper.
func (m *SectorMapper) GetSector(subset game.SectorSubset) (game.Sector, error) {
	atomic.AddInt32(&m.GetSectorCount, 1)
	if m.GetSectorFunc == nil {
		return game.Sector{}, nil
	}
	return m.GetSectorFunc(subset)
}

// NewSectorMapper returns a event service mock ready for usage.
func NewSectorMapper() *SectorMapper {
	return &SectorMapper{}
}
