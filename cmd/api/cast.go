package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/dto"
	"github.com/elojah/mux"
)

func (h *handler) cast(ctx context.Context, msg dto.Message) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	a := msg.Action.(dto.Cast)
	source := game.ID(a.Source)
	target := game.ID(a.Target)

	event := game.Event{
		ID:     game.NewULID(),
		Source: game.ID(msg.Token),
		TS:     time.Unix(0, msg.TS),
		Action: game.Cast{
			AbilityID: game.ID(a.AbilityID),
			Source:    source,
			Target:    target,
			Position:  game.Vec3(a.Position),
		},
	}

	go func() {
		if err := h.SendEvent(event, source); err != nil {
			logger.Error().Err(err).Str("event", event.ID.String()).Msg("event rejected")
		}
	}()
	go func() {
		if err := h.SendEvent(event, target); err != nil {
			logger.Error().Err(err).Str("event", event.ID.String()).Msg("event rejected")
		}
	}()
	return nil
}
