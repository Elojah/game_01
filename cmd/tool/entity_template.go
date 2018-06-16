package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
)

func (h *handler) entityTemplate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.postEntityTemplates(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *handler) postEntityTemplates(w http.ResponseWriter, r *http.Request) {
	logger := log.With().Str("method", "POST").Str("route", "/entity/template").Logger()

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var entityTemplates []game.EntityTemplate
	if err := decoder.Decode(&entityTemplates); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("entity_templates", len(entityTemplates)).Msg("found")

	for _, entityTemplate := range entityTemplates {
		if err := h.SetEntityTemplate(entityTemplate); err != nil {
			logger.Error().Err(err).Str("entity_template", entityTemplate.ID.String()).Msg("failed to set entity_template")
			return
		}
	}
}
