package srg

import (
	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/redis"
)

var _ ability.Store = (*Store)(nil)
var _ ability.FeedbackStore = (*Store)(nil)
var _ ability.TemplateStore = (*Store)(nil)
var _ ability.StarterStore = (*Store)(nil)

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
