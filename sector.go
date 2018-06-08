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

// Adjacents returns ids of adjacent sectors.
func (s Sector) Adjacents() []ID {
	sectorIDsMap := make(map[ID]struct{})
	for _, bp := range s.BondPoints {
		sectorIDsMap[bp.SectorID] = struct{}{}
	}
	sectorIDs := make ([]ID, len(sectorIDsMap))
	for i := range sectorIDsMap {
		sectorIDs[i] = sectorIDsMap[i]
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

// SectorMapper is the service for Sector.
type SectorMapper interface {
	SetSector(Sector) error
	GetSector(SectorSubset) (Sector, error)
}

// SectorSubset allows to retrieve on sector by ID.
type SectorSubset struct {
	ID ID
}
