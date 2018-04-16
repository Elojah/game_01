package tile38

import (
	"github.com/elojah/redis"
)

// Service is a redis service connected to tile38.
type Service struct {
	*redis.Service
}

// NewService returns a tile38 service based on redis argument.
func NewService(s *redis.Service) *Service {
	return &Service{
		Service: s,
	}
}
