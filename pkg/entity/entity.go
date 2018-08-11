package entity

import (
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/ulid"
)

// Move moves entity to position p.
func (e *E) Move(p geometry.Vec3) {
	e.Position.Coord = p
}

// Mapper is an interface for E object.
type Mapper interface {
	SetEntity(E, int64) error
	GetEntity(Subset) (E, error)
	DelEntity(Subset) error
}

// Subset is a subset to retrieve one entity.
type Subset struct {
	ID     ulid.ID
	MaxTS  int64
	Cursor uint64
	Count  int64
}
