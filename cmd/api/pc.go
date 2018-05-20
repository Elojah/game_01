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

	a := msg.Action.(dto.SetPC)
	token := game.ID(msg.Token)
	event := game.Event{
		ID:     game.NewULID(),
		Source: token,
		TS:     time.Unix(0, msg.TS),
		Action: game.SetPC{
			Type: game.EntityType(a.Type),
		},
	}

	go func() {
		if err := h.SendEvent(event, token); err != nil {
			logger.Error().Err(err).Str("event", event.ID.String()).Msg("event rejected")
		}
	}()
	return nil
}

func (h *handler) connectPC(ctx context.Context, msg dto.Message) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	a := msg.Action.(dto.ConnectPC)
	token := game.ID(msg.Token)
	event := game.Event{
		ID:     game.NewULID(),
		Source: token,
		TS:     time.Unix(0, msg.TS),
		Action: game.ConnectPC{
			Target: game.ID(a.Target),
		},
	}

	go func() {
		if err := h.SendEvent(event, token); err != nil {
			logger.Error().Err(err).Str("event", event.ID.String()).Msg("event rejected")
		}
	}()
	return nil

}
