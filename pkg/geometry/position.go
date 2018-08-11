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

// MoveReference moves v to a new reference. lhs and rhs must refer to a identical point in both references.
func (v *Vec3) MoveReference(lhs Vec3, rhs Vec3) {
	v.X += rhs.X - lhs.X
	v.Y += rhs.Y - lhs.Y
	v.Z += rhs.Z - lhs.Z
}

// Segment returns the distance between 2 points following XYZ axis.
func Segment(lhs Vec3, rhs Vec3) float64 {
	return math.Abs(lhs.X-rhs.X) + math.Abs(lhs.Y-rhs.Y) + math.Abs(lhs.Z-rhs.Z)
}
