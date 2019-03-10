package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
)

// Store mocks sector.Store.
type Store struct {
	SetSectorFunc  func(sector.S) error
	GetSectorFunc  func(ulid.ID) (sector.S, error)
	SetSectorCount int32
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
func (s *Store) GetSector(id ulid.ID) (sector.S, error) {
	atomic.AddInt32(&s.GetSectorCount, 1)
	if s.GetSectorFunc == nil {
		return sector.S{}, nil
	}
	return s.GetSectorFunc(id)
}

// NewSectorStore returns a sec service mock ready for usage.
func NewSectorStore() *Store {
	return &Store{}
}
