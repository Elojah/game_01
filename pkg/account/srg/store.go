package srg

import (
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/redis"
)

var _ account.Store = (*Store)(nil)
var _ account.TokenStore = (*Store)(nil)
var _ account.TokenHCStore = (*Store)(nil)

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
