package sector

import (
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/ulid"
)

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
