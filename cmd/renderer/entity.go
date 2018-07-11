package main

import (
	"github.com/elojah/game_01/pkg/entity"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

// Entity represents a rendered entity.
type Entity struct {
	Position entity.Position
}

// NewEntity returns a new rendered entity from a domain entity.
func NewEntity(e entity.E) Entity {
	return Entity{Position: e.Position}
}

// Draw draws entity e with a new imdraw.
func (e Entity) Draw(imd *imdraw.IMDraw) {
	imd.Color = pixel.RGB(255, 0, 0)
	imd.Push(pixel.V(e.Position.Coord.X, e.Position.Coord.Y))
	imd.Circle(150, 100)
}
