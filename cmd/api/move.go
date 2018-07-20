package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/dto"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/mux"
)

func (h *handler) move(ctx context.Context, msg dto.Event) error {

	logger := log.With().
		Str("packet", ctx.Value(mux.Key("packet")).(string)).
		Str("action", "move").
		Logger()

	a := msg.Action.(dto.Move)
	source := ulid.ID(a.Source)
	target := ulid.ID(a.Target)
	e := event.E{
		ID:     ulid.NewID(),
		Source: ulid.ID(msg.Token),
		TS:     time.Unix(0, msg.TS),
		Action: event.Move{
			Source:   source,
			Target:   target,
			Position: geometry.Vec3(a.Position),
		},
	}

	logger = logger.With().Str("event", ulid.String(e.ID)).Logger()

	go func() {
		if err := h.PublishEvent(e, source); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
		logger.Info().Str("source", ulid.String(source)).Msg("send event")
	}()

	go func() {
		if err := h.PublishEvent(e, target); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
		logger.Info().Str("target", ulid.String(target)).Msg("send event")
	}()

	return nil
}
