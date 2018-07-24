package recurrer

import (
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
)

// R wraps usecases for recurrer object.
type R struct {
	infra.QRecurrerMapper
	infra.RecurrerMapper
	infra.SyncMapper
}

// New creates a new recurrer on a random sync for id id.
func (r *R) New(entityID ulid.ID, tokenID ulid.ID) (infra.Recurrer, error) {

	sync, err := r.GetRandomSync(infra.SyncSubset{})
	if err != nil {
		return infra.Recurrer{}, errors.Wrap(err, "get random sync")
	}
	recurrer := infra.Recurrer{
		EntityID: entityID,
		TokenID:  tokenID,
		Action:   infra.Open,
		Pool:     sync.ID,
	}
	if err := r.PublishRecurrer(recurrer, sync.ID); err != nil {
		return infra.Recurrer{}, errors.Wrapf(err, "open recurrer %s on pool %s", ulid.String(entityID), ulid.String(sync.ID))
	}
	if err := r.SetRecurrer(recurrer); err != nil {
		return infra.Recurrer{}, errors.Wrapf(err, "set recurrer", ulid.String(recurrer.EntityID))
	}
	return recurrer, nil
}

// Delete deletes a recurrer id on any pool.
func (r *R) Delete(id ulid.ID) error {
	recurrer, err := r.GetRecurrer(infra.RecurrerSubset{TokenID: id})
	if err != nil {
		return errors.Wrapf(err, "get recurrer %s", ulid.String(id))
	}
	recurrer.Action = infra.Close
	if err := r.PublishRecurrer(recurrer, recurrer.Pool); err != nil {
		return errors.Wrapf(err, "close recurrer %s on pool %s", ulid.String(recurrer.EntityID), ulid.String(recurrer.Pool))
	}
	if err := r.DelRecurrer(infra.RecurrerSubset{TokenID: recurrer.TokenID}); err != nil {
		return errors.Wrapf(err, "delete recurrer for token %s", ulid.String(recurrer.TokenID))
	}
	return nil
}
