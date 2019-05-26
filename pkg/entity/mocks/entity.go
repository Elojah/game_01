package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/sector"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

var _ entity.Store = (*Store)(nil)
var _ entity.App = (*App)(nil)

// Store mocks entity.Store.
type Store struct {
	UpsertFunc      func(entity.E, uint64) error
	FetchFunc       func(gulid.ID, uint64) (entity.E, error)
	RemoveFunc      func(gulid.ID) error
	RemoveByTSFunc  func(gulid.ID, uint64) error
	UpsertCount     int32
	FetchCount      int32
	RemoveCount     int32
	RemoveByTSCount int32
}

// Upsert mocks entity.Store.
func (s *Store) Upsert(e entity.E, ts uint64) error {
	atomic.AddInt32(&s.UpsertCount, 1)
	if s.UpsertFunc == nil {
		return nil
	}
	return s.UpsertFunc(e, ts)
}

// Fetch mocks entity.Store.
func (s *Store) Fetch(id gulid.ID, maxTS uint64) (entity.E, error) {
	atomic.AddInt32(&s.FetchCount, 1)
	if s.FetchFunc == nil {
		return entity.E{}, nil
	}
	return s.FetchFunc(id, maxTS)
}

// Remove mocks entity.Store.
func (s *Store) Remove(id gulid.ID) error {
	atomic.AddInt32(&s.RemoveCount, 1)
	if s.RemoveFunc == nil {
		return nil
	}
	return s.RemoveFunc(id)
}

// RemoveByTS mocks entity.Store.
func (s *Store) RemoveByTS(id gulid.ID, minTS uint64) error {
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

// App implementation of entity applications.
type App struct {
	entity.InventoryStore
	entity.MRInventoryStore
	entity.PCLeftStore
	entity.PCStore
	entity.PermissionStore
	entity.SpawnStore
	entity.Store
	entity.TemplateStore

	AbilityStore ability.Store

	SectorEntitiesStore sector.EntitiesStore

	Sequencer infra.SequencerApp

	DisconnectFunc                 func(gulid.ID) error
	ErasePCFunc                    func(gulid.ID, gulid.ID) error
	FetchMRInventoryFromCacheFunc  func(gulid.ID, gulid.ID) (entity.Inventory, error)
	UpsertMRInventoryWithCacheFunc func(gulid.ID, entity.Inventory) error
	CheckPermissionFunc            func(gulid.ID, gulid.ID) error

	DisconnectCount                 int32
	ErasePCCount                    int32
	FetchMRInventoryFromCacheCount  int32
	UpsertMRInventoryWithCacheCount int32
	CheckPermissionCount            int32
}

func (a *App) Disconnect(id gulid.ID) error {
	atomic.AddInt32(&a.DisconnectCount, 1)
	if a.DisconnectFunc == nil {
		return nil
	}
	return a.DisconnectFunc(id)
}

func (a *App) ErasePC(accountID gulid.ID, id gulid.ID) error {
	atomic.AddInt32(&a.ErasePCCount, 1)
	if a.ErasePCFunc == nil {
		return nil
	}
	return a.ErasePCFunc(accountID, id)
}

func (a *App) FetchMRInventoryFromCache(id gulid.ID, entityID gulid.ID) (entity.Inventory, error) {
	atomic.AddInt32(&a.FetchMRInventoryFromCacheCount, 1)
	if a.FetchMRInventoryFromCacheFunc == nil {
		return entity.Inventory{}, nil
	}
	return a.FetchMRInventoryFromCacheFunc(id, entityID)
}

func (a *App) UpsertMRInventoryWithCache(id gulid.ID, inv entity.Inventory) error {
	atomic.AddInt32(&a.UpsertMRInventoryWithCacheCount, 1)
	if a.UpsertMRInventoryWithCacheFunc == nil {
		return nil
	}
	return a.UpsertMRInventoryWithCacheFunc(id, inv)

}

func (a *App) CheckPermission(tokenID gulid.ID, id gulid.ID) error {
	atomic.AddInt32(&a.CheckPermissionCount, 1)
	if a.CheckPermissionFunc == nil {
		return nil
	}
	return a.CheckPermissionFunc(tokenID, id)

}
