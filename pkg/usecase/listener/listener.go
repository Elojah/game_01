package listener

import (
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/pkg/errors"
)

// L wraps usecases for listener object.
type L struct {
	infra.QListenerMapper
	infra.ListenerMapper
	infra.CoreMapper
}

// New creates a new listener on a random core for id id.
func (l *L) New(id ulid.ID) (infra.Listener, error) {

	core, err := l.GetRandomCore(infra.CoreSubset{})
	if err != nil {
		return infra.Listener{}, errors.Wrap(err, "get random core")
	}
	listener := infra.Listener{ID: id, Action: infra.Open, Pool: core.ID}
	if err := l.PublishListener(listener, core.ID); err != nil {
		return infra.Listener{}, errors.Wrapf(err, "open listener %s on core %s", ulid.String(listener.ID), ulid.String(core.ID))
	}
	if err := l.SetListener(listener); err != nil {
		return infra.Listener{}, errors.Wrapf(err, "set listener %s", listener.ID)
	}
	return listener, nil
}

// Delete deletes a listener id on any pool.
func (l *L) Delete(id ulid.ID) error {
	listener, err := l.GetListener(infra.ListenerSubset{ID: id})
	if err != nil {
		return errors.Wrapf(err, "get listener %s", ulid.String(id))
	}
	listener.Action = infra.Close
	if err := l.PublishListener(listener, listener.Pool); err != nil {
		return errors.Wrapf(err, "close listener %s on pool %s", ulid.String(listener.ID), ulid.String(listener.Pool))
	}
	if err := l.DelListener(infra.ListenerSubset{ID: listener.ID}); err != nil {
		return errors.Wrapf(err, "delete listener %s", ulid.String(listener.ID))
	}
	return nil
}
