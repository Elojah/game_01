package main

import (
	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

func (a *app) CreateEntity(event game.Event) error {

	create := event.Action.(game.CreateEntity)

	// #Check permission token/source.
	permission, err := a.GetPermission(game.PermissionSubset{
		Source: event.Source,
		Target: create.Source,
	})
	if err == storage.ErrNotFound || (err != nil && permission.Value != game.Owner) {
		return game.ErrInsufficientRights
	}
	if err != nil {
		return err
	}

	return nil
}
