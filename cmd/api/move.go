package main

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/mux"
)

func (h *handler) move(ctx context.Context, msg event.DTO) error {

	logger := log.With().
		Str("packet", ctx.Value(mux.Key("packet")).(string)).
		Str("action", "move").
		Logger()

	move := msg.Query.GetMove()
	e := event.E{
		ID:    msg.ID,
		Token: msg.Token,
		Action: event.Action{
			MoveSource: &event.MoveSource{
				Targets:  move.Targets,
				Position: move.Position,
			},
		},
	}

	logger = logger.With().Str("event", e.ID.String()).Logger()

	go func() {
		if err := h.PublishEvent(e, move.Source); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
		logger.Info().Str("source", move.Source.String()).Msg("send event")
	}()

	return nil
}
