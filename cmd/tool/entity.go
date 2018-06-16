package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
)

func (h *handler) entity(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.postEntities(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *handler) postEntities(w http.ResponseWriter, r *http.Request) {
	logger := log.With().Str("method", "POST").Str("route", "/entity").Logger()

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var entities []game.Entity
	if err := decoder.Decode(&entities); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("entities", len(entities)).Msg("found")

	for _, entity := range entities {
		if err := h.SetEntity(entity, time.Now().UnixNano()); err != nil {
			logger.Error().Err(err).Str("entity", entity.ID.String()).Msg("failed to set entity")
			return
		}
	}
}
