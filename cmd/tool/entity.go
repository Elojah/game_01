package main

import (
	"encoding/json"
	"net/http"

	"github.com/oklog/ulid"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/entity"
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

	var entities []entity.E
	if err := decoder.Decode(&entities); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("entities", len(entities)).Msg("found")

	for _, e := range entities {
		if err := h.EntityStore.SetEntity(e, ulid.Now()); err != nil {
			logger.Error().Err(err).Str("entity", e.ID.String()).Msg("failed to set entity")
			http.Error(w, "store failure", http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
