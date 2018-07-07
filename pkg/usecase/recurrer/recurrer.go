package recurrer

import (
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
)

// R wraps usecases for recurrer object.
type R struct {
	event.QRecurrerMapper
	event.RecurrerMapper
	infra.SyncMapper
}

// New creates a new recurrer on a random sync for id id.
func (r *R) New(entityID ulid.ID, tokenID ulid.ID) (event.Recurrer, error) {

	sync, err := r.GetRandomSync(infra.SyncSubset{})
	if err != nil {
		return event.Recurrer{}, errors.Wrap(err, "get random sync")
	}
	recurrer := event.Recurrer{
		EntityID: entityID,
		TokenID:  tokenID,
		Action:   event.Open,
		Pool:     sync.ID,
	}
	if err := r.PublishRecurrer(recurrer, sync.ID); err != nil {
		return event.Recurrer{}, errors.Wrapf(err, "open recurrer %s on pool %s", entityID.String(), sync.ID.String())
	}
	if err := r.SetRecurrer(recurrer); err != nil {
		return event.Recurrer{}, errors.Wrapf(err, "set recurrer", recurrer.EntityID.String())
	}
	return recurrer, nil
}

// Delete deletes a recurrer id on any pool.
func (r *R) Delete(id ulid.ID) error {
	recurrer, err := r.GetRecurrer(event.RecurrerSubset{TokenID: id})
	if err != nil {
		return errors.Wrapf(err, "get recurrer %s", id.String())
	}
	recurrer.Action = event.Close
	if err := r.PublishRecurrer(recurrer, recurrer.Pool); err != nil {
		return errors.Wrapf(err, "close recurrer %s on pool %s", recurrer.EntityID.String(), recurrer.Pool.String())
	}
	if err := r.DelRecurrer(event.RecurrerSubset{TokenID: recurrer.TokenID}); err != nil {
		return errors.Wrapf(err, "delete recurrer for token %s", recurrer.TokenID.String())
	}
	return nil
}
