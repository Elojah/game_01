package redis

import (
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	qrecurrerKey = "qrecurrer:"
)

// PublishRecurrer implementation with redis pubsub.
func (s *Service) PublishRecurrer(r event.Recurrer, id ulid.ID) error {
	raw, err := r.Marshal(nil)
	if err != nil {
		return err
	}
	return s.Publish(qrecurrerKey+ulid.String(id), raw).Err()
}

// SubscribeRecurrer implementation with redis pubsub.
func (s *Service) SubscribeRecurrer(id ulid.ID) *event.Subscription {
	return (*event.Subscription)(s.Subscribe(qrecurrerKey + ulid.String(id)))
}
