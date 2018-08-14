package srg

import (
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	qlistenerKey = "qlistener:"
)

// PublishListener implementation with redis pubsub.
func (s *Store) PublishListener(l infra.Listener, id ulid.ID) error {
	raw, err := l.Marshal()
	if err != nil {
		return err
	}
	return s.Publish(qlistenerKey+id.String(), raw).Err()
}

// SubscribeListener implementation with redis pubsub.
func (s *Store) SubscribeListener(id ulid.ID) *infra.Subscription {
	return (*infra.Subscription)(s.Subscribe(qlistenerKey + id.String()))
}
