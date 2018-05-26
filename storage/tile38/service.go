package tile38

import (
	"github.com/elojah/game_01"
	"github.com/elojah/redis"
)

var _ game.EntityPositionMapper = (*Service)(nil)

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
