package entity

import (
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/ulid"
)

// Type represents the type of an entity.
type Type = ulid.ID

// E represents a dynamic entity.
type E struct {
	ID       ulid.ID           `json:"id"`
	Type     Type              `json:"type"`
	Name     string            `json:"name"`
	HP       uint64            `json:"hp"`
	MP       uint64            `json:"mp"`
	Position geometry.Position `json:"position"`
}

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

// Equal returns if both entities are equal.
func (e E) Equal(en E) bool {
	if ulid.Compare(e.ID, en.ID) != 0 {
		return false
	}
	if ulid.Compare(e.Type, en.Type) != 0 {
		return false
	}
	if e.Name != en.Name {
		return false
	}
	if e.HP != en.HP {
		return false
	}
	if e.MP != en.MP {
		return false
	}
	if ulid.Compare(e.Position.SectorID, en.Position.SectorID) != 0 {
		return false
	}
	if e.Position.Coord.X != en.Position.Coord.X {
		return false
	}
	if e.Position.Coord.Y != en.Position.Coord.Y {
		return false
	}
	if e.Position.Coord.Z != en.Position.Coord.Z {
		return false
	}
	return true
}
