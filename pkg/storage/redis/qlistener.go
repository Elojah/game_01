package redis

import (
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/storage"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	qlistenerKey = "qlistener:"
)

// PublishListener implementation with redis pubsub.
func (s *Service) PublishListener(l event.Listener, id ulid.ID) error {
	raw, err := storage.NewListener(l).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Publish(qlistenerKey+id.String(), raw).Err()
}

// SubscribeListener implementation with redis pubsub.
func (s *Service) SubscribeListener(id ulid.ID) *event.Subscription {
	return (*event.Subscription)(s.Subscribe(qlistenerKey + id.String()))
}
