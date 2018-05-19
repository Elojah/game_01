package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/dto"
	"github.com/elojah/mux"
)

func (h *handler) createEntity(ctx context.Context, msg dto.Message) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	id := game.NewULID()
	a := msg.Action.(dto.SetEntity)
	token := game.ID(msg.Token)
	ts := time.Unix(0, msg.TS)
	source := game.ID(a.Source)
	entityType := game.EntityType(a.Type)
	position := game.Vec3(a.Position)

	go func() {
		err := h.SendEvent(game.Event{
			ID:     id,
			Source: token,
			TS:     ts,
			Action: game.SetEntity{
				Source:   source,
				Type:     entityType,
				Position: position,
			},
		}, source)
		if err != nil {
			logger.Error().Err(err).Str("event", "unmarshalable").Msg("event rejected")
		}
	}()
	return nil
}
