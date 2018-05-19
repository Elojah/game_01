package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/dto"
	"github.com/elojah/mux"
)

func (h *handler) move(ctx context.Context, msg dto.Message) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	id := game.NewULID()
	a := msg.Action.(dto.Move)
	token := game.ID(msg.Token)
	ts := time.Unix(0, msg.TS)
	source := game.ID(a.Source)
	target := game.ID(a.Target)
	position := game.Vec3(a.Position)

	go func() {
		if err := h.SendEvent(game.Event{
			ID:     id,
			Source: token,
			TS:     ts,
			Action: game.MoveDone{
				Source:   source,
				Target:   target,
				Position: position,
			},
		}, source); err != nil {
			logger.Error().Err(err).Str("event", "unmarshalable").Msg("event rejected")
		}
	}()

	go func() {
		if err := h.SendEvent(game.Event{
			ID:     id,
			Source: token,
			TS:     ts,
			Action: game.MoveReceived{
				Source:   source,
				Target:   target,
				Position: position,
			},
		}, target); err != nil {
			logger.Error().Err(err).Str("event", "unmarshalable").Msg("event rejected")
		}
	}()

	return nil
}
