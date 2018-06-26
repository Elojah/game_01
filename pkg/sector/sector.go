package sector

import (
	"math"

	game "github.com/elojah/game_01"
)

// BondPoint represents a connecting point to another sector.
type BondPoint struct {
	ID       game.ID
	SectorID game.ID
	Position game.Vec3
}

// S represents a cuboid in the world.
type S struct {
	ID         game.ID
	Size       game.Vec3
	BondPoints []BondPoint
}

// Out returns if a position is still in the sector.
func (s S) Out(position game.Vec3) bool {
	return position.X < 0 ||
		position.X > s.Size.X ||
		position.Y < 0 ||
		position.Y > s.Size.Y ||
		position.Z < 0 ||
		position.Z > s.Size.Z
}

// Adjacents returns ids of adjacent sectors.
func (s S) Adjacents() map[game.ID]struct{} {
	sectorIDs := make(map[game.ID]struct{})
	for _, bp := range s.BondPoints {
		sectorIDs[bp.SectorID] = struct{}{}
	}
	return sectorIDs
}

// ClosestBP returns the closest bond points in bps of position.
func (s S) ClosestBP(position game.Vec3) BondPoint {
	min := math.MaxFloat64
	var iMin int
	for i, bp := range s.BondPoints {
		dist := game.Segment(bp.Position, position)
		if dist < min {
			min = dist
			iMin = i
		}
	}
	return s.BondPoints[iMin]
}

// FindBP returns a bond point corresponding to this id for this sector.
func (s S) FindBP(id game.ID) BondPoint {
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
	ID game.ID
}
