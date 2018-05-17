package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/dto"
	"github.com/elojah/mux"
)

func (h *handler) move(ctx context.Context, m dto.Move, token game.Token, ts time.Time) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	id := game.NewULID()
	source := game.ID(m.Source)
	target := game.ID(m.Target)
	go func() {
		if err := h.SendEvent(game.Event{
			ID:     id,
			Source: token.ID,
			TS:     ts,
			Action: game.MoveDone{
				Source:   source,
				Target:   target,
				Position: game.Vec3(m.Position),
			},
		}, source); err != nil {
			logger.Error().Err(err).Str("event", "unmarshalable").Msg("event rejected")
		}
	}()

	go func() {
		if err := h.SendEvent(game.Event{
			ID:     id,
			Source: token.ID,
			TS:     ts,
			Action: game.MoveReceived{
				Source:   source,
				Target:   target,
				Position: game.Vec3(m.Position),
			},
		}, target); err != nil {
			logger.Error().Err(err).Str("event", "unmarshalable").Msg("event rejected")
		}
	}()

	return nil
}
