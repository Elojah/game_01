package recurrer

import (
	"github.com/rs/zerolog/log"

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
	logger := log.With().
		Str("recurrer", "").
		Str("entity", entityID.String()).
		Str("token", tokenID.String()).
		Str("action", "new").
		Logger()

	sync, err := r.GetRandomSync(infra.SyncSubset{})
	if err != nil {
		logger.Error().Err(err).Msg("failed to get random sync")
		return event.Recurrer{}, err
	}
	recurrer := event.Recurrer{
		EntityID: entityID,
		TokenID:  tokenID,
		Action:   event.Open,
		Pool:     sync.ID,
	}
	logger = logger.With().Str("recurrer", recurrer.TokenID.String()).Logger()
	if err := r.PublishRecurrer(recurrer, sync.ID); err != nil {
		logger.Error().Err(err).Str("sync", sync.ID.String()).Msg("failed to publish recurrer")
		return event.Recurrer{}, err
	}
	if err := r.SetRecurrer(recurrer); err != nil {
		logger.Error().Err(err).Msg("failed to set recurrer")
		return event.Recurrer{}, err
	}
	return recurrer, nil
}

// Delete deletes a recurrer id on any pool.
func (r *R) Delete(id ulid.ID) error {
	logger := log.With().
		Str("recurrer", id.String()).
		Str("action", "delete").
		Logger()
	recurrer, err := r.GetRecurrer(event.RecurrerSubset{TokenID: id})
	if err != nil {
		logger.Error().Err(err).Msg("failed to get recurrer")
		return err
	}
	recurrer.Action = event.Close
	if err := r.PublishRecurrer(recurrer, recurrer.Pool); err != nil {
		logger.Error().Err(err).Msg("failed to publish recurrer")
		return err
	}
	if err := r.DelRecurrer(event.RecurrerSubset{TokenID: recurrer.TokenID}); err != nil {
		logger.Error().Err(err).Msg("failed to delete recurrer")
		return err
	}
	return nil
}
