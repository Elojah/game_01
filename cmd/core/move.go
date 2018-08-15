package main

import (
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	serrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/pkg/errors"
)

func (a *app) Move(id ulid.ID, e event.E) error {

	move := e.Action.GetValue().(*event.Move)

	if id.Compare(move.Source) == 0 {
		return a.MoveSource(e)
	}
	if id.Compare(move.Target) == 0 {
		return a.MoveTarget(e)
	}
	return nil
}

func (a *app) MoveSource(e event.E) error {
	// TODO
	// check if source is not stun/slience/unable to move units.
	// in this case cancel (what mechanism ?) the move on both source + target.
	// And if there is some bonus for moving one, add it here.
	return nil
}

func (a *app) MoveTarget(e event.E) error {

	move := e.Action.GetValue().(*event.Move)

	// #Check permission token/source.
	permission, err := a.GetPermission(entity.PermissionSubset{
		Source: e.Source.String(),
		Target: move.Source.String(),
	})
	if err == serrors.ErrNotFound || (err != nil && account.ACL(permission.Value) != account.Owner) {
		return errors.Wrapf(account.ErrInsufficientACLs, "get permission token %s for %s", e.Source.String(), move.Source.String())
	}
	if err != nil {
		return errors.Wrapf(err, "get permission token %s for %s", e.Source.String(), move.Source.String())
	}

	// #Check permission source/target if source != target.
	if move.Source != move.Target {
		permission, err := a.GetPermission(entity.PermissionSubset{
			Source: move.Source.String(),
			Target: move.Target.String(),
		})
		if err == serrors.ErrNotFound || (err != nil && account.ACL(permission.Value) != account.Owner) {
			return errors.Wrapf(account.ErrInsufficientACLs, "get permission entity %s for %s", move.Source.String(), move.Target.String())
		}
		if err != nil {
			return errors.Wrapf(err, "get permission entity %s for %s", move.Source.String(), move.Target.String())
		}
	}

	// #Retrieve previous state target.
	target, err := a.EntityStore.GetEntity(entity.Subset{ID: move.Target, MaxTS: e.TS.UnixNano()})
	if err != nil {
		return errors.Wrapf(err, "get entity %s at max ts %s", move.Target.String(), e.TS.UnixNano())
	}

	// #Retrieve current sector
	s, err := a.SectorStore.GetSector(sector.Subset{ID: target.Position.SectorID})
	if err != nil {
		return errors.Wrapf(err, "get sector %s", target.Position.SectorID.String())
	}

	if target.Position.SectorID.Compare(move.Position.SectorID) == 0 {

		// #Check if target has moved in correct boundaries in same sector.
		if s.Out(target.Position.Coord) {
			return errors.Wrapf(
				account.ErrInvalidAction,
				"check in sector %s (%f , %f , %f) from (%f , %f , %f) to (%f , %f , %f) for entity %s",
				s.ID.String(),
				s.Dim.X,
				s.Dim.Y,
				s.Dim.Z,
				target.Position.Coord.X,
				target.Position.Coord.Y,
				target.Position.Coord.Z,
				move.Position.Coord.X,
				move.Position.Coord.Y,
				move.Position.Coord.Z,
				target.ID.String(),
			)
		}

		// #Check if target has moved at a tolerable distance in same sector.
		if geometry.Segment(target.Position.Coord, move.Position.Coord) > a.moveTolerance {
			return errors.Wrapf(
				account.ErrInvalidAction,
				"check move tolerance %f from (%f , %f , %f) to (%f , %f , %f) for entity %s",
				a.moveTolerance,
				target.Position.Coord.X,
				target.Position.Coord.Y,
				target.Position.Coord.Z,
				move.Position.Coord.X,
				move.Position.Coord.Y,
				move.Position.Coord.Z,
				target.ID.String(),
			)
		}

		// #Move target
		target.Position.Coord = move.Position.Coord

	} else {

		// #Check if new sector is a neighbour.
		neigh, ok := s.Neighbours[move.Position.SectorID.String()]
		if !ok {
			return errors.Wrapf(
				account.ErrInvalidAction,
				"invalid next neighbour sector %s with previous %s",
				move.Position.SectorID.String(),
				target.Position.SectorID.String(),
			)
		}

		// #Check if target has moved at a tolerable distance in different sectors.
		if geometry.Segment(target.Position.Coord, move.Position.Coord.MoveReference(neigh)) > a.moveTolerance {
			return errors.Wrapf(
				account.ErrInvalidAction,
				"check move tolerance %f from %s (%f , %f , %f) to %s (%f , %f , %f) for entity %s",
				a.moveTolerance,
				target.Position.SectorID.String(),
				target.Position.Coord.X,
				target.Position.Coord.Y,
				target.Position.Coord.Z,
				move.Position.SectorID.String(),
				move.Position.Coord.X,
				move.Position.Coord.Y,
				move.Position.Coord.Z,
				target.ID.String(),
			)
		}

		// #Add entity to new sector and remove from previous.
		if err := a.AddEntityToSector(target.ID, move.Position.SectorID); err != nil {
			return errors.Wrapf(err, "add entity %s to sector %s", target.ID.String(), move.Position.SectorID.String())
		}
		if err := a.RemoveEntityFromSector(target.ID, target.Position.SectorID); err != nil {
			return errors.Wrapf(err, "remove entity %s from sector %s", target.ID.String(), s.ID.String())
		}

		// #Move target
		target.Position = move.Position
	}

	// #Write new target state.
	return errors.Wrapf(a.EntityStore.SetEntity(target, e.TS.UnixNano()), "set entity %s for ts %d", target.ID.String(), e.TS.UnixNano())
}
