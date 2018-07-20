package sector

import (
	"math"

	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/ulid"
)

// BondPoint represents a connecting point to another sector.
type BondPoint struct {
	ID       ulid.ID
	SectorID ulid.ID
	Position geometry.Vec3
}

// S represents a cuboid in the world.
type S struct {
	ID         ulid.ID
	Dim        geometry.Vec3
	BondPoints []BondPoint
}

// Out returns if a position is still in the sector.
func (s S) Out(position geometry.Vec3) bool {
	return position.X < 0 ||
		position.X > s.Dim.X ||
		position.Y < 0 ||
		position.Y > s.Dim.Y ||
		position.Z < 0 ||
		position.Z > s.Dim.Z
}

// Adjacents returns ids of adjacent sectors.
func (s S) Adjacents() map[ulid.ID]struct{} {
	sectorIDs := make(map[ulid.ID]struct{})
	for _, bp := range s.BondPoints {
		sectorIDs[bp.SectorID] = struct{}{}
	}
	return sectorIDs
}

// ClosestBP returns the closest bond points in bps of position.
func (s S) ClosestBP(position geometry.Vec3) BondPoint {
	min := math.MaxFloat64
	var iMin int
	for i, bp := range s.BondPoints {
		dist := geometry.Segment(bp.Position, position)
		if dist < min {
			min = dist
			iMin = i
		}
	}
	return s.BondPoints[iMin]
}

// FindBP returns a bond point corresponding to this id for this sector.
func (s S) FindBP(id ulid.ID) BondPoint {
	for _, bp := range s.BondPoints {
		if id == bp.ID {
			return bp
		}
	}
	return BondPoint{}
}

// Mapper is the service for S.
type Mapper interface {
	SetSector(S) error
	GetSector(Subset) (S, error)
}

// Subset allows to retrieve on sector by ID.
type Subset struct {
	ID ulid.ID
}
