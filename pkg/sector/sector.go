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

// Store contains basic operations for sector s object.
type Store interface {
	Upsert(S) error
	Fetch(ulid.ID) (S, error)
}

// App contains sector stores and applications.
type App interface {
	Store
	EntitiesStore
	Dial(tolerance float64)
	Move(entity.E, geometry.Position) (entity.E, error)
	Segment(geometry.Position, geometry.Position) (float64, error)
}
