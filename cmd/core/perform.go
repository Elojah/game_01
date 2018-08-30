package main

import (
	"sync"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

func (a *app) PerformSource(id ulid.ID, e event.E) error {

	perform := e.Action.GetValue().(*event.PerformSource)

	// #Retrieve entity
	source, err := a.EntityStore.GetEntity(entity.Subset{
		ID:    id,
		MaxTS: e.TS.UnixNano(),
	})
	if err != nil {
		return errors.Wrapf(err, "get entity %s", id.String())
	}

	// #Retrieve ability.
	ab, err := a.AbilityStore.GetAbility(ability.Subset{
		ID: perform.AbilityID,
	})
	if err == gerrors.ErrNotFound {
		return errors.Wrapf(gerrors.ErrInsufficientACLs, "get ability %s for %s", perform.AbilityID.String(), id.String())
	}
	if err != nil {
		return errors.Wrapf(err, "get ability %s for %s", perform.AbilityID.String(), id.String())
	}

	// #Check cast was not interrupted.
	if source.Cast == nil ||
		source.Cast.AbilityID != perform.AbilityID ||
		source.Cast.TS.Add(ab.CastTime) != e.TS {
		// normal behavior, don't return errors
		return nil
	}

	// #Set ability LastUsed
	ab.LastUsed = e.TS
	if err := a.AbilityStore.SetAbility(ab, source.ID); err != nil {
		return errors.Wrapf(err, "set ability %s for %s", ab.ID.String(), source.ID.String())
	}

	// #For all ability components.
	var result *multierror.Error
	errC := make(chan error, 0)
	go func() {
		for err := range errC {
			result = multierror.Append(result, err)
		}
	}()
	for cid := range ab.Components {

		// #Retrieve targets for this component.
		target, ok := perform.Targets[cid]
		if !ok {
			return errors.Wrapf(gerrors.ErrMissingTarget, "component %s for ability %s", cid, ab.ID.String())
		}

		// #Send event to all targets
		if len(target.Positions) != 0 {
			return gerrors.ErrNotImplementedYet
		}

		var wg sync.WaitGroup
		wg.Add(len(target.Entities))
		for _, id := range target.Entities {
			go func(id ulid.ID) {
				if err := a.EventQStore.PublishEvent(event.E{
					ID: ulid.NewID(),
					TS: e.TS.Add(time.Nanosecond), // Add TS + 1 ns to apply damage
					Action: event.Action{
						PerformTarget: &event.PerformTarget{
							AbilityID:   ab.ID,
							ComponentID: ulid.MustParse(cid),
							Source:      source.ID,
						},
					},
				}, id); err != nil {
					errC <- err
				}
			}(id)
		}
		wg.Wait()
		close(errC)
	}
	return result.ErrorOrNil()
}

func (a *app) PerformTarget(id ulid.ID, e event.E) error {

	perform := e.Action.GetValue().(*event.PerformTarget)

	// #Retrieve source state.
	source, err := a.EntityStore.GetEntity(entity.Subset{ID: perform.Source, MaxTS: e.TS.UnixNano()})
	if err != nil {
		return errors.Wrapf(err, "get entity %s at max ts %s", id.String(), e.TS.UnixNano())
	}

	// #Retrieve previous target state.
	target, err := a.EntityStore.GetEntity(entity.Subset{ID: id, MaxTS: e.TS.UnixNano()})
	if err != nil {
		return errors.Wrapf(err, "get entity %s at max ts %s", id.String(), e.TS.UnixNano())
	}

	// #Retrieve ability.
	ab, err := a.AbilityStore.GetAbility(ability.Subset{
		ID: perform.AbilityID,
	})
	if err == gerrors.ErrNotFound {
		return errors.Wrapf(gerrors.ErrInsufficientACLs, "get ability %s for %s", perform.AbilityID.String(), id.String())
	}
	if err != nil {
		return errors.Wrapf(err, "get ability %s for %s", perform.AbilityID.String(), id.String())
	}

	// #Initialize feedback.
	fb := ability.Feedback{
		ID:          ulid.NewID(),
		AbilityID:   ab.ID,
		ComponentID: perform.ComponentID,
	}

	// #Apply all ability components.
	cid := perform.ComponentID.String()
	component, ok := ab.Components[cid]
	if !ok {
		return errors.Wrapf(gerrors.ErrMissingTarget, "component %s for ability %s", cid, ab.ID.String())
	}
	var result *multierror.Error
	for _, effect := range component.Effects {
		veffect := effect.GetValue()
		switch veffect.(type) {
		case ability.Damage:
			fb.Effects = append(fb.Effects, ability.EffectFeedback{
				DamageFeedback: target.Damage(source, veffect.(ability.Damage)),
			})
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

	// #Set entity new state.
	if err := a.EntityStore.SetEntity(target, e.TS.UnixNano()); err != nil {
		return errors.Wrapf(err, "set entity %s", source.ID.String())
	}

	// #Set feedback.
	if err := a.FeedbackStore.SetFeedback(fb); err != nil {
		return errors.Wrapf(err, "set feedback %s from %s", fb.ID.String(), target.ID.String())
	}

	// #Publish feedback to source.
	return a.EventQStore.PublishEvent(event.E{
		ID: ulid.NewID(),
		TS: e.TS.Add(ab.CastTime),
		Action: event.Action{
			FeedbackTarget: &event.FeedbackTarget{
				ID:     fb.ID,
				Source: target.ID,
			},
		},
	}, perform.Source)
}
