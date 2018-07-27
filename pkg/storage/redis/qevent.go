package redis

import (
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	qeventKey = "qevent:"
)

// PublishEvent implementation with redis pubsub.
func (s *Service) PublishEvent(e event.E, id ulid.ID) error {
	raw, err := e.Marshal(nil)
	if err != nil {
		return err
	}
	return s.Publish(qeventKey+id.String(), raw).Err()
}

// SubscribeEvent implementation with redis pubsub.
func (s *Service) SubscribeEvent(id ulid.ID) *infra.Subscription {
	return (*infra.Subscription)(s.Subscribe(qeventKey + id.String()))
}
