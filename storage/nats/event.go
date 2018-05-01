package nats

import (
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
