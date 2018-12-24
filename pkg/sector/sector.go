package sector

import (
	"github.com/elojah/game_01/pkg/entity"
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
	GetSector(ulid.ID) (S, error)
}

// Service represents a sector service helper to move between sectors.
type Service interface {
	Up(float64) error
	Move(entity.E, geometry.Position) (entity.E, error)
}
