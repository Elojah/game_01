package nats

import (
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/game_01/storage"
)

// PublishListener implemented by NATS queue messaging.
func (s *Service) PublishListener(listener event.Listener, target ulid.ID) error {
	l := storage.NewListener(listener)
	raw, err := l.Marshal(nil)
	if err != nil {
		return err
	}
	return s.Publish(target.String(), raw)
}
