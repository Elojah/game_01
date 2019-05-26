package main

import (
	"encoding/json"
	"net/http"

	gulid "github.com/elojah/game_01/pkg/ulid"
	"github.com/rs/zerolog/log"
)

func (h *handler) lootHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.postLoots(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *handler) postLoots(w http.ResponseWriter, r *http.Request) {
	logger := log.With().Str("method", "POST").Str("route", "/loot").Logger()

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var loots []gulid.ID
	if err := decoder.Decode(&loots); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("loots", len(loots)).Msg("found")

	for _, l := range loots {
		if err := h.item.UpsertLoot(l); err != nil {
			logger.Error().Err(err).Str("loot", l.String()).Msg("failed to set loot")
			http.Error(w, "store failure", http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
