package nats

import (
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/game_01/storage"
)

// SendRecurrer implemented by NATS queue messaging.
func (s *Service) SendRecurrer(recurrer event.Recurrer, target ulid.ID) error {
	l := storage.NewRecurrer(recurrer)
	raw, err := l.Marshal(nil)
	if err != nil {
		return err
	}
	return s.Publish(target.String(), raw)
}
