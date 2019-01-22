package svc

import (
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/sector"
)

// Service is a wrapping service around sector usecases.
type Service struct {
	SectorEntitiesStore sector.EntitiesStore
	SectorStore         sector.Store

	Tolerance float64
}

// Up assigns allowed tolerance for entity move.
func (s *Service) Up(tolerance float64) error {
	s.Tolerance = tolerance
	return nil
}

// Move moves a target to a new position, considering tolerance service tolerance and sector neighbours.
func (s *Service) Move(target entity.E, newPosition geometry.Position) (entity.E, error) {

	// #Retrieve current sector
	sec, err := s.SectorStore.GetSector(target.Position.SectorID)
	if err != nil {
		return target, errors.Wrapf(err, "get sector %s", target.Position.SectorID.String())
	}

	// #If moved in same sector
	if target.Position.SectorID.Compare(newPosition.SectorID) == 0 {

		// #Check if target has moved in correct boundaries in same sector.
		if sec.Out(target.Position.Coord) {
			return target, errors.Wrap(
				errors.Wrap(
					gerrors.ErrInvalidMove{
						TargetID:       target.ID.String(),
						SectorID:       sec.ID.String(),
						SectorDim:      sec.Dim,
						TargetPosition: target.Position.Coord,
						NewSectorID:    sec.ID.String(),
						NewSectorDim:   sec.Dim,
						NewPosition:    newPosition.Coord,
					},
					"check boundaries",
				),
				"move",
			)
		}

		// #Check if target has moved at a tolerable distance in same sector.
		if geometry.Segment(target.Position.Coord, newPosition.Coord) > s.Tolerance {
			return target, errors.Wrap(
				errors.Wrap(
					gerrors.ErrInvalidMove{
						TargetID:       target.ID.String(),
						SectorID:       sec.ID.String(),
						SectorDim:      sec.Dim,
						TargetPosition: target.Position.Coord,
						NewSectorID:    sec.ID.String(),
						NewSectorDim:   sec.Dim,
						NewPosition:    newPosition.Coord,
					},
					"check tolerance",
				),
				"move",
			)
		}

		// #Move target
		target.Position.Coord = newPosition.Coord

		// #Else
	} else {

		// #Check if new sector is a neighbour.
		neigh, ok := sec.Neighbours[newPosition.SectorID.String()]
		if !ok {
			return target, errors.Wrap(
				errors.Wrap(
					gerrors.ErrInvalidMove{
						TargetID:       target.ID.String(),
						SectorID:       sec.ID.String(),
						SectorDim:      sec.Dim,
						TargetPosition: target.Position.Coord,
						NewSectorID:    newPosition.SectorID.String(),
						NewPosition:    newPosition.Coord,
					},
					"check neighbour",
				),
				"move",
			)
		}

		// #Retrieve neighbour sector (for boundaries checking)
		neighSec, err := s.SectorStore.GetSector(newPosition.SectorID)
		if err != nil {
			return target, errors.Wrap(err, "move")
		}

		// #Check if target has moved in correct boundaries in neighbour sector.
		if neighSec.Out(newPosition.Coord) {
			return target, errors.Wrap(
				errors.Wrap(
					gerrors.ErrInvalidMove{
						TargetID:       target.ID.String(),
						SectorID:       sec.ID.String(),
						SectorDim:      sec.Dim,
						TargetPosition: target.Position.Coord,
						NewSectorID:    neighSec.ID.String(),
						NewSectorDim:   neighSec.Dim,
						NewPosition:    newPosition.Coord,
					},
					"check boundaries",
				),
				"move",
			)
		}

		// #Check if target has moved at a tolerable distance in different sectors.
		if geometry.Segment(target.Position.Coord, newPosition.Coord.MoveReference(neigh)) > s.Tolerance {
			return target, errors.Wrap(
				errors.Wrap(
					gerrors.ErrInvalidMove{
						TargetID:       target.ID.String(),
						SectorID:       sec.ID.String(),
						SectorDim:      sec.Dim,
						TargetPosition: target.Position.Coord,
						NewSectorID:    neighSec.ID.String(),
						NewSectorDim:   neighSec.Dim,
						NewPosition:    newPosition.Coord,
					},
					"check tolerance",
				),
				"move",
			)
		}

		// #Add entity to new sector and remove from previous.
		if err := s.SectorEntitiesStore.AddEntityToSector(target.ID, newPosition.SectorID); err != nil {
			return target, errors.Wrap(err, "move")
		}
		if err := s.SectorEntitiesStore.RemoveEntityFromSector(target.ID, target.Position.SectorID); err != nil {
			return target, errors.Wrap(err, "move")
		}

		// #Move target
		target.Position = newPosition
	}

	// #Return new target state.
	return target, nil
}

// Segment returns the segment distance between two positions, even if different sectors.
// WORKS FOR NEIGHBOUR SEGMENTS ONLY
func (s *Service) Segment(lhs geometry.Position, rhs geometry.Position) (float64, error) {

	// #If both positions have same sector, return plain segment
	if lhs.SectorID.Compare(rhs.SectorID) == 0 {
		return geometry.Segment(lhs.Coord, rhs.Coord), nil
	}

	// #Retrieve current sector
	sec, err := s.SectorStore.GetSector(lhs.SectorID)
	if err != nil {
		return 0, errors.Wrap(err, "calculate segment")
	}

	// #Check if new sector is a neighbour.
	neigh, ok := sec.Neighbours[rhs.SectorID.String()]
	if !ok {
		return 0, errors.Wrap(
			gerrors.ErrInvalidNeighbourSector{
				SectorID:        lhs.SectorID.String(),
				SectorNeighbour: rhs.SectorID.String(),
			},
			"calculate segment",
		)
	}

	// #Check if target has moved at a tolerable distance in different sectors.
	return geometry.Segment(lhs.Coord, rhs.Coord.MoveReference(neigh)), nil
}
