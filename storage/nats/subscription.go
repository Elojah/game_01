package nats

import (
	"github.com/elojah/game_01"
)

// CreateSubscription returns a subscription which follows events received in NATS queue.
func (s *Service) CreateSubscription(subject string, consumer game.MsgHandler) (*game.Subscription, error) {
	return s.Subscribe(subject, consumer)
}
