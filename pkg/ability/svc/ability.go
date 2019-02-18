package svc

import (
	"net/http"

	"github.com/elojah/game_01/pkg/ability"
	gulid "github.com/elojah/game_01/pkg/ulid"
	"github.com/pkg/errors"
)

type Service struct {
	AbilityStore         ability.Store
	AbilityTemplateStore ability.TemplateStore
}

func (s *Service) SetStarterAbilities(entityID gulid.ID, typeID gulid.ID) error {
	st, err := h.AbilityStarterStore.GetStarter(template.ID)
	if err != nil {
		logger.Error().Err(err).Str("template", template.ID.String()).Msg("failed to get starter abilities")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, abilityID := range st.AbilityIDs {
		ab, err := s.AbilityTemplateStore.GetTemplate(abilityID)
		if err != nil {
			return errors.Wrap(err, "set abilities")
		}
		if err := s.AbilityStore.SetAbility(ab, entityID); err != nil {
			return errors.Wrap(err, "set abilities")
		}
	}
	return nil
}
