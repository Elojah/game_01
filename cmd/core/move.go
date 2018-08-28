package main

import (
	"sync"
	"time"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

func (a *app) MoveSource(id ulid.ID, e event.E) error {

	move := e.Action.GetValue().(*event.Move)

	// #Check permission token/source.
	permission, err := a.GetPermission(entity.PermissionSubset{
		Source: e.Source.String(),
		Target: move.Source.String(),
	})
	if err == gerrors.ErrNotFound || (err != nil && account.ACL(permission.Value) != account.Owner) {
		return errors.Wrapf(gerrors.ErrInsufficientACLs, "get permission token %s for %s", e.Source.String(), move.Source.String())
	}
	if err != nil {
		return errors.Wrapf(err, "get permission token %s for %s", e.Source.String(), move.Source.String())
	}

	// #TODO Check if source is not stun or forbidden to move other entities

	// #Send event to all targets
	var result *multierror.Error
	errC := make(chan error, 0)
	go func() {
		for err := range errC {
			result = multierror.Append(result, err)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(len(move.Targets))
	for _, target := range move.Targets {
		go func(target ulid.ID) {
			if err := a.EventQStore.PublishEvent(event.E{
				ID: ulid.NewID(),
				TS: e.TS.Add(time.Nanosecond), // Add TS + 1 ns to apply damage
				Action: event.Action{
					Move: move,
				},
			}, id); err != nil {
				errC <- err
			}
		}(target)
	}
	wg.Wait()
	close(errC)

	return result.ErrorOrNil()
}

func (a *app) MoveTarget(id ulid.ID, e event.E) error {

	move := e.Action.GetValue().(*event.Move)

	// #Retrieve previous state target.
	target, err := a.EntityStore.GetEntity(entity.Subset{ID: id, MaxTS: e.TS.UnixNano()})
	if err != nil {
		return errors.Wrapf(err, "get entity %s at max ts %s", id.String(), e.TS.UnixNano())
	}

	// #Retrieve current sector
	s, err := a.SectorStore.GetSector(sector.Subset{ID: target.Position.SectorID})
	if err != nil {
		return errors.Wrapf(err, "get sector %s", target.Position.SectorID.String())
	}

	// #If moved in same sector
	if target.Position.SectorID.Compare(move.Position.SectorID) == 0 {

		// #Check if target has moved in correct boundaries in same sector.
		if s.Out(target.Position.Coord) {
			return errors.Wrapf(
				gerrors.ErrInvalidAction,
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
				gerrors.ErrInvalidAction,
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

		// #Else
	} else {

		// #Check if new sector is a neighbour.
		neigh, ok := s.Neighbours[move.Position.SectorID.String()]
		if !ok {
			return errors.Wrapf(
				gerrors.ErrInvalidAction,
				"invalid next neighbour sector %s with previous %s",
				move.Position.SectorID.String(),
				target.Position.SectorID.String(),
			)
		}

		// #Check if target has moved at a tolerable distance in different sectors.
		if geometry.Segment(target.Position.Coord, move.Position.Coord.MoveReference(neigh)) > a.moveTolerance {
			return errors.Wrapf(
				gerrors.ErrInvalidAction,
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
