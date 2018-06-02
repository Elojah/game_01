package game

import (
	"math"
)

// Vec2 is a 2D position. Coordinates are in float64.
type Vec2 struct {
	X float64
	Y float64
}

// Vec3 is a 3D position. Coordinates are in float64.
type Vec3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// Add add vector coordinates to v.
func (v *Vec3) Add(p Vec3) {
	v.X += p.X
	v.Y += p.Y
	v.Z += p.Z
}

// Segment returns the distance between 2 points following XYZ axis.
func Segment(lhs Vec3, rhs Vec3) float64 {
	return math.Abs(lhs.X-rhs.X) + math.Abs(lhs.Y-rhs.Y) + math.Abs(lhs.Z-rhs.Z)
}
