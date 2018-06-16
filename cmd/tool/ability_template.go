package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
)

func (h *handler) abilityTemplate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.postAbilityTemplates(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *handler) postAbilityTemplates(w http.ResponseWriter, r *http.Request) {
	logger := log.With().Str("method", "POST").Str("route", "/ability/template").Logger()

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var abilityTemplates []game.AbilityTemplate
	if err := decoder.Decode(&abilityTemplates); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("ability_templates", len(abilityTemplates)).Msg("found")

	for _, abilityTemplate := range abilityTemplates {
		if err := h.SetAbilityTemplate(abilityTemplate); err != nil {
			logger.Error().Err(err).Str("ability_template", abilityTemplate.ID.String()).Msg("failed to set ability_template")
			return
		}
	}
}
