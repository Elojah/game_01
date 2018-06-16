package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
)

func (h *handler) sector(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.postSectors(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *handler) postSectors(w http.ResponseWriter, r *http.Request) {
	logger := log.With().Str("method", "POST").Str("route", "/sector").Logger()

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var sectors []game.Sector
	if err := decoder.Decode(&sectors); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("sectors", len(sectors)).Msg("found")

	for _, sector := range sectors {
		if err := h.SetSector(sector); err != nil {
			logger.Error().Err(err).Str("sector", sector.ID.String()).Msg("failed to set sector")
			return
		}
	}
}
