package nats

import (
	"github.com/nats-io/go-nats"

	"github.com/elojah/game_01"
)

// CreateSubscription returns a chan which follows events received in NATS queue.
func (s *Service) CreateSubscription(subject string, bufsize int) (game.Subscription, error) {
	ch := make(chan *nats.Msg, bufsize)
	sub, err := s.ChanSubscribe(subject, ch)
	return game.Subscription{
		Subscription: sub,
		Ch:           game.MsgChan(ch),
	}, err
}
