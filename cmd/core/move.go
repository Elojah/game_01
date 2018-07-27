package main

import (
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/storage"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/pkg/errors"
)

func (a *app) Move(id ulid.ID, e event.E) error {

	move := e.Action.(event.Move)

	if ulid.Compare(id, move.Source) == 0 {
		return a.MoveSource(e)
	}
	if ulid.Compare(id, move.Target) == 0 {
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

	move := e.Action.(event.Move)

	// #Check permission token/source.
	permission, err := a.PermissionMapper.GetPermission(entity.PermissionSubset{
		Source: ulid.String(e.Source),
		Target: ulid.String(move.Source),
	})
	if err == storage.ErrNotFound || (err != nil && account.ACL(permission.Value) != account.Owner) {
		return errors.Wrapf(account.ErrInsufficientACLs, "get permission token %s for %s", ulid.String(e.Source), ulid.String(move.Source))
	}
	if err != nil {
		return errors.Wrapf(err, "get permission token %s for %s", ulid.String(e.Source), ulid.String(move.Source))
	}

	// #Check permission source/target if source != target.
	if move.Source != move.Target {
		permission, err := a.PermissionMapper.GetPermission(entity.PermissionSubset{
			Source: ulid.String(move.Source),
			Target: ulid.String(move.Target),
		})
		if err == storage.ErrNotFound || (err != nil && account.ACL(permission.Value) != account.Owner) {
			return errors.Wrapf(account.ErrInsufficientACLs, "get permission entity %s for %s", ulid.String(move.Source), ulid.String(move.Target))
		}
		if err != nil {
			return errors.Wrapf(err, "get permission entity %s for %s", ulid.String(move.Source), ulid.String(move.Target))
		}
	}

	// #Retrieve previous state target.
	target, err := a.EntityMapper.GetEntity(entity.Subset{ID: move.Target, MaxTS: e.TS.UnixNano()})
	if err != nil {
		return errors.Wrapf(err, "get entity %s at max ts %s", ulid.String(move.Target), e.TS.UnixNano())
	}

	// #Retrieve current sector
	s, err := a.SectorMapper.GetSector(sector.Subset{ID: target.Position.SectorID})
	if err != nil {
		return errors.Wrapf(err, "get sector %s", ulid.String(target.Position.SectorID))
	}

	if ulid.Compare(target.Position.SectorID, move.Position.SectorID) == 0 {

		// #Check if target has moved in correct boundaries in same sector.
		if s.Out(target.Position.Coord) {
			return errors.Wrapf(
				account.ErrInvalidAction,
				"check in sector %s (%f , %f , %f) from (%f , %f , %f) to (%f , %f , %f) for entity %s",
				ulid.String(s.ID),
				s.Dim.X,
				s.Dim.Y,
				s.Dim.Z,
				target.Position.Coord.X,
				target.Position.Coord.Y,
				target.Position.Coord.Z,
				move.Position.Coord.X,
				move.Position.Coord.Y,
				move.Position.Coord.Z,
				ulid.String(target.ID),
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
				ulid.String(target.ID),
			)
		}

		// #Move target
		target.Move(move.Position.Coord)

	} else {

		// #Find closest connection point for current out coord.
		con := s.Closest(target.Position.Coord)

		// #Move coordinates to new reference.
		target.Position.Coord.MoveReference(con.Coord, con.External.Coord)

		// #Add entity to new sector and remove from previous.
		if err := a.EntitiesMapper.AddEntityToSector(target.ID, con.External.SectorID); err != nil {
			return errors.Wrapf(err, "add entity %s to sector %s", ulid.String(target.ID), ulid.String(con.External.SectorID))
		}
		if err := a.EntitiesMapper.RemoveEntityToSector(target.ID, target.Position.SectorID); err != nil {
			return errors.Wrapf(err, "remove entity %s from sector %s", ulid.String(target.ID), ulid.String(s.ID))
		}

		// #Change entity SectorID.
		target.Position.SectorID = con.External.SectorID
	}

	// #Write new target state.
	return errors.Wrapf(a.EntityMapper.SetEntity(target, e.TS.UnixNano()), "set entity %s for ts %d", ulid.String(target.ID), e.TS.UnixNano())
}
