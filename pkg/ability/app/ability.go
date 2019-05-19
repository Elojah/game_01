package app

import (
	"github.com/elojah/game_01/pkg/ability"
	gulid "github.com/elojah/game_01/pkg/ulid"
	"github.com/pkg/errors"
)

// A implements ability applications.
type A struct {
	ability.FeedbackStore
	ability.StarterStore
	ability.Store
	ability.TemplateStore
}

// SetStarters implement App with local stores.
func (app A) SetStarters(entityID gulid.ID, typeID gulid.ID) error {
	st, err := app.StarterStore.FetchStarter(typeID)
	if err != nil {
		return errors.Wrap(err, "set starter abilities")
	}

	for _, abilityID := range st.AbilityIDs {
		ab, err := app.TemplateStore.FetchTemplate(abilityID)
		if err != nil {
			return errors.Wrap(err, "set starter abilities")
		}
		if err := app.Store.Insert(ab, entityID); err != nil {
			return errors.Wrap(err, "set starter abilities")
		}
	}
	return nil
}

// Copy implement App with local stores.
func (app A) Copy(sourceID gulid.ID, targetID gulid.ID) error {
	abilities, err := app.Store.List(sourceID)
	if err != nil {
		return errors.Wrap(err, "copy abilities")
	}

	for _, ab := range abilities {
		if err := app.Store.Insert(ab, targetID); err != nil {
			return errors.Wrap(err, "copy abilities")
		}
	}
	return nil
}
