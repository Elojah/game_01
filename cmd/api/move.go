package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/mux"
)

func (h *handler) move(ctx context.Context, msg event.DTO) error {

	logger := log.With().
		Str("packet", ctx.Value(mux.Key("packet")).(string)).
		Str("action", "move").
		Logger()

	a := msg.Action.(event.Move)
	source := ulid.ID(a.Source)
	target := ulid.ID(a.Target)
	e := event.E{
		ID:     ulid.NewID(),
		Source: msg.Token,
		TS:     time.Unix(0, msg.TS),
		Action: event.Move{
			Source:   source,
			Target:   target,
			Position: a.Position,
		},
	}

	logger = logger.With().Str("event", e.ID.String()).Logger()

	go func() {
		if err := h.PublishEvent(e, source); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
		logger.Info().Str("source", source.String()).Msg("send event")
	}()

	go func() {
		if err := h.PublishEvent(e, target); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
		logger.Info().Str("target", target.String()).Msg("send event")
	}()

	return nil
}
