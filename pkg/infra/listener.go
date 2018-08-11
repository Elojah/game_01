package infra

import (
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"

	"github.com/pkg/errors"
)

// QListenerStore handles send/receive methods for listeners.
type QListenerStore interface {
	PublishListener(Listener, ulid.ID) error
	SubscribeListener(ulid.ID) *Subscription
}

// ListenerStore handle listener data interactions.
type ListenerStore interface {
	SetListener(Listener) error
	GetListener(ListenerSubset) (Listener, error)
	DelListener(ListenerSubset) error
}

// ListenerSubset retrieves listener per ID.
type ListenerSubset struct {
	ID ulid.ID
}

// ListenerService represents listener usecases.
type ListenerService struct {
	QListenerStore QListenerStore
	ListenerStore  ListenerStore
	CoreStore      CoreStore
}

// New creates a new listener on a random core for id id.
func (s *ListenerService) New(id ulid.ID) (infra.Listener, error) {
	c, err := s.GetRandomCore(infra.CoreSubset{})
	if err != nil {
		return infra.Listener{}, errors.Wrap(err, "get random core")
	}
	l := infra.Listener{ID: id, Action: infra.Open, Pool: c.ID}
	if err := s.PublishListener(l, c.ID); err != nil {
		return infra.Listener{}, errors.Wrapf(err, "open listener %s on core %s", l.ID.String(), c.ID.String())
	}
	if err := s.SetListener(l); err != nil {
		return infra.Listener{}, errors.Wrapf(err, "set listener %s", l.ID)
	}
	return l, nil
}

// Remove deletes a listener id on any pool.
func (s *ListenerService) Remove(id ulid.ID) error {
	l, err := s.GetListener(infra.ListenerSubset{ID: id})
	if err != nil {
		return errors.Wrapf(err, "get listener %s", id.String())
	}
	l.Action = infra.Close
	if err := s.PublishListener(l, l.Pool); err != nil {
		return errors.Wrapf(err, "close listener %s on pool %s", l.ID.String(), l.Pool.String())
	}
	if err := s.DelListener(infra.ListenerSubset{ID: l.ID}); err != nil {
		return errors.Wrapf(err, "delete listener %s", l.ID.String())
	}
	return nil
}
