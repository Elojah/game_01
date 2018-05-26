package nats

import (
	"github.com/elojah/game_01"
	"github.com/elojah/nats"
)

var _ game.QEventMapper = (*Service)(nil)
var _ game.QListenerMapper = (*Service)(nil)
var _ game.SubscriptionMapper = (*Service)(nil)

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
