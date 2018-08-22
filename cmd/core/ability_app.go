package main

import (
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

type abilityApp struct {
	EntityStore entity.Store

	Source  entity.E
	Ability ability.A
	Targets ability.Targets
	TS      time.Time

	Feedback ability.Feedback
	err      *multierror.Error
}

func newAbilityApp(source entity.E, ab ability.A, casted *event.Casted, ts time.Time) *abilityApp {
	return &abilityApp{
		Source:  source,
		Ability: ab,
		Targets: casted.Targets,
		TS:      ts,
		Feedback: ability.Feedback{
			ID:         ulid.NewID(),
			AbilityID:  ab.ID,
			Components: make([]ability.ComponentFeedback, len(ab.Components)),
		},
	}

}

func (a *abilityApp) Run() {

	var result *multierror.Error
	for i, c := range a.Ability.Components {
		switch c.GetValue().(type) {
		case ability.Damage:
			if err := a.Damage(a.Source, c, a.Targets, a.TS); err != nil {
				result = multierror.Append(result, errors.Wrapf(err, "damage direct component %d", i))
			}
		case ability.Heal:
			result = multierror.Append(result, gerrors.ErrNotImplementedYet)
		case ability.HealOverTime:
			result = multierror.Append(result, gerrors.ErrNotImplementedYet)
		case ability.DamageOverTime:
			result = multierror.Append(result, gerrors.ErrNotImplementedYet)
		default:
			result = multierror.Append(result, gerrors.ErrNotImplementedYet)
		}
	}
}
