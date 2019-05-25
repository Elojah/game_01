package app

import (
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"

	"github.com/pkg/errors"
)

var _ infra.RecurrerApp = (*RecurrerApp)(nil)

// RecurrerApp represents recurrer usecases.
type RecurrerApp struct {
	infra.QRecurrerStore
	infra.RecurrerStore
	infra.SyncStore
}

// Create creates a new recurrer on a random sync for id id.
func (app RecurrerApp) Create(entityID ulid.ID, tokenID ulid.ID) (infra.Recurrer, error) {

	// #Open recurrer on a random sync
	sy, err := app.SyncStore.FetchRandomSync()
	if err != nil {
		return infra.Recurrer{}, errors.Wrap(err, "create recurrer")
	}
	r := infra.Recurrer{
		EntityID: entityID,
		TokenID:  tokenID,
		Action:   infra.Open,
		Pool:     sy.ID,
	}
	if err := app.QRecurrerStore.PublishRecurrer(r, sy.ID); err != nil {
		return infra.Recurrer{}, errors.Wrap(err, "create recurrer")
	}

	// #Set recurrer with saved sync id
	if err := app.RecurrerStore.UpsertRecurrer(r); err != nil {
		return infra.Recurrer{}, errors.Wrap(err, "create recurrer")
	}
	return r, nil
}

// Erase deletes a recurrer id on any pool.
func (app RecurrerApp) Erase(id ulid.ID) error {

	// #Retrieve and close recurrer
	r, err := app.RecurrerStore.FetchRecurrer(id)
	if err != nil {
		return errors.Wrap(err, "erase recurrer")
	}
	r.Action = infra.Close
	if err := app.QRecurrerStore.PublishRecurrer(r, r.Pool); err != nil {
		return errors.Wrap(err, "erase recurrer")
	}

	// #Delete recurrer
	if err := app.RecurrerStore.RemoveRecurrer(r.TokenID); err != nil {
		return errors.Wrap(err, "erase recurrer")
	}
	return nil
}
