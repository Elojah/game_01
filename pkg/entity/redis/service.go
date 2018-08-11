package redis

import (
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/redis"
)

var _ entity.Store = (*Service)(nil)
var _ entity.PermissionStore = (*Service)(nil)
var _ entity.TemplateStore = (*Service)(nil)
var _ entity.PCStore = (*Service)(nil)
var _ entity.PCLeftStore = (*Service)(nil)

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
