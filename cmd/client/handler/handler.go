package handler

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/dto"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/mux"
)

// H is a handler to handle entities updates or ACK.
type H struct {
	*mux.M
	ACK    chan infra.ACK
	Entity chan entity.E
}

// Dial starts handler listening.
func (h *H) Dial() error {
	h.M.Listen()
	h.ACK = make(chan infra.ACK, 100)
	h.Entity = make(chan entity.E, 1000)
	return nil
}

// Close closes the handler and ACK chan.
func (h *H) Close() error {
	if err := h.M.Close(); err != nil {
		return err
	}
	close(h.ACK)
	close(h.Entity)
	return nil
}

// HandleEntity is the handler for entities updates from sync server.
func (h *H) HandleEntity(ctx context.Context, raw []byte) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	// #Unmarshal entity.
	var e dto.Entity
	if _, err := e.Unmarshal(raw); err != nil {
		logger.Error().Err(err).Str("status", "unformatted").Msg("packet rejected")
		return err
	}

	h.Entity <- e.Domain()
	return nil
}

// HandleACK is the handler for ack packets from api server.
func (h *H) HandleACK(ctx context.Context, raw []byte) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	// #Unmarshal entity.
	var ack dto.ACK
	if _, err := ack.Unmarshal(raw); err != nil {
		logger.Error().Err(err).Str("status", "unformatted").Msg("packet rejected")
		return err
	}

	h.ACK <- ack.Domain()
	return nil
}
