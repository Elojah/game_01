package main

import (
	"encoding/json"
	"net/http"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/rs/zerolog/log"
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

	var templates []entity.Template
	if err := decoder.Decode(&templates); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("entity_templates", len(templates)).Msg("found")

	for _, t := range templates {
		if err := h.EntityTemplateStore.SetTemplate(t); err != nil {
			logger.Error().Err(err).Str("entity_template", t.ID.String()).Msg("failed to set entity_template")
			http.Error(w, "store failure", http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
