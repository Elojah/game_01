package main

import (
	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/entity"
	serrors "github.com/elojah/game_01/pkg/errors"
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
	if err == serrors.ErrNotFound {
		return errors.Wrapf(serrors.ErrInsufficientACLs, "get ability %s for %s", casted.AbilityID.String(), id.String())
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

	var result *multierror.Error
	for i, c := range ab.Components {
		switch c.GetValue().(type) {
		case ability.DamageDirect:
			if err := a.DamageDirect(source, c, casted.Targets, e.TS); err != nil {
				result = multierror.Append(result, errors.Wrapf(err, "damage direct component %d", i))
			}
		case ability.HealDirect:
			result = multierror.Append(result, serrors.ErrNotImplementedYet)
		case ability.HealOverTime:
			result = multierror.Append(result, serrors.ErrNotImplementedYet)
		case ability.DamageOverTime:
			result = multierror.Append(result, serrors.ErrNotImplementedYet)
		}
	}

	return result.ErrorOrNil()
}
