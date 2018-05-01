package nats

import (
	"github.com/elojah/nats"
)

// Service implements event.
type Service struct {
	*nats.Service
}

// NewService returns a new game_01 nats Service.
func NewService(s *nats.Service) *Service {
	return &Service{
		Service: s,
	}
}
