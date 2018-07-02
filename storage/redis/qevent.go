package redis

import (
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/game_01/storage"
)

const (
	qeventKey = "qevent:"
)

// PublishEvent implementation with redis pubsub.
func (s *Service) PublishEvent(e event.E, id ulid.ID) error {
	raw, err := storage.NewEvent(e).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Publish(qeventKey+id.String(), raw).Err()
}

// SubscribeEvent implementation with redis pubsub.
func (s *Service) SubscribeEvent(id ulid.ID) *event.Subscription {
	return (*event.Subscription)(s.Subscribe(qeventKey + id.String()))
}
