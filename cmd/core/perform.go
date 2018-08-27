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

func (a *app) Perform(id ulid.ID, e event.E) error {
	if e.Action.GetValue().(*event.Perform).Source.Compare(id) == 0 {
		return a.PerformSource(id, e)
	}
	return a.PerformTarget(id, e)
}

func (a *app) PerformTarget(id ulid.ID, e event.E) error {

	perform := e.Action.GetValue().(*event.Perform)

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
		ID:        ulid.NewID(),
		AbilityID: ab.ID,
	}
	// #Apply all ability components.
	var result *multierror.Error
	for i, comp := range ab.Components {
		c := comp.GetValue()
		switch c.(type) {
		case ability.Damage:
			cfb, err := target.Damage(source, c.(ability.Damage))
			if err != nil {
				result = multierror.Append(result, errors.Wrapf(err, "damage direct component %d", i))
				continue
			}
			fb.Components = append(fb.Components, cfb)
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

	return nil
}

func (a *app) PerformSource(id ulid.ID, e event.E) error {

	perform := e.Action.GetValue().(*event.Perform)

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

	// #Send event to all targets
	if len(perform.Targets.Positions) != 0 {
		return gerrors.ErrNotImplementedYet
	}
	var result *multierror.Error
	errC := make(chan error, 0)
	go func() {
		for err := range errC {
			result = multierror.Append(result, err)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(len(perform.Targets.Entities))
	for _, ens := range perform.Targets.Entities {
		for _, id := range ens.IDs {
			go func(id ulid.ID) {
				if err := a.EventQStore.PublishEvent(event.E{
					ID: ulid.NewID(),
					TS: e.TS.Add(time.Nanosecond), // Add TS + 1 ns to apply damage
					Action: event.Action{
						Perform: perform,
					},
				}, id); err != nil {
					errC <- err
				}
			}(id)
		}
	}
	wg.Wait()
	close(errC)

	return result.ErrorOrNil()
}
