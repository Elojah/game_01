package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/dto"
	"github.com/elojah/mux"
)

func (h *handler) attack(ctx context.Context, a dto.Attack, token game.Token, ts time.Time) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	id := game.NewULID()
	source := game.ID(a.Source)
	target := game.ID(a.Target)
	go func() {
		err := h.SendEvent(game.Event{
			ID:     id,
			Source: token.ID,
			TS:     ts,
			Action: game.AttackDone{
				Source: source,
				Target: target,
			},
		}, source)
		if err != nil {
			logger.Error().Err(err).Str("event", "unmarshalable").Msg("event rejected")
		}
	}()
	go func() {
		err := h.SendEvent(game.Event{
			ID:     id,
			Source: token.ID,
			TS:     ts,
			Action: game.AttackReceived{
				Source: source,
				Target: target,
			},
		}, target)
		if err != nil {
			logger.Error().Err(err).Str("event", "unmarshalable").Msg("event rejected")
		}
	}()
	return nil
}
