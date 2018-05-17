package main

import (
	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

func (a *app) MoveDone(event game.Event) error {
	return nil
}

func (a *app) MoveReceived(event game.Event) error {

	move := event.Action.(game.MoveReceived)

	// #Check permission token/source.
	permission, err := a.GetPermission(game.PermissionSubset{
		Source: event.Source.String(),
		Target: move.Source.String(),
	})
	if err == storage.ErrNotFound || (err != nil && game.Right(permission.Value) != game.Owner) {
		return game.ErrInsufficientRights
	}
	if err != nil {
		return err
	}

	// #Check permission source/target if source != target.
	if move.Source != move.Target {
		permission, err := a.GetPermission(game.PermissionSubset{
			Source: move.Source.String(),
			Target: move.Target.String(),
		})
		if err == storage.ErrNotFound || (err != nil && game.Right(permission.Value) != game.Owner) {
			return game.ErrInsufficientRights
		}
		if err != nil {
			return err
		}
	}

	// #Retrieve previous state target.
	target, err := a.GetEntity(game.EntitySubset{Key: move.Target.String(), Max: event.TS.UnixNano()})
	if err != nil {
		return err
	}

	// #Check if target has move in correct boundaries.
	if game.Segment(target.Position, move.Position) > a.moveTolerance {
		return game.ErrInvalidAction
	}

	// #Move target
	target.MoveTo(move.Position)

	// #Write new target state.
	return a.SetEntity(target, event.TS.UnixNano())
}
