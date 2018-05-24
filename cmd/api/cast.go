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
	targets := make([]game.ID, len(a.Targets))
	for i, target := range a.Targets {
		targets[i] = game.ID(target)
	}

	event := game.Event{
		ID:     game.NewULID(),
		Source: game.ID(msg.Token),
		TS:     time.Unix(0, msg.TS),
		Action: game.Cast{
			AbilityID: game.ID(a.AbilityID),
			Source:    source,
			Targets:   targets,
			Position:  game.Vec3(a.Position),
		},
	}

	go func() {
		if err := h.SendEvent(event, source); err != nil {
			logger.Error().Err(err).Str("event", event.ID.String()).Msg("event rejected")
		}
	}()
	for _, target := range targets {
		go func(target game.ID) {
			if err := h.SendEvent(event, target); err != nil {
				logger.Error().Err(err).Str("event", event.ID.String()).Msg("event rejected")
			}
		}(target)
	}
	return nil
}