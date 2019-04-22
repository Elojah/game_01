package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/entity"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

func (h *handler) spawn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.postSpawns(w, r)
	case "GET":
		h.getSpawn(w, r)
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

func (h *handler) getSpawn(w http.ResponseWriter, r *http.Request) {

	logger := log.With().Str("method", "GET").Str("route", "/spawn").Logger()

	paramIDs, ok := r.URL.Query()["ids"]
	if !ok || len(paramIDs) == 0 {
		http.Error(w, "missing ids", http.StatusBadRequest)
		return
	}
	strIDs := strings.Split(paramIDs[0], ",")
	spawnIDs := make([]gulid.ID, len(strIDs))
	for i, strID := range strIDs {
		var err error
		spawnIDs[i], err = gulid.Parse(strID)
		if err != nil {
			http.Error(w, fmt.Sprintf("id invalid: %s", strID), http.StatusBadRequest)
			return
		}
	}

	logger.Info().Int("spawns", len(spawnIDs)).Msg("found")

	spawns := make([]entity.Spawn, len(spawnIDs))
	for i, id := range spawnIDs {
		var err error
		spawns[i], err = h.EntitySpawnStore.GetSpawn(id)
		if err != nil {
			logger.Error().Err(err).Str("spawn", id.String()).Msg("failed to get spawn")
			http.Error(w, "store failure", http.StatusInternalServerError)
			return
		}
	}
	raw, err := json.Marshal(spawns)
	if err != nil {
		logger.Error().Err(err).Msg("failed to marshal spawns")
		http.Error(w, "marshal failure", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(raw); err != nil {
		logger.Error().Err(err).Msg("failed to write response")
		http.Error(w, "write failure", http.StatusInternalServerError)
		return
	}
}
