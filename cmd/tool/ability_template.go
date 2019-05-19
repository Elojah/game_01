package main

import (
	"encoding/json"
	"net/http"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/rs/zerolog/log"
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

	var templates []ability.Template
	if err := decoder.Decode(&templates); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("ability_templates", len(templates)).Msg("found")

	for _, t := range templates {
		if err := h.AbilityTemplateStore.InsertTemplate(t); err != nil {
			logger.Error().Err(err).Str("ability_template", t.ID.String()).Msg("failed to set ability_template")
			http.Error(w, "store failure", http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
