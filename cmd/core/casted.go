package main

import (
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

func (a *app) Casted(id ulid.ID, e event.E) error {

	casted := e.Action.GetValue().(*event.Casted)

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
		ID: casted.AbilityID,
	})
	if err == gerrors.ErrNotFound {
		return errors.Wrapf(gerrors.ErrInsufficientACLs, "get ability %s for %s", casted.AbilityID.String(), id.String())
	}
	if err != nil {
		return errors.Wrapf(err, "get ability %s for %s", casted.AbilityID.String(), id.String())
	}

	// #Check cast was not interrupted.
	if source.Cast == nil ||
		source.Cast.AbilityID != casted.AbilityID ||
		source.Cast.TS.Add(ab.CastTime) != e.TS {
		// normal behavior, don't return errors
		return nil
	}

	abApp := newAbilityApp(source, ab, casted, e.TS)
	abApp.EntityStore = a.EntityStore

	abApp.Run()

	return result.ErrorOrNil()
}
