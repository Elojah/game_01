package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
)

func (h *handler) sectorEntities(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.postSectorEntities(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *handler) postSectorEntities(w http.ResponseWriter, r *http.Request) {
	logger := log.With().Str("method", "POST").Str("route", "/sector/entities").Logger()

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var sectorEntities []game.SectorEntities
	if err := decoder.Decode(&sectorEntities); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("sector_entities", len(sectorEntities)).Msg("found")

	for _, se := range sectorEntities {
		for _, entityID := range se.EntityIDs {
			if err := h.AddEntityToSector(entityID, se.SectorID); err != nil {
				logger.Error().Err(err).
					Str("sector", se.SectorID.String()).
					Str("entity", entityID.String()).
					Msg("failed to add entity to sector")
				return
			}
		}
	}
}
