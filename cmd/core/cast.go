package main

import (
	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/game_01/storage"
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
		return account.ErrInsufficientACLs
	}
	if err != nil {
		return err
	}

	// #Retrieve ability.
	ability, err := a.AbilityMapper.GetAbility(ability.Subset{
		ID:       cast.AbilityID,
		EntityID: cast.Source,
	})
	if err == storage.ErrNotFound {
		return account.ErrInsufficientACLs
	}
	if err != nil {
		return err
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
		return account.ErrInsufficientACLs
	}
	if err != nil {
		return err
	}

	// #Retrieve ability.
	ability, err := a.AbilityMapper.GetAbility(ability.Subset{
		ID:       cast.AbilityID,
		EntityID: cast.Source,
	})
	if err == storage.ErrNotFound {
		return account.ErrInsufficientACLs
	}
	if err != nil {
		return err
	}

	source, err := a.EntityMapper.GetEntity(entity.Subset{ID: cast.Source, MaxTS: e.TS.UnixNano()})
	if err != nil {
		return err
	}

	target, err := a.EntityMapper.GetEntity(entity.Subset{ID: id, MaxTS: e.TS.UnixNano()})
	if err != nil {
		return err
	}

	afb := ability.Affect(&target)
	if err := a.FeedbackMapper.SetAbilityFeedback(afb); err != nil {
		return err
	}
	return a.SendEvent(event.E{
		ID:     ulid.NewID(),
		TS:     e.TS,
		Source: e.Source,
		Action: event.Feedback{
			AbilityID: afb.ID,
			Source:    source.ID,
			Target:    target.ID,
		},
	}, source.ID)
}
