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

// Cuboid represents a square box.
type Cuboid struct {
	X    int64
	Y    int64
	Z    int64
	Size Vec3
}

// Circle represents a area circle.
type Circle struct {
	Centre Vec2
	Radius float64
}

// Segment returns the distance between 2 points following XYZ axis.
func Segment(lhs Vec3, rhs Vec3) float64 {
	return math.Abs(lhs.X-rhs.X) + math.Abs(lhs.Y-rhs.Y) + math.Abs(lhs.Z-rhs.Z)
}
