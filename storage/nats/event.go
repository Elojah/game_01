package nats

import (
	"github.com/elojah/game_01"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/storage"
)

// SendEvent implemented by NATS queue messaging.
func (s *Service) SendEvent(e event.E, target game.ID) error {
	raw, err := storage.NewEvent(e).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Publish(target.String(), raw)
}
