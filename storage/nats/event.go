package nats

import (
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/game_01/storage"
)

// PublishEvent implemented by NATS queue messaging.
func (s *Service) PublishEvent(e event.E, target ulid.ID) error {
	raw, err := storage.NewEvent(e).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Publish(target.String(), raw)
}
