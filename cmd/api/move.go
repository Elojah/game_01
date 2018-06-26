package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/dto"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/mux"
)

func (h *handler) move(ctx context.Context, msg dto.Message) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	a := msg.Action.(dto.Move)
	source := game.ID(a.Source)
	target := game.ID(a.Target)
	e := event.E{
		ID:     game.NewID(),
		Source: game.ID(msg.Token),
		TS:     time.Unix(0, msg.TS),
		Action: event.Move{
			Source:   source,
			Target:   target,
			Position: game.Vec3(a.Position),
		},
	}

	go func() {
		if err := h.SendEvent(e, source); err != nil {
			logger.Error().Err(err).Str("event", e.ID.String()).Msg("event rejected")
		}
	}()

	go func() {
		if err := h.SendEvent(e, target); err != nil {
			logger.Error().Err(err).Str("event", e.ID.String()).Msg("event rejected")
		}
	}()

	return nil
}
