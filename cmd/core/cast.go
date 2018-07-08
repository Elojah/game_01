package main

import (
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/storage"
	"github.com/elojah/game_01/pkg/ulid"
)

func (a *app) Cast(id ulid.ID, e event.E) error {

	cast := e.Action.(event.Cast)

	if id.Compare(cast.Source) == 0 {
		return a.CastSource(id, e)
	}
	return a.CastTarget(id, e)
}

func (a *app) CastSource(id ulid.ID, e event.E) error {

	cast := e.Action.(event.Cast)

	// #Check permission token/source.
	permission, err := a.PermissionMapper.GetPermission(entity.PermissionSubset{
		Source: e.Source.String(),
		Target: cast.Source.String(),
	})
	if err == storage.ErrNotFound || (err != nil && account.ACL(permission.Value) != account.Owner) {
		return errors.Wrapf(account.ErrInsufficientACLs, "get permission token %s for %s", e.Source.String(), cast.Source.String())
	}
	if err != nil {
		return errors.Wrapf(err, "get permission token %s for %s", e.Source.String(), cast.Source.String())
	}

	// #Retrieve ability.
	ability, err := a.AbilityMapper.GetAbility(ability.Subset{
		ID:       cast.AbilityID,
		EntityID: cast.Source,
	})
	if err == storage.ErrNotFound {
		return errors.Wrapf(account.ErrInsufficientACLs, "get ability %s for %s", cast.AbilityID.String(), cast.Source.String())
	}
	if err != nil {
		return errors.Wrapf(err, "get ability %s for %s", cast.AbilityID.String(), cast.Source.String())
	}

	_ = ability

	return nil
}

func (a *app) CastTarget(id ulid.ID, e event.E) error {

	cast := e.Action.(event.Cast)

	// #Check permission token/source.
	permission, err := a.PermissionMapper.GetPermission(entity.PermissionSubset{
		Source: e.Source.String(),
		Target: cast.Source.String(),
	})
	if err == storage.ErrNotFound || (err != nil && account.ACL(permission.Value) != account.Owner) {
		return errors.Wrapf(account.ErrInsufficientACLs, "get permission token %s for %s", e.Source.String(), cast.Source.String())
	}
	if err != nil {
		return errors.Wrapf(err, "get permission token %s for %s", e.Source.String(), cast.Source.String())
	}

	// #Retrieve ability.
	ability, err := a.AbilityMapper.GetAbility(ability.Subset{
		ID:       cast.AbilityID,
		EntityID: cast.Source,
	})
	if err == storage.ErrNotFound {
		return errors.Wrapf(account.ErrInsufficientACLs, "get ability %s for %s", cast.AbilityID.String(), cast.Source.String())
	}
	if err != nil {
		return errors.Wrapf(err, "get ability %s for %s", cast.AbilityID.String(), cast.Source.String())
	}

	source, err := a.EntityMapper.GetEntity(entity.Subset{ID: cast.Source, MaxTS: e.TS.UnixNano()})
	if err != nil {
		return errors.Wrapf(err, "get entity %s at max ts %d", cast.Source.String(), e.TS.UnixNano())
	}

	target, err := a.EntityMapper.GetEntity(entity.Subset{ID: id, MaxTS: e.TS.UnixNano()})
	if err != nil {
		return errors.Wrapf(err, "get entity %s at max ts %d", id.String(), e.TS.UnixNano())
	}

	afb := ability.Affect(&target)
	if err := a.FeedbackMapper.SetAbilityFeedback(afb); err != nil {
		return errors.Wrapf(err, "set ability feedback %s", afb.ID.String())
	}
	fb := event.E{
		ID:     ulid.NewID(),
		TS:     e.TS,
		Source: e.Source,
		Action: event.Feedback{
			AbilityID: afb.ID,
			Source:    source.ID,
			Target:    target.ID,
		},
	}
	return errors.Wrapf(a.PublishEvent(fb, source.ID), "publish event %s to %s", fb.ID.String(), e.Source.String())
}
