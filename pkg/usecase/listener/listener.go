package listener

import (
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/pkg/errors"
)

// L wraps usecases for listener object.
type L struct {
	event.QListenerMapper
	event.ListenerMapper
	infra.CoreMapper
}

// New creates a new listener on a random core for id id.
func (l *L) New(id ulid.ID) (event.Listener, error) {

	core, err := l.GetRandomCore(infra.CoreSubset{})
	if err != nil {
		return event.Listener{}, errors.Wrap(err, "get random core")
	}
	listener := event.Listener{ID: id, Action: event.Open, Pool: core.ID}
	if err := l.PublishListener(listener, core.ID); err != nil {
		return event.Listener{}, errors.Wrapf(err, "open listener %s on core %s", listener.ID.String(), core.ID.String())
	}
	if err := l.SetListener(listener); err != nil {
		return event.Listener{}, errors.Wrapf(err, "set listener %s", listener.ID)
	}
	return listener, nil
}

// Delete deletes a listener id on any pool.
func (l *L) Delete(id ulid.ID) error {
	listener, err := l.GetListener(event.ListenerSubset{ID: id})
	if err != nil {
		return errors.Wrapf(err, "get listener %s", id.String())
	}
	listener.Action = event.Close
	if err := l.PublishListener(listener, listener.Pool); err != nil {
		return errors.Wrapf(err, "close listener %s on pool %s", listener.ID.String(), listener.Pool.String())
	}
	if err := l.DelListener(event.ListenerSubset{ID: listener.ID}); err != nil {
		return errors.Wrapf(err, "delete listener %s", listener.ID.String())
	}
	return nil
}
