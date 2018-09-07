package srg

import (
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/redis"
)

var _ infra.CoreStore = (*Store)(nil)
var _ infra.SequencerStore = (*Store)(nil)
var _ infra.QSequencerStore = (*Store)(nil)
var _ infra.RecurrerStore = (*Store)(nil)
var _ infra.QRecurrerStore = (*Store)(nil)
var _ infra.SyncStore = (*Store)(nil)

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
