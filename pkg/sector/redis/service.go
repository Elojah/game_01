package redis

import (
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/redis"
)

var _ sector.Store = (*Service)(nil)
var _ sector.EntitiesStore = (*Service)(nil)
var _ sector.StarterStore = (*Service)(nil)

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
