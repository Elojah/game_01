package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01/pkg/entity"
)

// Store mocks entity.Store.
type Store struct {
	SetEntityFunc  func(entity.E, int64) error
	SetEntityCount int32
	GetEntityFunc  func(entity.Subset) (entity.E, error)
	GetEntityCount int32
	DelEntityFunc  func(entity.Subset) error
	DelEntityCount int32
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
func (s *Store) GetEntity(subset entity.Subset) (entity.E, error) {
	atomic.AddInt32(&s.GetEntityCount, 1)
	if s.GetEntityFunc == nil {
		return entity.E{}, nil
	}
	return s.GetEntityFunc(subset)
}

// DelEntity mocks entity.Store.
func (s *Store) DelEntity(subset entity.Subset) error {
	atomic.AddInt32(&s.DelEntityCount, 1)
	if s.DelEntityFunc == nil {
		return nil
	}
	return s.DelEntityFunc(subset)
}

// NewStore returns a entity service mock ready for usage.
func NewStore() *Store {
	return &Store{}
}
