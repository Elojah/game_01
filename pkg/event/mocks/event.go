package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// Store mocks event.Store.
type Store struct {
	UpsertFunc  func(event.E, gulid.ID) error
	FetchFunc   func(gulid.ID, gulid.ID) (event.E, error)
	ListFunc    func(gulid.ID, gulid.ID) ([]event.E, error)
	RemoveFunc  func(gulid.ID, gulid.ID) error
	UpsertCount int32
	FetchCount  int32
	ListCount   int32
	RemoveCount int32
}

// Upsert mocks event.Store.
func (s *Store) Upsert(e event.E, id gulid.ID) error {
	atomic.AddInt32(&s.UpsertCount, 1)
	if s.UpsertFunc == nil {
		return nil
	}
	return s.UpsertFunc(e, id)
}

// Fetch mocks event.Store.
func (s *Store) Fetch(id gulid.ID, entityID gulid.ID) (event.E, error) {
	atomic.AddInt32(&s.FetchCount, 1)
	if s.FetchFunc == nil {
		return event.E{}, nil
	}
	return s.FetchFunc(id, entityID)
}

// List mocks event.Store.
func (s *Store) List(id gulid.ID, min gulid.ID) ([]event.E, error) {
	atomic.AddInt32(&s.ListCount, 1)
	if s.ListFunc == nil {
		return nil, nil
	}
	return s.ListFunc(id, min)
}

// Remove mocks event.Store.
func (s *Store) Remove(id gulid.ID, eID gulid.ID) error {
	atomic.AddInt32(&s.RemoveCount, 1)
	if s.RemoveFunc == nil {
		return nil
	}
	return s.RemoveFunc(id, eID)
}

// NewStore returns a event service mock ready for usage.
func NewStore() *Store {
	return &Store{}
}

// QStore mocks event qstore.
type QStore struct {
	PublishEventFunc    func(event.E, gulid.ID) error
	SubscribeEventFunc  func(gulid.ID) *infra.Subscription
	PublishEventCount   int32
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
