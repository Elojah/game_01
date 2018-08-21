package main

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/entity"
	serrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

func (a *app) Casted(id ulid.ID, e event.E) error {

	casted := e.Action.GetValue().(*event.Casted)

	// #Retrieve entity
	en, err := a.EntityStore.GetEntity(entity.Subset{
		ID:    id,
		MaxTS: e.TS.UnixNano(),
	})
	if err != nil {
		return errors.Wrapf(err, "get entity %s", id.String())
	}

	// #Retrieve ability.
	ab, err := a.AbilityStore.GetAbility(ability.Subset{
		ID:       casted.AbilityID,
		EntityID: id,
	})
	if err == serrors.ErrNotFound {
		return errors.Wrapf(serrors.ErrInsufficientACLs, "get ability %s for %s", casted.AbilityID.String(), id.String())
	}
	if err != nil {
		return errors.Wrapf(err, "get ability %s for %s", casted.AbilityID.String(), id.String())
	}

	// #Check cast was not interrupted.
	if en.Cast == nil ||
		en.Cast.AbilityID != casted.AbilityID ||
		en.Cast.TS.Add(ab.CastTime) != e.TS {
		log.Info().
			Str("entity", id.String()).
			Int64("ts", e.TS.UnixNano()).
			Str("ability", ab.ID.String()).
			Msg("interrupted cast")
			// normal behavior, don't return errors
		return nil
	}

	for _, c := range ab.Components {
		switch c.GetValue().(type) {
		case ability.DamageDirect:
			return serrors.ErrNotImplementedYet
		case ability.HealDirect:
			return serrors.ErrNotImplementedYet
		case ability.HealOverTime:
			return serrors.ErrNotImplementedYet
		case ability.DamageOverTime:
			return serrors.ErrNotImplementedYet
		}
	}

	return nil
}
