package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/ulid"
)

// AbilityWithEntity represents the payload to create/associate new ability.
type AbilityWithEntity struct {
	ability.A
	ID       string `json:"id"`
	Type     string `json:"type"`
	EntityID ulid.ID
	Entity   string `json:"entity"`
}

// UnmarshalJSON customs unmarshal to define an ability.
func (a *AbilityWithEntity) UnmarshalJSON(data []byte) error {
	type alias AbilityWithEntity
	var al alias
	if err := json.Unmarshal(data, &al); err != nil {
		return err
	}
	var err error
	if al.A.ID, err = ulid.Parse(al.ID); err != nil {
		return err
	}
	if al.A.Type, err = ulid.Parse(al.Type); err != nil {
		return err
	}
	if al.EntityID, err = ulid.Parse(al.Entity); err != nil {
		return err
	}
	return nil
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
	decoder.DisallowUnknownFields()
	defer r.Body.Close()

	var abilities []AbilityWithEntity
	if err := decoder.Decode(&abilities); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("abilities", len(abilities)).Msg("found")

	for _, a := range abilities {
		if err := h.AbilityMapper.SetAbility(a.A, a.EntityID); err != nil {
			logger.Error().Err(err).Str("ability", ulid.String(a.A.ID)).Msg("failed to set ability")
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
