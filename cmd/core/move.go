package main

import (
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

func (svc *service) MoveTarget(id ulid.ID, e event.E) error {

	mt := e.Action.MoveTarget
	ts := e.ID.Time()

	// #Check permission source/token
	if err := svc.EntityPermissionService.CheckPermission(e.Token, id); err != nil {
		return errors.Wrap(err, "move target")
	}

	// #Retrieve previous state target.
	target, err := svc.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrap(err, "move target")
	}

	// #Check if entity is alive
	if target.Dead {
		return errors.Wrap(gerrors.ErrIsDead{EntityID: id.String()}, "move target")
	}

	if target, err = svc.SectorService.Move(target, mt.Position); err != nil {
		return errors.Wrap(err, "move target")
	}

	return errors.Wrap(svc.EntityStore.SetEntity(target, ts+1), "move target")
}
