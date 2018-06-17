package main

import (
	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

func (a *app) Cast(id game.ID, event game.Event) error {

	cast := event.Action.(game.Cast)

	if id.Compare(cast.Source) == 0 {
		return a.CastSource(id, event)
	}
	return a.CastTarget(id, event)
}

func (a *app) CastSource(id game.ID, event game.Event) error {

	cast := event.Action.(game.Cast)

	// #Check permission token/source.
	permission, err := a.GetPermission(game.PermissionSubset{
		Source: event.Source.String(),
		Target: cast.Source.String(),
	})
	if err == storage.ErrNotFound || (err != nil && game.ACL(permission.Value) != game.Owner) {
		return game.ErrInsufficientACLs
	}
	if err != nil {
		return err
	}

	// #Retrieve ability.
	ability, err := a.GetAbility(game.AbilitySubset{
		ID:       cast.AbilityID,
		EntityID: cast.Source,
	})
	if err == storage.ErrNotFound {
		return game.ErrInsufficientACLs
	}
	if err != nil {
		return err
	}

	_ = ability

	return nil
}

func (a *app) CastTarget(id game.ID, event game.Event) error {

	cast := event.Action.(game.Cast)

	// #Check permission token/source.
	permission, err := a.GetPermission(game.PermissionSubset{
		Source: event.Source.String(),
		Target: cast.Source.String(),
	})
	if err == storage.ErrNotFound || (err != nil && game.ACL(permission.Value) != game.Owner) {
		return game.ErrInsufficientACLs
	}
	if err != nil {
		return err
	}

	// #Retrieve ability.
	ability, err := a.GetAbility(game.AbilitySubset{
		ID:       cast.AbilityID,
		EntityID: cast.Source,
	})
	if err == storage.ErrNotFound {
		return game.ErrInsufficientACLs
	}
	if err != nil {
		return err
	}

	source, err := a.GetEntity(game.EntitySubset{Key: cast.Source.String(), MaxTS: event.TS.UnixNano()})
	if err != nil {
		return err
	}

	target, err := a.GetEntity(game.EntitySubset{Key: id.String(), MaxTS: event.TS.UnixNano()})
	if err != nil {
		return err
	}

	afb := ability.Affect(&target)
	if err := a.SetAbilityFeedback(afb); err != nil {
		return err
	}
	return a.SendEvent(game.Event{
		ID:     game.NewID(),
		TS:     event.TS,
		Source: event.Source,
		Action: game.Feedback{
			AfbID:  afb.ID,
			Source: source.ID,
			Target: target.ID,
		},
	}, source.ID)
}
