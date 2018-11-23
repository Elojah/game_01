package main

import (
	"encoding/json"
	"net/http"

	"github.com/elojah/game_01/pkg/infra"
	"github.com/rs/zerolog/log"
)

func (h *handler) sequencer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.postSequencer(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *handler) postSequencer(w http.ResponseWriter, r *http.Request) {
	logger := log.With().Str("method", "POST").Str("route", "/sequencer").Logger()

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var sequencers []infra.Sequencer
	if err := decoder.Decode(&sequencers); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("sequencers", len(sequencers)).Msg("found")

	for _, s := range sequencers {
		switch s.Action {
		case infra.Open:
			if _, err := h.SequencerService.New(s.ID); err != nil {
				logger.Error().Err(err).Str("sequencer", s.ID.String()).Msg("failed to set sequencer")
				return
			}
		case infra.Close:
			if err := h.SequencerService.Remove(s.ID); err != nil {
				logger.Error().Err(err).Str("sequencer", s.ID.String()).Msg("failed to set sequencer")
				return
			}
		}
	}
	w.WriteHeader(http.StatusOK)
}
