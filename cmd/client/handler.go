package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/dto"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/mux"
)

type handler struct {
	*mux.M
	ACK chan ulid.ID
}

func (h *handler) Dial() error {
	h.M.Listen()
	h.ACK = make(chan ulid.ID)
	return nil
}

func (h *handler) Close() error {
	if err := h.M.Close(); err != nil {
		return err
	}
	close(h.ACK)
	return nil
}

func (h *handler) handleEntity(ctx context.Context, raw []byte) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	// #Unmarshal entity.
	var e entity.E
	if _, err := e.Unmarshal(raw); err != nil {
		logger.Error().Err(err).Str("status", "unformatted").Msg("packet rejected")
		return err
	}

	raw, err := json.Marshal(e)
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

func (h *handler) handleACK(ctx context.Context, raw []byte) error {

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
