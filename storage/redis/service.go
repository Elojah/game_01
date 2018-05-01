package redis

import (
	"github.com/elojah/redis"
)

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
