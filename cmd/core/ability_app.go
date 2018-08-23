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

// AbilityApp wraps ability applications mechanism.
type AbilityApp struct {
	EntityStore entity.Store

	Source  entity.E
	Targets ability.Targets
	TS      time.Time
}

// NewAbilityApp returns a new ability app ready to run.
func NewAbilityApp(source entity.E, targets ability.Targets, ts time.Time) *AbilityApp {
	return &AbilityApp{
		Source:  source,
		Targets: targets,
		TS:      ts,
	}

}

// Run applies the ability from source to targets at ts defined in NewAbilityApp.
func (a *AbilityApp) Run(ab ability.A) (ability.Feedback, error) {

	fb := ability.Feedback{
		ID:        ulid.NewID(),
		AbilityID: ab.ID,
	}

	var result *multierror.Error
	for i, comp := range ab.Components {
		c := comp.GetValue()
		switch c.(type) {
		case ability.Damage:
			cfbs, err := a.ApplyDamage(c.(ability.Damage))
			if err != nil {
				result = multierror.Append(result, errors.Wrapf(err, "damage direct component %d", i))
				continue
			}
			fb.Components = append(fb.Components, cfbs...)
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

	return fb, result.ErrorOrNil()
}
