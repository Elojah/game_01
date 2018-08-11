package redis

import (
	"github.com/elojah/game_01/pkg/entity/redis"
	"github.com/elojah/game_01/pkg/event"
)

var _ event.Store = (*Service)(nil)
var _ event.QStore = (*Service)(nil)

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
