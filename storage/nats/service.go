package nats

import (
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/nats"
)

var _ event.QListenerMapper = (*Service)(nil)
var _ event.QRecurrerMapper = (*Service)(nil)
var _ event.SubscriptionMapper = (*Service)(nil)

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
