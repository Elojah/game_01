package geometry

import (
	"math"
)

// Add add vector coordinates to v.
func (v *Vec3) Add(p Vec3) {
	v.X += p.X
	v.Y += p.Y
	v.Z += p.Z
}

// MoveReference returns v position moved to a new reference ref relative to previous origin.
func (v Vec3) MoveReference(ref Vec3) Vec3 {
	return Vec3{
		X: v.X + ref.X,
		Y: v.Y + ref.Y,
		Z: v.Z + ref.Z,
	}
}

// Segment returns the distance between 2 points following XYZ axis.
func Segment(lhs Vec3, rhs Vec3) float64 {
	return math.Abs(lhs.X-rhs.X) + math.Abs(lhs.Y-rhs.Y) + math.Abs(lhs.Z-rhs.Z)
}
