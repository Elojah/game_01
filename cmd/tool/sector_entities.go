package main

import (
	"encoding/json"
	"net/http"

	"github.com/elojah/game_01/pkg/sector"
	"github.com/rs/zerolog/log"
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

	var entities []sector.Entities
	if err := decoder.Decode(&entities); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("sector_entities", len(entities)).Msg("found")

	for _, e := range entities {
		for _, entityID := range e.EntityIDs {
			if err := h.SectorEntitiesStore.AddEntityToSector(entityID, e.SectorID); err != nil {
				logger.Error().Err(err).
					Str("sector", e.SectorID.String()).
					Str("entity", entityID.String()).
					Msg("failed to add entity to sector")
				http.Error(w, "store failure", http.StatusInternalServerError)
				return
			}
		}
	}
	w.WriteHeader(http.StatusOK)
}
