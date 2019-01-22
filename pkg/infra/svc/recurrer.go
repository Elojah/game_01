package svc

import (
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"

	"github.com/pkg/errors"
)

// RecurrerService represents recurrer usecases.
type RecurrerService struct {
	InfraQRecurrer infra.QRecurrerStore
	InfraRecurrer  infra.RecurrerStore
	InfraSync      infra.SyncStore
}

// New creates a new recurrer on a random sync for id id.
func (s RecurrerService) New(entityID ulid.ID, tokenID ulid.ID) (infra.Recurrer, error) {

	// #Open recurrer on a random sync
	sy, err := s.InfraSync.GetRandomSync()
	if err != nil {
		return infra.Recurrer{}, errors.Wrap(err, "new recurrer")
	}
	r := infra.Recurrer{
		EntityID: entityID,
		TokenID:  tokenID,
		Action:   infra.Open,
		Pool:     sy.ID,
	}
	if err := s.InfraQRecurrer.PublishRecurrer(r, sy.ID); err != nil {
		return infra.Recurrer{}, errors.Wrap(err, "new recurrer")
	}

	// #Set recurrer with saved sync id
	if err := s.InfraRecurrer.SetRecurrer(r); err != nil {
		return infra.Recurrer{}, errors.Wrap(err, "new recurrer")
	}
	return r, nil
}

// Remove deletes a recurrer id on any pool.
func (s RecurrerService) Remove(id ulid.ID) error {

	// #Retrieve and close recurrer
	r, err := s.InfraRecurrer.GetRecurrer(id)
	if err != nil {
		return errors.Wrap(err, "remove recurrer")
	}
	r.Action = infra.Close
	if err := s.InfraQRecurrer.PublishRecurrer(r, r.Pool); err != nil {
		return errors.Wrap(err, "remove recurrer")
	}

	// #Delete recurrer
	if err := s.InfraRecurrer.DelRecurrer(r.TokenID); err != nil {
		return errors.Wrap(err, "remove recurrer")
	}
	return nil
}
