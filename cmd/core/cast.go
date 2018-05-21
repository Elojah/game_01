package main

import (
	"github.com/elojah/game_01"
	// "github.com/elojah/game_01/storage"
)

func (a *app) Cast(id game.ID, event game.Event) error {

	cast := event.Action.(game.Cast)

	if id.Compare(cast.Source) == 0 {
		return a.CastSource(event)
	}
	if id.Compare(cast.Target) == 0 {
		return a.CastTarget(event)
	}
	return nil
}

func (a *app) CastSource(id game.ID, event game.Event) error {

	cast := event.Action.(game.Cast)

	// #Check permission token/source.
	permission, err := a.GetPermission(game.PermissionSubset{
		Source: event.Source.String(),
		Target: cast.Source.String(),
	})
	if err == storage.ErrNotFound || (err != nil && game.Right(permission.Value) != game.Owner) {
		return game.ErrInsufficientRights
	}
	if err != nil {
		return err
	}

	// #Retrieve skill.
	skill, err := a.GetSkill(game.SkillSubset{
		ID:       cast.SkillID,
		EntityID: cast.Source,
	})
	if err == storage.ErrNotFound {
		return game.ErrInsufficientRights
	}
	if err != nil {
		return err
	}

	return nil
}
