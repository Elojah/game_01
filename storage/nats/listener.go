package nats

import (
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
