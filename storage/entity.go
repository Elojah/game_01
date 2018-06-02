package storage

import (
	"github.com/elojah/game_01"
)

// NewEntity convert a game.Entity into a storage Entity.
func NewEntity(a game.Entity) *Entity {
	return &Entity{
		ID:       [16]byte(a.ID),
		Type:     [16]byte(a.Type),
		Name:     a.Name,
		HP:       a.HP,
		MP:       a.MP,
		SectorID: a.Position.SectorID,
		X:        a.Position.Coord.X,
		Y:        a.Position.Coord.Y,
		Z:        a.Position.Coord.Z,
	}
}

// Domain converts a storage Entity into a game Entity.
func (a Entity) Domain() game.Entity {
	return game.Entity{
		ID:   game.ID(a.ID),
		Type: game.EntityType(a.Type),
		Name: a.Name,
		HP:   a.HP,
		MP:   a.MP,
		Position: game.Position{
			SectorID: a.SectorID,
			Coord: game.Vec3{
				X: a.X,
				Y: a.Y,
				Z: a.Z,
			},
		},
	}
}
