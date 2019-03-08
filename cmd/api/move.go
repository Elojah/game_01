package main

import (
	"context"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/mux"
)

func (h *handler) move(ctx context.Context, msg event.DTO) error {

	logger := log.With().
		Str("packet", ctx.Value(mux.Key("packet")).(string)).
		Str("action", "move").
		Logger()

	move := msg.Query.Move
	e := event.E{
		ID:    msg.ID,
		Token: msg.Token,
		Action: event.Action{
			MoveTarget: &event.MoveTarget{
				Position: move.Position,
			},
		},
	}

	logger = logger.With().Str("event", e.ID.String()).Logger()

	var g errgroup.Group
	for _, target := range move.Targets {
		target := target
		g.Go(func() error {
			return h.PublishEvent(e, target)
		})
	}
	if err := g.Wait(); err != nil {
		logger.Error().Err(err).Msg("event rejected")
		return err
	}
	logger.Info().Str("id", e.ID.String()).Msg("send event")

	return nil
}
