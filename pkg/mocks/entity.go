package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01/pkg/entity"
)

// EntityService mocks entity.Service.
type EntityService struct {
	SetEntityFunc  func(entity.E, int64) error
	SetEntityCount int32
	GetEntityFunc  func(entity.Subset) (entity.E, error)
	GetEntityCount int32
	DelEntityFunc  func(entity.Subset) error
	DelEntityCount int32
}

// SetEntity mocks entity.Service.
func (m *EntityService) SetEntity(e entity.E, ts int64) error {
	atomic.AddInt32(&m.SetEntityCount, 1)
	if m.SetEntityFunc == nil {
		return nil
	}
	return m.SetEntityFunc(e, ts)
}

// GetEntity mocks entity.Service.
func (m *EntityService) GetEntity(subset entity.Subset) (entity.E, error) {
	atomic.AddInt32(&m.GetEntityCount, 1)
	if m.GetEntityFunc == nil {
		return entity.E{}, nil
	}
	return m.GetEntityFunc(subset)
}

// DelEntity mocks entity.Service.
func (m *EntityService) DelEntity(subset entity.Subset) error {
	atomic.AddInt32(&m.DelEntityCount, 1)
	if m.DelEntityFunc == nil {
		return nil
	}
	return m.DelEntityFunc(subset)
}

// NewEntityService returns a entity service mock ready for usage.
func NewEntityService() *EntityService {
	return &EntityService{}
}
