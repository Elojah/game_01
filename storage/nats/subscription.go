package nats

import (
	"github.com/elojah/game_01"
)

// SetSubscription returns a subscription which follows events received in NATS queue.
func (s *Service) SetSubscription(subject string, consumer game.MsgHandler) (*game.Subscription, error) {
	return s.Subscribe(subject, consumer)
}
