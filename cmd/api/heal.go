package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/dto"
	"github.com/elojah/mux"
)

func (h *handler) heal(ctx context.Context, a dto.Heal, token game.Token, ts time.Time) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	go func() {
		err := h.SendEvent(game.Event{
			ID:     game.NewULID(),
			Source: token.ID,
			TS:     ts,
			Action: game.HealDone{
				Target: game.ID(a.Target),
				Amount: 10,
			},
		}, game.ID(a.Source))
		if err != nil {
			logger.Error().Err(err).Str("event", "unmarshalable").Msg("event rejected")
		}
	}()
	go func() {
		err := h.SendEvent(game.Event{
			ID:     game.NewULID(),
			Source: token.ID,
			TS:     ts,
			Action: game.HealReceived{
				Source: game.ID(a.Source),
				Amount: 10,
			},
		}, game.ID(a.Target))
		if err != nil {
			logger.Error().Err(err).Str("event", "unmarshalable").Msg("event rejected")
		}
	}()
	return nil
}
