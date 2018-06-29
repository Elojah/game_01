package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/elojah/game_01/storage"
	"github.com/elojah/mux"
	"github.com/rs/zerolog/log"
)

type handler struct {
	*mux.M
}

func (h *handler) handle(ctx context.Context, raw []byte) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	// #Unmarshal entity.
	var entity storage.Entity
	if _, err := entity.Unmarshal(raw); err != nil {
		logger.Error().Err(err).Str("status", "unformatted").Msg("packet rejected")
		return err
	}

	raw, err := json.Marshal(entity)
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
