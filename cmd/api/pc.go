package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/dto"
	"github.com/elojah/mux"
)

func (h *handler) createPC(ctx context.Context, a dto.SetPC, token game.Token, ts time.Time) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	go func() {
		err := h.SendEvent(game.Event{
			ID:     game.NewULID(),
			Source: token.ID,
			TS:     ts,
			Action: game.SetPC{
				Type: game.EntityType(a.Type),
			},
		}, game.ID(token.ID))
		if err != nil {
			logger.Error().Err(err).Str("event", "unmarshalable").Msg("event rejected")
		}
	}()
	return nil
}
