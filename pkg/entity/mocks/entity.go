package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/ulid"
)

// Store mocks entity.Store.
type Store struct {
	SetEntityFunc      func(entity.E, int64) error
	SetEntityCount     int32
	GetEntityFunc      func(ulid.ID, int64) (entity.E, error)
	GetEntityCount     int32
	DelEntityFunc      func(ulid.ID) error
	DelEntityCount     int32
	DelEntityByTSFunc  func(ulid.ID, int64) error
	DelEntityByTSCount int32
}

// SetEntity mocks entity.Store.
func (s *Store) SetEntity(e entity.E, ts int64) error {
	atomic.AddInt32(&s.SetEntityCount, 1)
	if s.SetEntityFunc == nil {
		return nil
	}
	return s.SetEntityFunc(e, ts)
}

// GetEntity mocks entity.Store.
func (s *Store) GetEntity(id ulid.ID, maxTS int64) (entity.E, error) {
	atomic.AddInt32(&s.GetEntityCount, 1)
	if s.GetEntityFunc == nil {
		return entity.E{}, nil
	}
	return s.GetEntityFunc(id, maxTS)
}

// DelEntity mocks entity.Store.
func (s *Store) DelEntity(id ulid.ID) error {
	atomic.AddInt32(&s.DelEntityCount, 1)
	if s.DelEntityFunc == nil {
		return nil
	}
	return s.DelEntityFunc(id)
}

// DelEntityByTS mocks entity.Store.
func (s *Store) DelEntityByTS(id ulid.ID, minTS int64) error {
	atomic.AddInt32(&s.DelEntityByTSCount, 1)
	if s.DelEntityByTSFunc == nil {
		return nil
	}
	return s.DelEntityByTSFunc(id, minTS)
}

// NewStore returns a entity service mock ready for usage.
func NewStore() *Store {
	return &Store{}
}
