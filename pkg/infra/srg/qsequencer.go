package srg

import (
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/pkg/errors"
)

const (
	qsequencerKey = "qsequencer:"
)

// PublishSequencer implementation with redis pubsub.
func (s *Store) PublishSequencer(seq infra.Sequencer, id ulid.ID) error {
	raw, err := seq.Marshal()
	if err != nil {
		return errors.Wrapf(err, "publish sequencer %s on pool %s", seq.ID.String(), id.String())
	}
	return errors.Wrapf(
		s.Publish(qsequencerKey+id.String(), raw).Err(),
		"publish sequencer %s on pool %s",
		seq.ID.String(),
		id.String(),
	)
}

// SubscribeSequencer implementation with redis pubsub.
func (s *Store) SubscribeSequencer(id ulid.ID) *infra.Subscription {
	return (*infra.Subscription)(s.Subscribe(qsequencerKey + id.String()))
}
