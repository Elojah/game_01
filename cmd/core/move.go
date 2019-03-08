package main

import (
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

func (a *app) MoveTarget(id ulid.ID, e event.E) error {

	mt := e.Action.MoveTarget
	ts := e.ID.Time()

	// #Check permission source/token
	if err := a.EntityPermissionService.CheckPermission(e.Token, id); err != nil {
		return errors.Wrap(err, "move target")
	}

	// #Retrieve previous state target.
	target, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrap(err, "move target")
	}

	if target, err = a.SectorService.Move(target, mt.Position); err != nil {
		return errors.Wrap(err, "move target")
	}

	return errors.Wrap(a.EntityStore.SetEntity(target, ts), "move target")
}
