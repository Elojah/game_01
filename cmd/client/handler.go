package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/dto"
	"github.com/elojah/mux"
)

type handler struct {
	*mux.M
}

func (h *handler) Dial() error {
	h.M.Listen()
	return nil
}

func (h *handler) handleEntity(ctx context.Context, raw []byte) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	// #Unmarshal entity.
	var e dto.Entity
	if _, err := e.Unmarshal(raw); err != nil {
		logger.Error().Err(err).Str("status", "unformatted").Msg("packet rejected")
		return err
	}

	raw, err := json.Marshal(e)
	if err != nil {
		logger.Error().Err(err).Str("status", "unmarshalable").Msg("packet rejected")
		return err
	}
	n, err := os.Stdout.Write(raw)
	if err != nil {
		logger.Error().Err(err).Str("status", "unwritable").Msg("packet rejected")
		return err
	}
	logger.Info().Int("chars", n).Msg("packet rejected")
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

	raw, err := json.Marshal(ack)
	if err != nil {
		logger.Error().Err(err).Str("status", "unmarshalable").Msg("packet rejected")
		return err
	}
	n, err := os.Stdout.Write(raw)
	if err != nil {
		logger.Error().Err(err).Str("status", "unwritable").Msg("packet rejected")
		return err
	}
	logger.Info().Int("chars", n).Msg("packet rejected")
	return nil
}
