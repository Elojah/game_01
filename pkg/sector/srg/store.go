package srg

import (
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/redis"
)

var _ sector.Store = (*Store)(nil)
var _ sector.EntitiesStore = (*Store)(nil)
var _ sector.StarterStore = (*Store)(nil)

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
