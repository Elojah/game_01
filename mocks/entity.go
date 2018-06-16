package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01"
)

// EntityMapper mocks game.EntityMapper.
type EntityMapper struct {
	SetEntityFunc  func(game.Entity, int64) error
	SetEntityCount int32
	GetEntityFunc  func(game.EntitySubset) (game.Entity, error)
	GetEntityCount int32
}

// SetEntity mocks game.EntityMapper.
func (m *EntityMapper) SetEntity(event game.Entity, ts int64) error {
	atomic.AddInt32(&m.SetEntityCount, 1)
	if m.SetEntityFunc == nil {
		return nil
	}
	return m.SetEntityFunc(event, ts)
}

// GetEntity mocks game.EntityMapper.
func (m *EntityMapper) GetEntity(subset game.EntitySubset) (game.Entity, error) {
	atomic.AddInt32(&m.GetEntityCount, 1)
	if m.GetEntityFunc == nil {
		return game.Entity{}, nil
	}
	return m.GetEntityFunc(subset)
}

// NewEntityMapper returns a event service mock ready for usage.
func NewEntityMapper() *EntityMapper {
	return &EntityMapper{}
}
