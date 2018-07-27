package sector

import (
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/ulid"
)

// Connection represents a connection points between two sectors.
type Connection struct {
	Coord    geometry.Vec3
	External geometry.Position
}

// S represents a cuboid in the world.
type S struct {
	ID          ulid.ID
	Dim         geometry.Vec3
	Connections []Connection
}

// Adjacents return ids of all adjacent sectors.
func (s S) Adjacents() []ulid.ID {
	ids := make([]ulid.ID, len(s.Connections))
	for i, co := range s.Connections {
		ids[i] = co.External.SectorID
	}
	return ids
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

// Mapper is the service for S.
type Mapper interface {
	SetSector(S) error
	GetSector(Subset) (S, error)
}

// Subset allows to retrieve on sector by ID.
type Subset struct {
	ID ulid.ID
}
