package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// Store mocks event.Store.
type Store struct {
	SetEventFunc   func(event.E, gulid.ID) error
	SetEventCount  int32
	GetEventFunc   func(gulid.ID, gulid.ID) (event.E, error)
	GetEventCount  int32
	ListEventFunc  func(gulid.ID, gulid.ID) ([]event.E, error)
	ListEventCount int32
	DelEventFunc   func(gulid.ID, gulid.ID) error
	DelEventCount  int32
}

// SetEvent mocks event.Store.
func (s *Store) SetEvent(e event.E, id gulid.ID) error {
	atomic.AddInt32(&s.SetEventCount, 1)
	if s.SetEventFunc == nil {
		return nil
	}
	return s.SetEventFunc(e, id)
}

// GetEvent mocks event.Store.
func (s *Store) GetEvent(id gulid.ID, entityID gulid.ID) (event.E, error) {
	atomic.AddInt32(&s.GetEventCount, 1)
	if s.GetEventFunc == nil {
		return event.E{}, nil
	}
	return s.GetEventFunc(id, entityID)
}

// ListEvent mocks event.Store.
func (s *Store) ListEvent(id gulid.ID, min gulid.ID) ([]event.E, error) {
	atomic.AddInt32(&s.ListEventCount, 1)
	if s.ListEventFunc == nil {
		return nil, nil
	}
	return s.ListEventFunc(id, min)
}

// DelEvent mocks event.Store.
func (s *Store) DelEvent(id gulid.ID, eID gulid.ID) error {
	atomic.AddInt32(&s.DelEventCount, 1)
	if s.DelEventFunc == nil {
		return nil
	}
	return s.DelEventFunc(id, eID)
}

// NewStore returns a event service mock ready for usage.
func NewStore() *Store {
	return &Store{}
}

// QStore mocks event qstore.
type QStore struct {
	PublishEventFunc    func(event.E, gulid.ID) error
	PublishEventCount   int32
	SubscribeEventFunc  func(gulid.ID) *infra.Subscription
	SubscribeEventCount int32
}

// PublishEvent mocks PublishEvent in event qstore.
func (s *QStore) PublishEvent(e event.E, entityID gulid.ID) error {
	atomic.AddInt32(&s.PublishEventCount, 1)
	if s.PublishEventFunc == nil {
		return nil
	}
	return s.PublishEventFunc(e, entityID)
}

// SubscribeEvent mocks SubscribeEvent in event qstore.
func (s *QStore) SubscribeEvent(entityID gulid.ID) *infra.Subscription {
	atomic.AddInt32(&s.SubscribeEventCount, 1)
	if s.SubscribeEventFunc == nil {
		return nil
	}
	return s.SubscribeEventFunc(entityID)
}

// NewQStore returns a event queue service mock ready for usage.
func NewQStore() *QStore {
	return &QStore{}
}
