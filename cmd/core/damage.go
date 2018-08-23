package main

import (
	"sync"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

// ApplyDamage applies damage to all targets in ability app.
func (a *AbilityApp) ApplyDamage(dd ability.Damage) ([]ability.DamageFeedback, error) {

	// #Retrieve all targets
	var dfbs []ability.DamageFeedback

	if len(a.Targets.Positions) != 0 {
		return nil, gerrors.ErrNotImplementedYet
	}

	dfbC := make(chan ability.DamageFeedback, 0)
	go func() {
		for dfb := range dfbC {
			dfbs = append(dfbs, dfb)
		}
	}()

	var result *multierror.Error
	errC := make(chan error, 0)
	go func() {
		for err := range errC {
			result = multierror.Append(result, err)
		}
	}()

	// #For all targeted entities
	var wg sync.WaitGroup
	wg.Add(len(a.Targets.Entities))
	for _, e := range a.Targets.Entities {
		for _, id := range e.IDs {
			go func(id ulid.ID) {
				// #Apply damage to one entity async
				dfb, err := a.ApplyDamageSingle(dd, id)
				dfbC <- dfb
				errC <- err
				wg.Done()
			}(id)
		}
	}
	wg.Wait()
	close(dfbC)
	close(errC)

	return dfbs, result.ErrorOrNil()
}

// ApplyDamageSingle applies damage to target id.
func (a *AbilityApp) ApplyDamageSingle(dd ability.Damage, id ulid.ID) (ability.DamageFeedback, error) {

	// #Retrieve target entity
	target, err := a.EntityStore.GetEntity(entity.Subset{
		ID:    id,
		MaxTS: a.TS.UnixNano(),
	})
	if err != nil {
		return ability.DamageFeedback{}, errors.Wrapf(err, "get entity %s", id.String())
	}

	// #Applies direct damage to target entity
	return target.Damage(a.Source, dd), nil
}
