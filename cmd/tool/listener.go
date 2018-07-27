package main

import (
	"encoding/json"
	"net/http"

	"github.com/elojah/game_01/pkg/infra"
	"github.com/rs/zerolog/log"
)

func (h *handler) listener(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.postListener(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *handler) postListener(w http.ResponseWriter, r *http.Request) {
	logger := log.With().Str("method", "POST").Str("route", "/listener").Logger()

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var listeners []infra.Listener
	if err := decoder.Decode(&listeners); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("listeners", len(listeners)).Msg("found")

	for _, l := range listeners {
		switch l.Action {
		case infra.Open:
			if _, err := h.L.New(l.ID); err != nil {
				logger.Error().Err(err).Str("listener", l.ID.String()).Msg("failed to set listener")
				return
			}
		case infra.Close:
			if err := h.L.Delete(l.ID); err != nil {
				logger.Error().Err(err).Str("listener", l.ID.String()).Msg("failed to set listener")
				return
			}
		}
	}
}
