package nats

import (
	"github.com/nats-io/go-nats"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

// SendEvent implemented by NATS queue messaging.
func (s *Service) SendEvent(event game.Event, target game.ID) error {
	raw, err := storage.NewEvent(event).Marshal(nil)
	if err != nil {
		return err
	}
	s.Publish(target.String(), raw)
	return nil
}

// ReceiveEvent returns a chan which follows events received in NATS queue.
func (s *Service) ReceiveEvent(subject string, bufsize int) (*nats.Subscription, chan *nats.Msg, error) {
	ch := make(chan *nats.Msg, bufsize)
	sub, err := s.ChanSubscribe(subject, ch)
	return sub, ch, err
}
