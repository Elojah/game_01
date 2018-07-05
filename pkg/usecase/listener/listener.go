package listener

import (
	"github.com/rs/zerolog/log"

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
	logger := log.With().
		Str("listener", id.String()).
		Str("action", "new").
		Logger()

	core, err := l.GetRandomCore(infra.CoreSubset{})
	if err != nil {
		logger.Error().Err(err).Msg("failed to get random core")
		return event.Listener{}, err
	}
	listener := event.Listener{ID: id, Action: event.Open, Pool: core.ID}
	if err := l.PublishListener(listener, core.ID); err != nil {
		logger.Error().Err(err).Str("core", core.ID.String()).Msg("failed to publish listener")
		return event.Listener{}, err
	}
	if err := l.SetListener(listener); err != nil {
		logger.Error().Err(err).Msg("failed to set listener")
		return event.Listener{}, err
	}
	return listener, nil
}

// Delete deletes a listener id on any pool.
func (l *L) Delete(id ulid.ID) error {
	logger := log.With().
		Str("listener", id.String()).
		Str("action", "delete").
		Logger()
	listener, err := l.GetListener(event.ListenerSubset{ID: id})
	if err != nil {
		logger.Error().Err(err).Msg("failed to get listener")
		return err
	}
	listener.Action = event.Close
	if err := l.PublishListener(listener, listener.Pool); err != nil {
		logger.Error().Err(err).Msg("failed to publish listener")
		return err
	}
	if err := l.DelListener(event.ListenerSubset{ID: listener.ID}); err != nil {
		logger.Error().Err(err).Msg("failed to delete listener")
		return err
	}
	return nil
}
