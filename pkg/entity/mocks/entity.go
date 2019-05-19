package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/ulid"
)

// Store mocks entity.Store.
type Store struct {
	InsertFunc      func(entity.E, uint64) error
	FetchFunc       func(ulid.ID, uint64) (entity.E, error)
	RemoveFunc      func(ulid.ID) error
	RemoveByTSFunc  func(ulid.ID, uint64) error
	InsertCount     int32
	FetchCount      int32
	RemoveCount     int32
	RemoveByTSCount int32
}

// Insert mocks entity.Store.
func (s *Store) Insert(e entity.E, ts uint64) error {
	atomic.AddInt32(&s.InsertCount, 1)
	if s.InsertFunc == nil {
		return nil
	}
	return s.InsertFunc(e, ts)
}

// Fetch mocks entity.Store.
func (s *Store) Fetch(id ulid.ID, maxTS uint64) (entity.E, error) {
	atomic.AddInt32(&s.FetchCount, 1)
	if s.FetchFunc == nil {
		return entity.E{}, nil
	}
	return s.FetchFunc(id, maxTS)
}

// Remove mocks entity.Store.
func (s *Store) Remove(id ulid.ID) error {
	atomic.AddInt32(&s.RemoveCount, 1)
	if s.RemoveFunc == nil {
		return nil
	}
	return s.RemoveFunc(id)
}

// RemoveByTS mocks entity.Store.
func (s *Store) RemoveByTS(id ulid.ID, minTS uint64) error {
	atomic.AddInt32(&s.RemoveByTSCount, 1)
	if s.RemoveByTSFunc == nil {
		return nil
	}
	return s.RemoveByTSFunc(id, minTS)
}

// NewStore returns a entity service mock ready for usage.
func NewStore() *Store {
	return &Store{}
}
