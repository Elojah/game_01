package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/dto"
	"github.com/elojah/mux"
)

func (h *handler) createPC(ctx context.Context, msg dto.Message) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	id := game.NewULID()
	a := msg.Action.(dto.SetPC)
	token := game.ID(msg.Token)
	ts := time.Unix(0, msg.TS)
	entityType := game.EntityType(a.Type)

	go func() {
		err := h.SendEvent(game.Event{
			ID:     id,
			Source: token,
			TS:     ts,
			Action: game.SetPC{
				Type: entityType,
			},
		}, token)
		if err != nil {
			logger.Error().Err(err).Str("event", "unmarshalable").Msg("event rejected")
		}
	}()
	return nil
}

func (h *handler) connectPC(ctx context.Context, msg dto.Message) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	id := game.NewULID()
	a := msg.Action.(dto.ConnectPC)
	token := game.ID(msg.Token)
	ts := time.Unix(0, msg.TS)
	target := game.ID(a.Target)

	go func() {
		err := h.SendEvent(game.Event{
			ID:     id,
			Source: token,
			TS:     ts,
			Action: game.ConnectPC{
				Target: target,
			},
		}, token)
		if err != nil {
			logger.Error().Err(err).Str("event", "unmarshalable").Msg("event rejected")
		}
	}()
	return nil

}
