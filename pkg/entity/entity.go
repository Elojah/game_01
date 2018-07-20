package entity

import (
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/ulid"
)

// Type represents the type of an entity.
type Type = ulid.ID

// Position represents an entity position in world.
type Position struct {
	Coord    geometry.Vec3
	SectorID ulid.ID
}

// E represents a dynamic entity.
type E struct {
	ID       ulid.ID  `json:"id"`
	Type     Type     `json:"type"`
	Name     string   `json:"name"`
	HP       uint64   `json:"hp"`
	MP       uint64   `json:"mp"`
	Position Position `json:"position"`
}

// Move moves entity to position p.
func (e *E) Move(p geometry.Vec3) {
	e.Position.Coord.Add(p)
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
func (e E) Equal(entity E) bool {
	if ulid.Compare(e.ID, entity.ID) != 0 {
		return false
	}
	if ulid.Compare(e.Type, entity.Type) != 0 {
		return false
	}
	if e.Name != entity.Name {
		return false
	}
	if e.HP != entity.HP {
		return false
	}
	if e.MP != entity.MP {
		return false
	}
	if ulid.Compare(e.Position.SectorID, entity.Position.SectorID) != 0 {
		return false
	}
	if e.Position.Coord.X != entity.Position.Coord.X {
		return false
	}
	if e.Position.Coord.Y != entity.Position.Coord.Y {
		return false
	}
	if e.Position.Coord.Z != entity.Position.Coord.Z {
		return false
	}
	return true
}
