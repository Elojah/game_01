package srg

import (
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	qrecurrerKey = "qrecurrer:"
)

// PublishRecurrer implementation with redis pubsub.
func (s *Store) PublishRecurrer(r infra.Recurrer, id ulid.ID) error {
	raw, err := r.Marshal()
	if err != nil {
		return err
	}
	return s.Publish(qrecurrerKey+id.String(), raw).Err()
}

// SubscribeRecurrer implementation with redis pubsub.
func (s *Store) SubscribeRecurrer(id ulid.ID) *infra.Subscription {
	return (*infra.Subscription)(s.Subscribe(qrecurrerKey + id.String()))
}
