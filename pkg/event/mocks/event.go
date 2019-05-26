package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

var _ event.Store = (*Store)(nil)
var _ event.QStore = (*QStore)(nil)
var _ event.App = (*App)(nil)

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
	PublishFunc    func(event.E, gulid.ID) error
	SubscribeFunc  func(gulid.ID) *infra.Subscription
	PublishCount   int32
	SubscribeCount int32
}

// Publish mocks Publish in event qstore.
func (s *QStore) Publish(e event.E, entityID gulid.ID) error {
	atomic.AddInt32(&s.PublishCount, 1)
	if s.PublishFunc == nil {
		return nil
	}
	return s.PublishFunc(e, entityID)
}

// Subscribe mocks Subscribe in event qstore.
func (s *QStore) Subscribe(entityID gulid.ID) *infra.Subscription {
	atomic.AddInt32(&s.SubscribeCount, 1)
	if s.SubscribeFunc == nil {
		return nil
	}
	return s.SubscribeFunc(entityID)
}

// NewQStore returns a event queue service mock ready for usage.
func NewQStore() *QStore {
	return &QStore{}
}

// App mocks event app.
type App struct {
	event.QStore
	event.Store
	event.TriggerStore

	CreateFunc  func(event.E, gulid.ID) error
	CancelFunc  func(event.E) error
	CreateCount int32
	CancelCount int32
}

// Create mocks Create in event app.
func (a *App) Create(e event.E, entityID gulid.ID) error {
	atomic.AddInt32(&a.CreateCount, 1)
	if a.CreateFunc == nil {
		return nil
	}
	return a.CreateFunc(e, entityID)
}

// Cancel mocks Cancel in event app.
func (a *App) Cancel(e event.E) error {
	atomic.AddInt32(&a.CancelCount, 1)
	if a.CancelFunc == nil {
		return nil
	}
	return a.CancelFunc(e)
}

// NewApp returns a app mock ready for usage.
func NewApp() *App {
	return &App{}
}
