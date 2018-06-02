package main

import (
	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

func (a *app) Move(id game.ID, event game.Event) error {

	move := event.Action.(game.Move)

	if id.Compare(move.Source) == 0 {
		return a.MoveSource(event)
	}
	if id.Compare(move.Target) == 0 {
		return a.MoveTarget(event)
	}
	return nil
}

func (a *app) MoveSource(event game.Event) error {
	// TODO
	// check if source is not stun/slience/unable to move units.
	// in this case cancel (what mechanism ?) the move on both source + target.
	// And if there is some bonus for moving one, add it here.
	return nil
}

func (a *app) MoveTarget(event game.Event) error {

	move := event.Action.(game.Move)

	// #Check permission token/source.
	permission, err := a.GetPermission(game.PermissionSubset{
		Source: event.Source.String(),
		Target: move.Source.String(),
	})
	if err == storage.ErrNotFound || (err != nil && game.ACL(permission.Value) != game.Owner) {
		return game.ErrInsufficientACLs
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
		if err == storage.ErrNotFound || (err != nil && game.ACL(permission.Value) != game.Owner) {
			return game.ErrInsufficientACLs
		}
		if err != nil {
			return err
		}
	}

	// #Retrieve previous state target.
	target, err := a.GetEntity(game.EntitySubset{Key: move.Target.String(), MaxTS: event.TS.UnixNano()})
	if err != nil {
		return err
	}

	// #Check if target has move in correct boundaries.
	if game.Segment(target.Position.Coord, move.Position) > a.moveTolerance {
		return game.ErrInvalidAction
	}

	// #Move target
	target.Move(move.Position)

	// #Write new target state.
	return a.SetEntity(target, event.TS.UnixNano())
}
