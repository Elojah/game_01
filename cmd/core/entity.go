package main

import (
	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

func (a *app) CreateEntity(event game.Event) error {

	se := event.Action.(game.SetEntity)
	_ = se

	// #Check permission token/source.
	permission, err := a.GetPermission(game.PermissionSubset{
		Source: event.Source.String(),
		Target: se.Source.String(),
	})
	if err == storage.ErrNotFound || (err != nil && game.Right(permission.Value) != game.Owner) {
		return game.ErrInsufficientRights
	}
	if err != nil {
		return err
	}

	return nil
}
