package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/mux"
)

func (h *handler) cast(ctx context.Context, msg event.DTO) error {

	logger := log.With().
		Str("packet", ctx.Value(mux.Key("packet")).(string)).
		Str("action", "cast").
		Logger()

	a := msg.Action.(event.Cast)
	source := ulid.ID(a.Source)
	targets := make([]ulid.ID, len(a.Targets))
	for i, target := range a.Targets {
		targets[i] = ulid.ID(target)
	}

	e := event.E{
		ID:     ulid.NewID(),
		Source: ulid.ID(msg.Token),
		TS:     time.Unix(0, msg.TS),
		Action: event.Cast{
			AbilityID: ulid.ID(a.AbilityID),
			Source:    source,
			Targets:   targets,
			Position:  geometry.Vec3(a.Position),
		},
	}

	logger = logger.With().Str("event", ulid.String(e.ID)).Logger()

	go func() {
		if err := h.PublishEvent(e, source); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
		logger.Info().Str("source", ulid.String(source)).Msg("send event")
	}()
	for _, target := range targets {
		go func(target ulid.ID) {
			if err := h.PublishEvent(e, target); err != nil {
				logger.Error().Err(err).Msg("event rejected")
			}
			logger.Info().Str("target", ulid.String(target)).Msg("send event")
		}(target)
	}
	return nil
}
