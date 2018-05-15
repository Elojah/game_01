package storage

import (
	"github.com/elojah/game_01"
)

// NewEntity convert a game.Entity into a storage Entity.
func NewEntity(a game.Entity) *Entity {
	return &Entity{
		ID: [16]byte(a.ID),
		HP: a.HP,
		MP: a.MP,
		X:  a.Position.X,
		Y:  a.Position.Y,
		Z:  a.Position.Z,
	}
}

// Domain converts a storage Entity into a game Entity.
func (a Entity) Domain() game.Entity {
	return game.Entity{
		ID: game.ID(a.ID),
		HP: a.HP,
		MP: a.MP,
		Position: game.Vec3{
			X: a.X,
			Y: a.Y,
			Z: a.Z,
		},
	}
}
