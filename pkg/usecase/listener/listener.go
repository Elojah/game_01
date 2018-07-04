package listener

import (
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
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
		return event.Listener{}, err
	}
	listener := event.Listener{ID: id, Action: event.Open, Pool: core.ID}
	if err := l.PublishListener(listener, core.ID); err != nil {
		return event.Listener{}, err
	}
	if err := l.SetListener(listener); err != nil {
		return event.Listener{}, err
	}
	return listener, nil
}

// Delete deletes a listener id on any pool.
func (l *L) Delete(id ulid.ID) error {
	listener, err := l.GetListener(event.ListenerSubset{ID: id})
	if err != nil {
		return err
	}
	listener.Action = event.Close
	if err := l.PublishListener(listener, listener.Pool); err != nil {
		return err
	}
	return nil
}
