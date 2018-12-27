package srg

import (
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/redis"
)

var _ event.Store = (*Store)(nil)
var _ event.QStore = (*Store)(nil)
var _ event.TriggerStore = (*Store)(nil)

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
