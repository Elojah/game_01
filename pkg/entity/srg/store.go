package srg

import (
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/redis"
)

var _ entity.Store = (*Store)(nil)
var _ entity.InventoryStore = (*Store)(nil)
var _ entity.MRInventoryStore = (*Store)(nil)
var _ entity.PCStore = (*Store)(nil)
var _ entity.PCLeftStore = (*Store)(nil)
var _ entity.PermissionStore = (*Store)(nil)
var _ entity.SpawnStore = (*Store)(nil)
var _ entity.TemplateStore = (*Store)(nil)

// Store implements token and entity.
type Store struct {
	*redis.Service
}

// NewStore returns a new game_01 redis Store.
func NewStore(s *redis.Service) *Store {
	return &Store{
		Service: s,
	}
}
