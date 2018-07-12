package main

import (
	"math/rand"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

// Entity represents a rendered entity.
type Entity struct {
	Position entity.Position
	color    pixel.RGBA
}

// NewEntity returns a new rendered entity from a domain entity.
func NewEntity(e entity.E) Entity {
	return Entity{
		Position: e.Position,
		color: pixel.RGB(
			float64(rand.Int()%255),
			float64(rand.Int()%255),
			float64(rand.Int()%255),
		),
	}
}

// Draw draws entity e with a new imdraw.
func (e Entity) Draw(imd *imdraw.IMDraw) {
	imd.Color = e.color
	imd.Push(pixel.V(e.Position.Coord.X, e.Position.Coord.Y))
	imd.Circle(10, 5)
}
