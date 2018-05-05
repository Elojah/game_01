package nats

import (
	"github.com/nats-io/go-nats"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

// SendListener implemented by NATS queue messaging.
func (s *Service) SendListener(listener game.Listener, target game.ID) error {
	l := storage.NewListener(listener)
	raw, err := l.Marshal(nil)
	if err != nil {
		return err
	}
	return s.Publish(target.String(), raw)
}

// ReceiveListener returns a chan which follows events received in NATS queue.
func (s *Service) ReceiveListener(subject string, bufsize int) (*game.Subscription, game.MsgChan, error) {
	ch := make(chan *nats.Msg, bufsize)
	sub, err := s.ChanSubscribe(subject, ch)
	return (*game.Subscription)(sub), game.MsgChan(ch), err
}
