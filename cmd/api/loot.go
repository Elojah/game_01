package main

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/mux"
)

func (h *handler) loot(ctx context.Context, msg event.DTO) error {

	logger := log.With().
		Str("packet", ctx.Value(mux.Key("packet")).(string)).
		Str("action", "loot").
		Logger()

	loot := msg.Query.Loot
	e := event.E{
		ID:    msg.ID,
		Token: msg.Token,
		Action: event.Action{
			LootSource: &event.LootSource{
				ItemID:   loot.ItemID,
				TargetID: loot.TargetID,
			},
		},
	}

	logger = logger.With().Str("event", e.ID.String()).Logger()

	if err := h.event.Publish(e, loot.Source); err != nil {
		logger.Error().Err(err).Msg("event rejected")
	}
	logger.Info().Str("source", loot.Source.String()).Msg("send event")

	return nil
}
