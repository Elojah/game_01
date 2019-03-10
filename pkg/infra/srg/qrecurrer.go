package srg

import (
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/pkg/errors"
)

const (
	qrecurrerKey = "qrecurrer:"
)

// PublishRecurrer implementation with redis pubsub.
func (s *Store) PublishRecurrer(r infra.Recurrer, id ulid.ID) error {
	raw, err := r.Marshal()
	if err != nil {
		return errors.Wrapf(err, "publish recurrer for token %s entity %s", r.TokenID.String(), r.EntityID.String())
	}
	return errors.Wrapf(
		s.Publish(qrecurrerKey+id.String(), raw).Err(),
		"publish recurrer for token %s entity %s",
		r.TokenID.String(),
		r.EntityID.String(),
	)
}

// SubscribeRecurrer implementation with redis pubsub.
func (s *Store) SubscribeRecurrer(id ulid.ID) *infra.Subscription {
	return s.Subscribe(qrecurrerKey + id.String())
}
