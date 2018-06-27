package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/ulid"
)

type abilityWithEntity struct {
	ability.A
	EntityID ulid.ID `json:"entity_id"`
}

func (h *handler) ability(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.postAbilities(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *handler) postAbilities(w http.ResponseWriter, r *http.Request) {

	logger := log.With().Str("method", "POST").Str("route", "/ability").Logger()

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var abilities []abilityWithEntity
	if err := decoder.Decode(&abilities); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("abilities", len(abilities)).Msg("found")

	for _, a := range abilities {
		if err := h.AbilityMapper.SetAbility(a.A, a.EntityID); err != nil {
			logger.Error().Err(err).Str("ability", a.ID.String()).Msg("failed to set ability")
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
