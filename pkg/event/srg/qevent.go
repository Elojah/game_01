package srg

import (
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/pkg/errors"
)

const (
	qeventKey = "qevent:"
)

// Publish implementation with redis pubsub.
func (s *Store) Publish(e event.E, id ulid.ID) error {
	raw, err := e.Marshal()
	if err != nil {
		return errors.Wrapf(err, "publish event %s to %s", e.ID.String(), id.String())
	}
	return errors.Wrapf(
		s.Service.Publish(qeventKey+id.String(), raw).Err(),
		"publish event %s to %s",
		e.ID.String(),
		id.String(),
	)
}

// Subscribe implementation with redis pubsub.
func (s *Store) Subscribe(id ulid.ID) *infra.Subscription {
	return s.Service.Subscribe(qeventKey + id.String())
}
