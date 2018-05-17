package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/dto"
	"github.com/elojah/mux"
)

func (h *handler) createEntity(ctx context.Context, a dto.CreateEntity, token game.Token, ts time.Time) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	go func() {
		err := h.SendEvent(game.Event{
			ID:     game.NewULID(),
			Source: token.ID,
			TS:     ts,
			Action: game.CreateEntity{
				Source:   game.ID(a.Source),
				Type:     game.EntityType(a.Type),
				Position: game.Vec3(a.Position),
			},
		}, game.ID(a.Source))
		if err != nil {
			logger.Error().Err(err).Str("event", "unmarshalable").Msg("event rejected")
		}
	}()
	return nil
}
