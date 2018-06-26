package nats

import "github.com/elojah/game_01/pkg/event"

// SetSubscription returns a subscription which follows events received in NATS queue.
func (s *Service) SetSubscription(subject string, consumer event.MsgHandler) (*event.Subscription, error) {
	return s.Subscribe(subject, consumer)
}
