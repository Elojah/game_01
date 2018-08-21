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
	e := event.E{
		ID:     ulid.NewID(),
		Source: msg.Token,
		TS:     time.Unix(0, msg.TS),
		Action: event.Action{
			Cast: a,
		},
	}

	logger = logger.With().Str("event", e.ID.String()).Logger()

	go func() {
		if err := h.PublishEvent(e, a.Source); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
		logger.Info().Str("source", a.Source.String()).Msg("send event")
	}()
	return nil
}
