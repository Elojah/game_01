package storage

import (
	"github.com/elojah/game_01"
	"github.com/elojah/game_01/pkg/entity"
)

// NewEntity convert a entity.E into a storage Entity.
func NewEntity(e entity.E) *Entity {
	return &Entity{
		ID:       [16]byte(e.ID),
		Type:     [16]byte(e.Type),
		Name:     e.Name,
		HP:       e.HP,
		MP:       e.MP,
		SectorID: e.Position.SectorID,
		X:        e.Position.Coord.X,
		Y:        e.Position.Coord.Y,
		Z:        e.Position.Coord.Z,
	}
}

// Domain converts a storage Entity into a game Entity.
func (e Entity) Domain() entity.E {
	return entity.E{
		ID:   game.ID(e.ID),
		Type: entity.Type(e.Type),
		Name: e.Name,
		HP:   e.HP,
		MP:   e.MP,
		Position: entity.Position{
			SectorID: e.SectorID,
			Coord: game.Vec3{
				X: e.X,
				Y: e.Y,
				Z: e.Z,
			},
		},
	}
}
