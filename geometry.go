package game

// Vec2 is a 2D position. Coordinates are in float64.
type Vec2 struct {
	X float64
	Y float64
}

// Vec3 is a 3D position. Coordinates are in float64.
type Vec3 struct {
	X float64
	Y float64
	Z float64
}

// Circle represents a area circle.
type Circle struct {
	Centre Vec2
	Radius float64
}
