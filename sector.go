package game

import (
	"math"
)

// Direction represents a direction relative to a sector. Must have 26 values + 1 neutral.
type Direction uint8

const (
	// In indicates the point is still in the sector.
	In Direction = 100
	// Back is the back direction relative to sector.
	Back Direction = 0
	// Down is the down direction relative to sector.
	Down Direction = 1
	// DownBack is the down back direction relative to sector.
	DownBack Direction = 2
	// DownFront is the down front direction relative to sector.
	DownFront Direction = 3
	// Front is the front direction relative to sector.
	Front Direction = 4
	// Left is the left direction relative to sector.
	Left Direction = 5
	// LeftBack is the left back direction relative to sector.
	LeftBack Direction = 6
	// LeftDown is the down left direction relative to sector.
	LeftDown Direction = 7
	// LeftDownBack is the down left back direction relative to sector.
	LeftDownBack Direction = 8
	// LeftDownFront is the down left front direction relative to sector.
	LeftDownFront Direction = 9
	// LeftFront is the left front direction relative to sector.
	LeftFront Direction = 10
	// LeftUp is the up left direction relative to sector.
	LeftUp Direction = 11
	// LeftUpBack is the up left back direction relative to sector.
	LeftUpBack Direction = 12
	// LeftUpFront is the up left front direction relative to sector.
	LeftUpFront Direction = 13
	// Right is the right direction relative to sector.
	Right Direction = 14
	// RightBack is the right back direction relative to sector.
	RightBack Direction = 15
	// RightDown is the down right direction relative to sector.
	RightDown Direction = 16
	// RightDownBack is the down right back direction relative to sector.
	RightDownBack Direction = 17
	// RightDownFront is the down right front direction relative to sector.
	RightDownFront Direction = 18
	// RightFront is the right front direction relative to sector.
	RightFront Direction = 19
	// RightUp is the up right direction relative to sector.
	RightUp Direction = 20
	// RightUpBack is the up right back direction relative to sector.
	RightUpBack Direction = 21
	// RightUpFront is the up right front direction relative to sector.
	RightUpFront Direction = 22
	// Up is the up direction relative to sector.
	Up Direction = 23
	// UpBack is the up back direction relative to sector.
	UpBack Direction = 24
	// UpFront is the up front direction relative to sector.
	UpFront Direction = 25
)

// Opposite returns the opposite direction of d.
func (d Direction) Opposite() Direction {
	switch d {
	case Back:
		return Front
	case Down:
		return Up
	case DownBack:
		return UpFront
	case DownFront:
		return UpBack
	case Front:
		return Back
	case Left:
		return Right
	case LeftBack:
		return RightFront
	case LeftDown:
		return RightUp
	case LeftDownBack:
		return RightUpFront
	case LeftDownFront:
		return RightUpBack
	case LeftFront:
		return RightBack
	case LeftUp:
		return RightDown
	case LeftUpBack:
		return RightDownFront
	case LeftUpFront:
		return RightDownBack
	case Right:
		return Left
	case RightBack:
		return LeftFront
	case RightDown:
		return LeftUp
	case RightDownBack:
		return LeftUpFront
	case RightDownFront:
		return LeftUpBack
	case RightFront:
		return LeftBack
	case RightUp:
		return LeftDown
	case RightUpBack:
		return LeftDownFront
	case RightUpFront:
		return LeftDownBack
	case Up:
		return Down
	case UpBack:
		return DownFront
	case UpFront:
		return DownBack
	}
	// Error case
	return In
}

// ExitPoint represents a central point to another sector.
type ExitPoint struct {
	ID       ID
	SectorID ID
	Position Vec3
}

// ExitPoints alias a slice of ExitPoint
type ExitPoints []ExitPoint

// Closest returns the closest exit points in exps of position.
func (exps ExitPoints) Closest(position Vec3) ExitPoint {
	min := math.MaxFloat64
	var iMin int
	for i, exp := range exps {
		dist := Segment(exp.Position, position)
		if dist < min {
			min = dist
			iMin = i
		}
	}
	return exps[iMin]
}

// Sector represents a cuboid in the world.
type Sector struct {
	ID         ID
	Size       Vec3
	ExitPoints [26][]ExitPoint
}

// Direction returns the relative position of v for sector s.
func (s Sector) Direction(v Vec3) Direction {
	deltaX := In
	if v.X > s.Size.X {
		deltaX = Right
	} else if v.X < 0 {
		deltaX = Left
	}
	deltaY := In
	if v.Y > s.Size.Y {
		deltaY = Up
	} else if v.Y < 0 {
		deltaY = Down
	}
	deltaZ := In
	if v.Z > s.Size.Z {
		deltaZ = Front
	} else if v.Z < 0 {
		deltaZ = Back
	}
	return combineDirections(deltaX, deltaY, deltaZ)
}

// TODO clean this with arrays.
func combineDirections(x Direction, y Direction, z Direction) Direction {
	if x == Right {
		if y == Up {
			if z == Front {
				return RightUpFront
			} else if z == Back {
				return RightUpBack
			}
			return RightUp
		} else if y == Down {
			if z == Front {
				return RightDownFront
			} else if z == Back {
				return RightDownBack
			}
			return RightDown
		}
		if z == Front {
			return RightFront
		} else if z == Back {
			return RightBack
		}
		return Right
	} else if x == Left {
		if y == Up {
			if z == Front {
				return LeftUpFront
			} else if z == Back {
				return LeftUpBack
			}
			return LeftUp
		} else if y == Down {
			if z == Front {
				return LeftDownFront
			} else if z == Back {
				return LeftDownBack
			}
			return LeftDown
		}
		if z == Front {
			return LeftFront
		} else if z == Back {
			return LeftBack
		}
		return Left
	}
	if y == Up {
		if z == Front {
			return UpFront
		} else if z == Back {
			return UpBack
		}
		return Up
	} else if y == Down {
		if z == Front {
			return DownFront
		} else if z == Back {
			return DownBack
		}
		return Down
	}
	if z == Front {
		return Front
	} else if z == Back {
		return Back
	}
	return In
}

// GetExitPoint returns the exit point in direction d with id id.
func (s Sector) GetExitPoint(id ID, d Direction) ExitPoint {
	for _, exp := range s.ExitPoints[d] {
		if exp.ID.Compare(id) == 0 {
			return exp
		}
	}
	return ExitPoint{}
}

// SectorMapper is the service for Sector.
type SectorMapper interface {
	SetSector(Sector) error
	GetSector(SectorSubset) (Sector, error)
}

// SectorSubset allows to retrieve on sector by ID.
type SectorSubset struct {
	ID ID
}
