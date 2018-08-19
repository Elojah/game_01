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

func (a *app) Casted(id ulid.ID, e event.E) error {

	cast := e.Action.GetValue().(*event.Casted)

	if id.Compare(cast.Source) == 0 {
		return a.CastedSource(id, e)
	}
	return a.CastedTarget(id, e)
}

func (a *app) CastedSource(id ulid.ID, e event.E) error {

	cast := e.Action.GetValue().(*event.Casted)

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

	e = event.E{
		Action: event.Action{
			Casted: (*event.Casted)(cast),
		},
	}
	_ = e
	_ = ab

	return nil
}

func (a *app) CastedTarget(id ulid.ID, e event.E) error {

	cast := e.Action.GetValue().(*event.Casted)

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

	source, err := a.EntityStore.GetEntity(entity.Subset{ID: cast.Source, MaxTS: e.TS.UnixNano()})
	if err != nil {
		return errors.Wrapf(err, "get entity %s at max ts %d", cast.Source.String(), e.TS.UnixNano())
	}

	target, err := a.EntityStore.GetEntity(entity.Subset{ID: id, MaxTS: e.TS.UnixNano()})
	if err != nil {
		return errors.Wrapf(err, "get entity %s at max ts %d", id.String(), e.TS.UnixNano())
	}

	afb := ab.Affect(&target)
	if err := a.SetFeedback(afb); err != nil {
		return errors.Wrapf(err, "set ability feedback %s", afb.ID.String())
	}
	fb := event.E{
		ID:     ulid.NewID(),
		TS:     e.TS,
		Source: e.Source,
		Action: event.Action{
			Feedback: &event.Feedback{
				ID:     afb.ID,
				Source: source.ID,
				Target: target.ID,
			},
		},
	}
	return errors.Wrapf(a.PublishEvent(fb, source.ID), "publish event %s to %s", fb.ID.String(), e.Source.String())
}