package handler

import (
	"context"
	"encoding/json"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/dto"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/mux"
)

// H is a handler to handle entities updates or ACK.
type H struct {
	*mux.M
	ACK chan ulid.ID
}

// Dial starts handler listening.
func (h *H) Dial() error {
	h.M.Listen()
	h.ACK = make(chan ulid.ID)
	return nil
}

// Close closes the handler and ACK chan.
func (h *H) Close() error {
	if err := h.M.Close(); err != nil {
		return err
	}
	close(h.ACK)
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

	raw, err := json.Marshal(e.Domain())
	if err != nil {
		logger.Error().Err(err).Str("status", "unmarshalable").Msg("packet rejected")
		return err
	}
	if _, err = os.Stdout.Write(append(raw, '\n')); err != nil {
		logger.Error().Err(err).Str("status", "unwritable").Msg("packet rejected")
		return err
	}
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

	h.ACK <- ack.ID
	return nil
}
