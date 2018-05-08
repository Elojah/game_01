package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/dto"
	"github.com/elojah/mux"
)

func (h *handler) attack(ctx context.Context, a dto.Attack, ts time.Time) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	go func() {
		err := h.SendEvent(game.Event{
			ID: game.NewULID(),
			TS: ts,
			Action: game.DamageInflict{
				Target: game.ID(a.Target),
				// TODO yes it is a const random number
				Amount: 10,
			},
		}, game.ID(a.Source))
		if err != nil {
			logger.Error().Err(err).Str("event", "unmarshalable").Msg("event rejected")
		}
	}()
	go func() {
		err := h.SendEvent(game.Event{
			ID: game.NewULID(),
			TS: ts,
			Action: game.Damage{
				Source: game.ID(a.Source),
				// TODO yes it is a const random number
				Amount: 10,
			},
		}, game.ID(a.Target))
		if err != nil {
			logger.Error().Err(err).Str("event", "unmarshalable").Msg("event rejected")
		}
	}()
	return nil
}