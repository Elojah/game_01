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
	return nil
}

func (a *app) PerformSource(id ulid.ID, e event.E) error {

	p := e.Action.GetValue().(*event.Perform)

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
		ID: p.AbilityID,
	})
	if err == gerrors.ErrNotFound {
		return errors.Wrapf(gerrors.ErrInsufficientACLs, "get ability %s for %s", p.AbilityID.String(), id.String())
	}
	if err != nil {
		return errors.Wrapf(err, "get ability %s for %s", p.AbilityID.String(), id.String())
	}

	// #Check cast was not interrupted.
	if source.Cast == nil ||
		source.Cast.AbilityID != p.AbilityID ||
		source.Cast.TS.Add(ab.CastTime) != e.TS {
		// normal behavior, don't return errors
		return nil
	}

	// #Send event to all targets
	if len(p.Targets.Positions) != 0 {
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
	wg.Add(len(p.Targets.Entities))
	for _, ens := range p.Targets.Entities {
		for _, id := range ens.IDs {
			go func(id ulid.ID) {
				if err := a.EventQStore.PublishEvent(event.E{
					ID: ulid.NewID(),
					TS: e.TS.Add(time.Nanosecond), // Add TS + 1 ns to apply damage
					Action: event.Action{
						Perform: p,
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
