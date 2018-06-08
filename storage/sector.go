package storage

import (
	"github.com/elojah/game_01"
)

// Domain converts a storage sector into a domain sector.
func (s *Sector) Domain() game.Sector {
	bps := make([]game.BondPoint, len(s.BondPoints))
	for i := range s.BondPoints {
		bps[i] = game.BondPoint{
			ID:       game.ID(s.BondPoints[i].ID),
			SectorID: game.ID(s.BondPoints[i].SectorID),
			Position: game.Vec3{
				X: s.BondPoints[i].X,
				Y: s.BondPoints[i].Y,
				Z: s.BondPoints[i].Z,
			},
		}
	}
	return game.Sector{
		ID: game.ID(s.ID),
		Size: game.Vec3{
			X: s.X,
			Y: s.Y,
			Z: s.Z,
		},
		BondPoints: bps,
	}
}

// NewSector converts a domain sector into a storage sector.
func NewSector(sector game.Sector) *Sector {
	bps := make([]BondPoint, len(sector.BondPoints))
	for i := range sector.BondPoints {
		bps[i] = BondPoint{
			ID:       [16]byte(sector.BondPoints[i].ID),
			SectorID: [16]byte(sector.BondPoints[i].SectorID),
			X:        sector.BondPoints[i].Position.X,
			Y:        sector.BondPoints[i].Position.Y,
			Z:        sector.BondPoints[i].Position.Z,
		}
	}
	return &Sector{
		ID:         [16]byte(sector.ID),
		X:          sector.Size.X,
		Y:          sector.Size.Y,
		Z:          sector.Size.Z,
		BondPoints: bps,
	}
}
