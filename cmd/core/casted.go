package main

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	serrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

func (a *app) Casted(id ulid.ID, e event.E) error {

	cast := e.Action.GetValue().(*event.Casted)

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
		ID:       cast.AbilityID,
		EntityID: id,
	})
	if err == serrors.ErrNotFound {
		return errors.Wrapf(account.ErrInsufficientACLs, "get ability %s for %s", cast.AbilityID.String(), id.String())
	}
	if err != nil {
		return errors.Wrapf(err, "get ability %s for %s", cast.AbilityID.String(), id.String())
	}

	// #Check cast was not interrupted.
	if en.Cast == nil ||
		en.Cast.AbilityID != cast.AbilityID ||
		en.Cast.TS.Add(ab.CastTime) != e.TS {
		log.Info().
			Str("entity", id.String()).
			Int64("ts", e.TS.UnixNano()).
			Str("ability", ab.ID.String()).
			Msg("interrupted cast")
			// normal behavior, don't return errors
		return nil
	}

	// switch ab

	return nil
}
