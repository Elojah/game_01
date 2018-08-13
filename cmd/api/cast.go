package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/mux"
)

func (h *handler) cast(ctx context.Context, msg event.DTO) error {

	logger := log.With().
		Str("packet", ctx.Value(mux.Key("packet")).(string)).
		Str("action", "cast").
		Logger()

	a := msg.Action.GetCast()
	source := ulid.ID(a.Source)
	targets := make([]ulid.ID, len(a.Targets))
	for i, target := range a.Targets {
		targets[i] = ulid.ID(target)
	}

	e := event.E{
		ID:     ulid.NewID(),
		Source: msg.Token,
		TS:     time.Unix(0, msg.TS),
	}
	e.Action.SetValue(&event.Cast{
		AbilityID: a.AbilityID,
		Source:    source,
		Targets:   targets,
		Position:  a.Position,
	})

	logger = logger.With().Str("event", e.ID.String()).Logger()

	go func() {
		if err := h.PublishEvent(e, source); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
		logger.Info().Str("source", source.String()).Msg("send event")
	}()
	for _, target := range targets {
		go func(target ulid.ID) {
			if err := h.PublishEvent(e, target); err != nil {
				logger.Error().Err(err).Msg("event rejected")
			}
			logger.Info().Str("target", target.String()).Msg("send event")
		}(target)
	}
	return nil
}
