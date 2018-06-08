package game

import (
	"math"
)

// BondPoint represents a connecting point to another sector.
type BondPoint struct {
	ID       ID
	SectorID ID
	Position Vec3
}

// Sector represents a cuboid in the world.
type Sector struct {
	ID         ID
	Size       Vec3
	BondPoints []BondPoint
}

// Out returns if a position is still in the sector.
func (s Sector) Out(position Vec3) bool {
	return position.X < 0 ||
		position.X > s.Size.X ||
		position.Y < 0 ||
		position.Y > s.Size.Y ||
		position.Z < 0 ||
		position.Z > s.Size.Z
}

// Adjacents returns ids of adjacent sectors.
func (s Sector) Adjacents() map[ID]struct{} {
	sectorIDs := make(map[ID]struct{})
	for _, bp := range s.BondPoints {
		sectorIDs[bp.SectorID] = struct{}{}
	}
	return sectorIDs
}

// ClosestBP returns the closest bond points in bps of position.
func (s Sector) ClosestBP(position Vec3) BondPoint {
	min := math.MaxFloat64
	var iMin int
	for i, bp := range s.BondPoints {
		dist := Segment(bp.Position, position)
		if dist < min {
			min = dist
			iMin = i
		}
	}
	return s.BondPoints[iMin]
}

// FindBP returns a bond point corresponding to this id for this sector.
func (s Sector) FindBP(id ID) BondPoint {
	for _, bp := range s.BondPoints {
		if id == bp.ID {
			return bp
		}
	}
	return BondPoint{}
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
