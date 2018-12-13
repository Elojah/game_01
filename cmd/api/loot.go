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

	loot := msg.Query.GetLoot()
	e := event.E{
		ID:    msg.ID,
		Token: msg.Token,
		Action: event.Action{
			LootSource: &event.LootSource{
				ID:     loot.Source,
				ItemID: loot.ItemID,
			},
		},
	}

	logger = logger.With().Str("event", e.ID.String()).Logger()

	go func() {
		if err := h.PublishEvent(e, loot.Source); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
		logger.Info().Str("source", loot.Source.String()).Msg("send event")
	}()
	return nil
}
