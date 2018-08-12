package storage

import (
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/redis"
)

var _ infra.CoreStore = (*Service)(nil)
var _ infra.ListenerStore = (*Service)(nil)
var _ infra.QListenerStore = (*Service)(nil)
var _ infra.RecurrerStore = (*Service)(nil)
var _ infra.QRecurrerStore = (*Service)(nil)
var _ infra.SyncStore = (*Service)(nil)

// Service implements token and entity.
type Service struct {
	*redis.Service
}

// NewService returns a new game_01 redis Service.
func NewService(s *redis.Service) *Service {
	return &Service{
		Service: s,
	}
}
