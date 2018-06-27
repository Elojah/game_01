package main

import (
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/perm"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/game_01/storage"
)

func (a *app) Move(id ulid.ID, e event.E) error {

	move := e.Action.(event.Move)

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

	move := e.Action.(event.Move)

	// #Check permission token/source.
	permission, err := a.GetPermission(perm.Subset{
		Source: e.Source.String(),
		Target: move.Source.String(),
	})
	if err == storage.ErrNotFound || (err != nil && account.ACL(permission.Value) != account.Owner) {
		return game.ErrInsufficientACLs
	}
	if err != nil {
		return err
	}

	// #Check permission source/target if source != target.
	if move.Source != move.Target {
		permission, err := a.GetPermission(perm.Subset{
			Source: move.Source.String(),
			Target: move.Target.String(),
		})
		if err == storage.ErrNotFound || (err != nil && account.ACL(permission.Value) != account.Owner) {
			return game.ErrInsufficientACLs
		}
		if err != nil {
			return err
		}
	}

	// #Retrieve previous state target.
	target, err := a.EntityMapper.GetEntity(entity.Subset{Key: move.Target.String(), MaxTS: e.TS.UnixNano()})
	if err != nil {
		return err
	}

	// #Check if target has move in correct boundaries.
	if game.Segment(target.Position.Coord, move.Position) > a.moveTolerance {
		return game.ErrInvalidAction
	}

	// #Retrieve current sector
	s, err := a.SectorMapper.GetSector(sector.Subset{ID: target.Position.SectorID})
	if err != nil {
		return err
	}

	// #Move target
	target.Move(move.Position)
	// If target is out of sector, move to next
	if s.Out(target.Position.Coord) {
		bp := s.ClosestBP(target.Position.Coord)
		nextSector, err := a.SectorMapper.GetSector(sector.Subset{ID: bp.SectorID})
		if err != nil {
			return err
		}
		oppBP := nextSector.FindBP(bp.ID)
		target.Position.SectorID = nextSector.ID
		target.Position.Coord.MoveReference(bp.Position, oppBP.Position)
		if err := a.EntitiesMapper.AddEntityToSector(target.ID, nextSector.ID); err != nil {
			return err
		}
		if err := a.EntitiesMapper.RemoveEntityToSector(target.ID, nextSector.ID); err != nil {
			return err
		}
	}

	// #Write new target state.
	return a.EntityMapper.SetEntity(target, e.TS.UnixNano())
}
