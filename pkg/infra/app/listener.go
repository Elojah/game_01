package app

import (
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"

	"github.com/pkg/errors"
)

// ListenerService represents listener usecases.
type ListenerService struct {
	InfraQListenerStore infra.QListenerStore
	InfraListenerStore  infra.ListenerStore
	InfraCoreStore      infra.CoreStore
}

// New creates a new listener on a random core for id id.
func (s *ListenerService) New(id ulid.ID) (infra.Listener, error) {

	// #Open listener on a random core
	c, err := s.InfraCoreStore.GetRandomCore(infra.CoreSubset{})
	if err != nil {
		return infra.Listener{}, errors.Wrap(err, "get random core")
	}
	l := infra.Listener{ID: id, Action: infra.Open, Pool: c.ID}
	if err := s.InfraQListenerStore.PublishListener(l, c.ID); err != nil {
		return infra.Listener{}, errors.Wrapf(err, "open listener %s on core %s", l.ID.String(), c.ID.String())
	}

	// #Set listener with saved core id
	if err := s.InfraListenerStore.SetListener(l); err != nil {
		return infra.Listener{}, errors.Wrapf(err, "set listener %s", l.ID)
	}
	return l, nil
}

// Remove deletes a listener id on any pool.
func (s *ListenerService) Remove(id ulid.ID) error {

	// #Retrieve and close listener
	l, err := s.InfraListenerStore.GetListener(infra.ListenerSubset{ID: id})
	if err != nil {
		return errors.Wrapf(err, "get listener %s", id.String())
	}
	l.Action = infra.Close
	if err := s.InfraQListenerStore.PublishListener(l, l.Pool); err != nil {
		return errors.Wrapf(err, "close listener %s on pool %s", l.ID.String(), l.Pool.String())
	}

	// #Delete listener
	if err := s.InfraListenerStore.DelListener(infra.ListenerSubset{ID: l.ID}); err != nil {
		return errors.Wrapf(err, "delete listener %s", l.ID.String())
	}
	return nil
}
