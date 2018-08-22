package main

import (
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

func (a *abilityApp) Damage(source entity.E, c ability.Component, targets ability.Targets, ts time.Time) (ability.Feedback, error) {

	// #Retrieve all targets
	if len(targets.Positions) != 0 {
		return ability.Feedback{}, gerrors.ErrNotImplementedYet
	}

	feedback := ability.Feedback{
		ID:         ulid.NewID(),
		Components: make([]ability.ComponentFeedback),
	}
	fbC := make(chan ability.DamageFeedback, 0)
	go func() {
		for fb := range <-fbC {
			feedback.Components = append(feedback.Components, fb)
		}
	}()

	var result *multierror.Error
	errC := make(chan ability.DamageFeedback, 0)
	go func() {
		for err := range <-errC {
			result = multierror.Append(result, err)
		}
	}()

	// #For all targeted entities
	for _, e := range targets.Entities {
		for _, id := range e.IDs {
			go func(id ulid.ID) {
				fb, err := a.DamageOne(source, c, id, ts)
				fbC <- fb
				errC <- err
			}(id)
		}
	}

	return nil
}

func (a *abilityApp) DamageOne(source entity.E, c ability.Component, id ulid.ID, ts time.Time) (ability.DamageFeedback, error) {

	// #Retrieve target entity
	target, err := a.EntityStore.GetEntity(entity.Subset{
		ID:    id,
		MaxTS: ts.UnixNano(),
	})
	if err != nil {
		return ability.DamageFeedback{}, errors.Wrapf(err, "get entity %s", id.String())
	}

	// #Applies direct damage to target entity
	return target.Damage(source, c), nil
}
