package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// TriggerStore mocks an event trigger store.
type TriggerStore struct {
	UpsertTriggerFunc  func(t event.Trigger) error
	FetchTriggerFunc   func(triggerID gulid.ID, entityID gulid.ID) (gulid.ID, error)
	ListTriggerFunc    func(triggerID gulid.ID) ([]event.Trigger, error)
	RemoveTriggerFunc  func(triggerID gulid.ID, entityID gulid.ID) error
	UpsertTriggerCount int32
	FetchTriggerCount  int32
	ListTriggerCount   int32
	RemoveTriggerCount int32
}

// UpsertTrigger mocks UpsertTrigger method in event trigger store.
func (s *TriggerStore) UpsertTrigger(t event.Trigger) error {
	atomic.AddInt32(&s.UpsertTriggerCount, 1)
	if s.UpsertTriggerFunc == nil {
		return nil
	}
	return s.UpsertTriggerFunc(t)
}

// FetchTrigger mocks FetchTrigger method in event trigger store.
func (s *TriggerStore) FetchTrigger(triggerID gulid.ID, entityID gulid.ID) (gulid.ID, error) {
	atomic.AddInt32(&s.FetchTriggerCount, 1)
	if s.FetchTriggerFunc == nil {
		return gulid.Zero(), nil
	}
	return s.FetchTriggerFunc(triggerID, entityID)
}

// ListTrigger mocks ListTrigger method in event trigger store.
func (s *TriggerStore) ListTrigger(triggerID gulid.ID) ([]event.Trigger, error) {
	atomic.AddInt32(&s.ListTriggerCount, 1)
	if s.ListTriggerFunc == nil {
		return nil, nil
	}
	return s.ListTriggerFunc(triggerID)
}

// RemoveTrigger mocks RemoveTrigger method in event trigger store.
func (s *TriggerStore) RemoveTrigger(triggerID gulid.ID, entityID gulid.ID) error {
	atomic.AddInt32(&s.RemoveTriggerCount, 1)
	if s.RemoveTriggerFunc == nil {
		return nil
	}
	return s.RemoveTriggerFunc(triggerID, entityID)
}

// NewTriggerStore returns a TriggerStore mock.
func NewTriggerStore() *TriggerStore {
	return &TriggerStore{}
}

// TriggerService is mock for event.TriggerService.
type TriggerService struct {
	UpsertFunc  func(event.E, gulid.ID) error
	CancelFunc  func(event.E) error
	UpsertCount int32
	CancelCount int32
}

// Upsert mocks Upsert method of TriggerService.
func (s *TriggerService) Upsert(e event.E, id gulid.ID) error {
	atomic.AddInt32(&s.UpsertCount, 1)
	if s.UpsertFunc == nil {
		return nil
	}
	return s.UpsertFunc(e, id)
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
