package main

import (
	"encoding/json"
	"net/http"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/rs/zerolog/log"
)

func (h *handler) abilityStarter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.postAbilityStarters(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *handler) postAbilityStarters(w http.ResponseWriter, r *http.Request) {
	logger := log.With().Str("method", "POST").Str("route", "/ability/starter").Logger()

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var starters []ability.Starter
	if err := decoder.Decode(&starters); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("ability_starters", len(starters)).Msg("found")

	for _, st := range starters {
		if err := h.AbilityStarterStore.InsertStarter(st); err != nil {
			logger.Error().Err(err).Str("ability_starter", st.EntityID.String()).Msg("failed to set ability_starter")
			http.Error(w, "store failure", http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
