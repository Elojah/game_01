package nats

import (
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/storage"
)

// SendListener implemented by NATS queue messaging.
func (s *Service) SendListener(listener event.Listener) error {
	l := storage.NewListener(listener)
	raw, err := l.Marshal(nil)
	if err != nil {
		return err
	}
	return s.Publish(listener.Pool.String(), raw)
}
