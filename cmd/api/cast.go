package main

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/mux"
)

func (h *handler) cast(ctx context.Context, msg event.DTO) error {

	logger := log.With().
		Str("packet", ctx.Value(mux.Key("packet")).(string)).
		Str("action", "cast").
		Logger()

	cast := msg.Query.GetCast()
	e := event.E{
		ID:    msg.ID,
		Token: msg.Token,
		Action: event.Action{
			CastSource: &event.CastSource{
				AbilityID: cast.AbilityID,
				Targets:   cast.Targets,
			},
		},
	}

	logger = logger.With().Str("event", e.ID.String()).Logger()

	go func() {
		if err := h.PublishEvent(e, cast.Source); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
		logger.Info().Str("source", cast.Source.String()).Msg("send event")
	}()
	return nil
}
