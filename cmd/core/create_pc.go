package main

import (
	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

const (
	characterKey = "pc"
)

func (a *app) CreatePC(event game.Event) error {

	create := event.Action.(game.CreatePC)

	// #Check token permission to create a new PC.
	permission, err := a.GetPermission(game.PermissionSubset{
		Source: event.Source.String(),
		Target: characterKey,
	})
	if err != nil {
		return err
	}
	if permission.Value <= 0 {
		return game.ErrInvalidAction
	}

	// #Decrease token permission to create a new PC by 1.
	if err := a.SetPermission(game.Permission{
		Source: event.Source.String(),
		Target: characterKey,
	}, permission.Value-1); err != nil {
		return err
	}

	return nil
}
