package main

import (
	"github.com/elojah/game_01"
)

func (a *app) Move(event game.Event) error {

	move := event.Action.(game.Move)

	// #Check permission on target
	permission, err := a.GetPermission(game.PermissionSubset{
		Source: event.Source,
		Target: move.Target,
	})
	if err != nil {
		return err
	}
	if permission.Value != game.Owner {
		return game.ErrInsufficientRights
	}

	// #Retrieve previous state target.
	target, err := a.GetEntity(game.EntitySubset{Key: move.Target.String(), Max: event.TS.UnixNano()})
	if err != nil {
		return err
	}

	// #Check if target has move in correct boundaries.
	if game.AxisDistance(target.Position, move.Position) > a.moveTolerance {
		return game.ErrInvalidAction
	}

	// #Move target
	target.MoveTo(move.Position)

	// #Write new target state.
	return a.CreateEntity(target, event.TS.UnixNano())
}
