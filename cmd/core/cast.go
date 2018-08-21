package main

import (
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	serrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

func (a *app) Cast(id ulid.ID, e event.E) error {

	cast := e.Action.GetValue().(*event.Cast)

	// #Check permission token/source.
	permission, err := a.GetPermission(entity.PermissionSubset{
		Source: e.Source.String(),
		Target: cast.Source.String(),
	})
	if err == serrors.ErrNotFound || (err != nil && account.ACL(permission.Value) != account.Owner) {
		return errors.Wrapf(account.ErrInsufficientACLs, "get permission token %s for %s", e.Source.String(), cast.Source.String())
	}
	if err != nil {
		return errors.Wrapf(err, "get permission token %s for %s", e.Source.String(), cast.Source.String())
	}

	// #Retrieve ability.
	ab, err := a.AbilityStore.GetAbility(ability.Subset{
		ID:       cast.AbilityID,
		EntityID: cast.Source,
	})
	if err == serrors.ErrNotFound {
		return errors.Wrapf(account.ErrInsufficientACLs, "get ability %s for %s", cast.AbilityID.String(), cast.Source.String())
	}
	if err != nil {
		return errors.Wrapf(err, "get ability %s for %s", cast.AbilityID.String(), cast.Source.String())
	}

	// #Check MP consumption
	source, err := a.EntityStore.GetEntity(entity.Subset{ID: cast.Source, MaxTS: e.TS.UnixNano()})
	if err != nil {
		return errors.Wrapf(err, "get entity %s at max ts %d", cast.Source.String(), e.TS.UnixNano())
	}
	if source.MP < ab.MPConsumption {
		return errors.Wrapf(
			account.ErrInvalidAction,
			"entity %s with MP %d for ability %s with MP: %d",
			cast.Source.String(),
			source.MP,
			ab.ID.String(),
			ab.MPConsumption,
		)
	}

	// #Check CD validity. if LastUsed + CD < now.
	if ab.LastUsed.Add(ab.CD).Before(e.TS) {
		return errors.Wrapf(account.ErrInvalidAction, "cd up for skill %s ", ab.ID.String())
	}

	// #Set entity new state with decreased MP and casting up.
	source.MP -= ab.MPConsumption
	source.Cast = &entity.Cast{AbilityID: ab.ID, TS: e.TS}
	if err := a.EntityStore.SetEntity(source, e.TS.UnixNano()); err != nil {
		return errors.Wrapf(err, "set entity %s", source.ID.String())
	}

	// #Add casted event to event set.
	if err := a.EventStore.SetEvent(event.E{
		ID: ulid.NewID(),
		TS: e.TS.Add(ab.CastTime),
		Action: event.Action{
			Casted: &event.Casted{AbilityID: ab.ID},
		},
	}, source.ID); err != nil {
		return errors.Wrapf(err, "set casted event %s", e.ID.String())
	}

	return nil
}
