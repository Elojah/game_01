package svc

import (
	"github.com/elojah/game_01/pkg/ability"
	gulid "github.com/elojah/game_01/pkg/ulid"
	"github.com/pkg/errors"
)

// Service implements ability service.
type Service struct {
	AbilityStore         ability.Store
	AbilityStarterStore  ability.StarterStore
	AbilityTemplateStore ability.TemplateStore
}

// SetStarterAbilities implement Service with local stores.
func (s *Service) SetStarterAbilities(entityID gulid.ID, typeID gulid.ID) error {
	st, err := s.AbilityStarterStore.GetStarter(typeID)
	if err != nil {
		return errors.Wrap(err, "set abilities")
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
