package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01/pkg/sector"
)

// Store mocks sector.Store.
type Store struct {
	SetSectorFunc  func(sector.S) error
	SetSectorCount int32
	GetSectorFunc  func(sector.Subset) (sector.S, error)
	GetSectorCount int32
}

// SetSector mocks sector.Store.
func (s *Store) SetSector(sec sector.S) error {
	atomic.AddInt32(&s.SetSectorCount, 1)
	if s.SetSectorFunc == nil {
		return nil
	}
	return s.SetSectorFunc(sec)
}

// GetSector mocks sector.Store.
func (s *Store) GetSector(subset sector.Subset) (sector.S, error) {
	atomic.AddInt32(&s.GetSectorCount, 1)
	if s.GetSectorFunc == nil {
		return sector.S{}, nil
	}
	return s.GetSectorFunc(subset)
}

// NewSectorStore returns a sec service mock ready for usage.
func NewSectorStore() *Store {
	return &Store{}
}
