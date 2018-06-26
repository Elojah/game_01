package entity

import (
	game "github.com/elojah/game_01"
)

// Type represents the type of an entity.
type Type = game.ID

// Position represents an entity position in world.
type Position struct {
	Coord    game.Vec3
	SectorID game.ID
}

// E represents a dynamic entity.
type E struct {
	ID       game.ID  `json:"id"`
	Type     Type     `json:"type"`
	Name     string   `json:"name"`
	HP       uint64   `json:"hp"`
	MP       uint64   `json:"mp"`
	Position Position `json:"position"`
}

// Move moves entity to position p.
func (e *E) Move(p game.Vec3) {
	e.Position.Coord.Add(p)
}

// Mapper is an interface for E object.
type Mapper interface {
	SetEntity(E, int64) error
	GetEntity(Subset) (E, error)
}

// Subset is a subset to retrieve one entity.
type Subset struct {
	Key    string
	MaxTS  int64
	Cursor uint64
	Count  int64
}

// Equal returns if both entities are equal.
func (e E) Equal(entity E) bool {
	if e.ID.Compare(entity.ID) != 0 {
		return false
	}
	if e.Type.Compare(entity.Type) != 0 {
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
	if e.Position.SectorID.Compare(entity.Position.SectorID) != 0 {
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