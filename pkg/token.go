package pkg

import (
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
)

// Token wraps use cases around token object.
type Token struct {
	account.TokenMapper

	EntityMapper entity.Mapper

	event.QRecurrerMapper
	event.QListenerMapper

	sector.EntitiesMapper
}

// DeleteToken closes a token and all entities/listener/sync associated.
func (t Token) DeleteToken(id ulid.ID) error {
	logger := log.With().
		Str("token", id.String()).
		Str("action", "close").
		Logger()
	go func() {
		// if err := t.SendListener(event.Listener{ID: id, Action: event.Close}); err != nil {
		// 	logger.Error().Err(err).Msg("failed to close listener")
		// }
	}()
	go func() {
		// if err := t.SendRecurrer(event.Recurrer{ID: id, Action: event.Close}); err != nil {
		// 	logger.Error().Err(err).Msg("failed to close recurrer")
		// }
	}()
	go func() {
		if err := t.EntityMapper.DelEntity(entity.Subset{Key: id.String()}); err != nil {
			logger.Error().Err(err).Msg("failed to delete entity")
		}
	}()
	return nil
}
