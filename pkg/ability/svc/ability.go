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
		return errors.Wrap(err, "set starter abilities")
	}

	for _, abilityID := range st.AbilityIDs {
		ab, err := s.AbilityTemplateStore.GetTemplate(abilityID)
		if err != nil {
			return errors.Wrap(err, "set starter abilities")
		}
		if err := s.AbilityStore.SetAbility(ab, entityID); err != nil {
			return errors.Wrap(err, "set starter abilities")
		}
	}
	return nil
}

// CopyAbilities implement Service with local stores.
func (s *Service) CopyAbilities(sourceID gulid.ID, targetID gulid.ID) error {
	abilities, err := s.AbilityStore.ListAbility(sourceID)
	if err != nil {
		return errors.Wrap(err, "copy abilities")
	}

	for _, ab := range abilities {
		if err := s.AbilityStore.SetAbility(ab, targetID); err != nil {
			return errors.Wrap(err, "copy abilities")
		}
	}
	return nil
}
