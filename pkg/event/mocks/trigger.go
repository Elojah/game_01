package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// TriggerStore mocks an event trigger store.
type TriggerStore struct {
	SetTriggerFunc   func(t event.Trigger) error
	GetTriggerFunc   func(triggerID gulid.ID, entityID gulid.ID) (gulid.ID, error)
	ListTriggerFunc  func(triggerID gulid.ID) ([]event.Trigger, error)
	DelTriggerFunc   func(triggerID gulid.ID, entityID gulid.ID) error
	SetTriggerCount  int32
	GetTriggerCount  int32
	ListTriggerCount int32
	DelTriggerCount  int32
}

// SetTrigger mocks SetTrigger method in event trigger store.
func (s *TriggerStore) SetTrigger(t event.Trigger) error {
	atomic.AddInt32(&s.SetTriggerCount, 1)
	if s.SetTriggerFunc == nil {
		return nil
	}
	return s.SetTriggerFunc(t)
}

// GetTrigger mocks GetTrigger method in event trigger store.
func (s *TriggerStore) GetTrigger(triggerID gulid.ID, entityID gulid.ID) (gulid.ID, error) {
	atomic.AddInt32(&s.GetTriggerCount, 1)
	if s.GetTriggerFunc == nil {
		return gulid.Zero(), nil
	}
	return s.GetTriggerFunc(triggerID, entityID)
}

// ListTrigger mocks ListTrigger method in event trigger store.
func (s *TriggerStore) ListTrigger(triggerID gulid.ID) ([]event.Trigger, error) {
	atomic.AddInt32(&s.ListTriggerCount, 1)
	if s.ListTriggerFunc == nil {
		return nil, nil
	}
	return s.ListTriggerFunc(triggerID)
}

// DelTrigger mocks DelTrigger method in event trigger store.
func (s *TriggerStore) DelTrigger(triggerID gulid.ID, entityID gulid.ID) error {
	atomic.AddInt32(&s.DelTriggerCount, 1)
	if s.DelTriggerFunc == nil {
		return nil
	}
	return s.DelTriggerFunc(triggerID, entityID)
}

// NewTriggerStore returns a TriggerStore mock.
func NewTriggerStore() *TriggerStore {
	return &TriggerStore{}
}

// TriggerService is mock for event.TriggerService.
type TriggerService struct {
	SetFunc     func(event.E, gulid.ID) error
	CancelFunc  func(event.E) error
	SetCount    int32
	CancelCount int32
}

// Set mocks Set method of TriggerService.
func (s *TriggerService) Set(e event.E, id gulid.ID) error {
	atomic.AddInt32(&s.SetCount, 1)
	if s.SetFunc == nil {
		return nil
	}
	return s.SetFunc(e, id)
}

// Cancel mocks Cancel method of TriggerService.
func (s *TriggerService) Cancel(e event.E) error {
	atomic.AddInt32(&s.CancelCount, 1)
	if s.CancelFunc == nil {
		return nil
	}
	return s.CancelFunc(e)
}

// NewTriggerService returns a TriggerService mock.
func NewTriggerService() *TriggerService {
	return &TriggerService{}
}
