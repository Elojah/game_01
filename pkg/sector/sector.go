package sector

import (
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/ulid"
)

// Adjacents return ids of all adjacent sectors.
func (s S) Adjacents() []ulid.ID {
	ids := make([]ulid.ID, len(s.Connections))
	for i, co := range s.Connections {
		ids[i] = co.External.SectorID
	}
	return ids
}

// ClosestConnection the closest connection from coord.
func (s S) ClosestConnection(coord geometry.Vec3) Connection {
	min := s.Dim.X + s.Dim.Y + s.Dim.Z
	iMin := 0
	for i, co := range s.Connections {
		if geometry.Segment(coord, co.Coord) < min {
			iMin = i
		}
	}
	return s.Connections[iMin]
}

// Out returns if a coord is still in the sector.
func (s S) Out(coord geometry.Vec3) bool {
	return coord.X < 0 ||
		coord.X > s.Dim.X ||
		coord.Y < 0 ||
		coord.Y > s.Dim.Y ||
		coord.Z < 0 ||
		coord.Z > s.Dim.Z
}

// Store is the service for S.
type Store interface {
	SetSector(S) error
	GetSector(Subset) (S, error)
}

// Subset allows to retrieve on sector by ID.
type Subset struct {
	ID ulid.ID
}
