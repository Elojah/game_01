package main

import (
	"encoding/json"
	"net/http"

	"github.com/elojah/game_01/pkg/sector"
	"github.com/rs/zerolog/log"
)

func (h *handler) sectorStarters(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.postSectors(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *handler) postSectorsStarters(w http.ResponseWriter, r *http.Request) {
	logger := log.With().Str("method", "POST").Str("route", "/sector/starters").Logger()

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var starters []sector.Starter
	if err := decoder.Decode(&starters); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("starters", len(starters)).Msg("found")

	for _, s := range starters {
		if err := h.SetStarter(s); err != nil {
			logger.Error().Err(err).Str("starter", s.SectorID.String()).Msg("failed to set starter")
			return
		}
	}
}