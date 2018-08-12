package app

import (
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"

	"github.com/pkg/errors"
)

// RecurrerStore represents recurrer usecases.
type InfraRecurrerStore struct {
	InfraQRecurrerStore infra.QRecurrerStore
	InfraRecurrerStore  infra.RecurrerStore
	InfraSyncStore      infra.SyncStore
}

// New creates a new recurrer on a random sync for id id.
func (s *InfraRecurrerStore) New(entityID ulid.ID, tokenID ulid.ID) (infra.Recurrer, error) {

	// #Open recurrer on a random sync
	sy, err := s.InfraSyncStore.GetRandomSync(infra.SyncSubset{})
	if err != nil {
		return infra.Recurrer{}, errors.Wrap(err, "get random sync")
	}
	r := infra.Recurrer{
		EntityID: entityID,
		TokenID:  tokenID,
		Action:   infra.Open,
		Pool:     sy.ID,
	}
	if err := s.InfraQRecurrerStore.PublishRecurrer(r, sy.ID); err != nil {
		return infra.Recurrer{}, errors.Wrapf(err, "open recurrer %s on pool %s", entityID.String(), sy.ID.String())
	}

	// #Set recurrer with saved sync id
	if err := s.InfraRecurrerStore.SetRecurrer(r); err != nil {
		return infra.Recurrer{}, errors.Wrapf(err, "set recurrer %s", r.EntityID.String())
	}
	return r, nil
}

// Remove deletes a recurrer id on any pool.
func (s *InfraRecurrerStore) Remove(id ulid.ID) error {

	// #Retrieve and close recurrer
	r, err := s.InfraRecurrerStore.GetRecurrer(infra.RecurrerSubset{TokenID: id})
	if err != nil {
		return errors.Wrapf(err, "get recurrer %s", id.String())
	}
	r.Action = infra.Close
	if err := s.InfraQRecurrerStore.PublishRecurrer(r, r.Pool); err != nil {
		return errors.Wrapf(err, "close recurrer %s on pool %s", r.EntityID.String(), r.Pool.String())
	}

	// #Delete recurrer
	if err := s.InfraRecurrerStore.DelRecurrer(infra.RecurrerSubset{TokenID: r.TokenID}); err != nil {
		return errors.Wrapf(err, "delete recurrer for token %s", r.TokenID.String())
	}
	return nil
}
