package srg

import (
	"github.com/elojah/game_01/pkg/item"
	"github.com/elojah/redis"
)

var _ item.Store = (*Store)(nil)
var _ item.LootStore = (*Store)(nil)

// Store implements token and item.
type Store struct {
	*redis.Service
}

// NewStore returns a new game_01 redis Store.
func NewStore(s *redis.Service) *Store {
	return &Store{
		Service: s,
	}
}
