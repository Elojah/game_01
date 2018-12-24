package main

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/mux"
)

func (h *handler) consume(ctx context.Context, msg event.DTO) error {

	logger := log.With().
		Str("packet", ctx.Value(mux.Key("packet")).(string)).
		Str("action", "consume").
		Logger()

	consume := msg.Query.GetConsume()
	e := event.E{
		ID:    msg.ID,
		Token: msg.Token,
		Action: event.Action{
			ConsumeSource: &event.ConsumeSource{
				ItemID: consume.ItemID,
			},
		},
	}

	logger = logger.With().Str("event", e.ID.String()).Logger()

	if err := h.PublishEvent(e, consume.Source); err != nil {
		logger.Error().Err(err).Msg("event rejected")
	}
	logger.Info().Str("source", consume.Source.String()).Msg("send event")

	return nil
}
