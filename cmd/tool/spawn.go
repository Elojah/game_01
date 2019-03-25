package main

import (
	"encoding/json"
	"net/http"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/rs/zerolog/log"
)

func (h *handler) spawn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.postSpawns(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *handler) postSpawns(w http.ResponseWriter, r *http.Request) {
	logger := log.With().Str("method", "POST").Str("route", "/spawn").Logger()

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var spawns []entity.Spawn
	if err := decoder.Decode(&spawns); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("spawns", len(spawns)).Msg("found")

	for _, sp := range spawns {
		if err := h.EntitySpawnStore.SetSpawn(sp); err != nil {
			logger.Error().Err(err).Str("spawn", sp.ID.String()).Msg("failed to set spawn")
			http.Error(w, "store failure", http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
