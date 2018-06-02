package storage

import (
	"github.com/elojah/game_01"
)

// Domain converts a storage sector into a domain sector.
func (s *Sector) Domain() game.Sector {
	eps := [26][]game.ExitPoint{}
	for i := 0; i < 26; i++ {
		eps[i] = make([]game.ExitPoint, len(s.ExitPoints[i]))
		for j := range s.ExitPoints[i] {
			eps[i][j] = game.ExitPoint{
				ID:       game.ID(s.ExitPoints[i][j].ID),
				SectorID: game.ID(s.ExitPoints[i][j].SectorID),
				Position: game.Vec3{
					X: s.ExitPoints[i][j].X,
					Y: s.ExitPoints[i][j].Y,
					Z: s.ExitPoints[i][j].Z,
				},
			}
		}
	}
	return game.Sector{
		ID: game.ID(s.ID),
		Size: game.Vec3{
			X: s.X,
			Y: s.Y,
			Z: s.Z,
		},
		ExitPoints: eps,
	}
}

// NewSector converts a domain sector into a storage sector.
func NewSector(sector game.Sector) *Sector {
	eps := [26][]ExitPoint{}
	for i := 0; i < 26; i++ {
		eps[i] = make([]ExitPoint, len(sector.ExitPoints[i]))
		for j := range sector.ExitPoints[i] {
			eps[i][j] = ExitPoint{
				ID:       [16]byte(sector.ExitPoints[i][j].ID),
				SectorID: [16]byte(sector.ExitPoints[i][j].SectorID),
				X:        sector.ExitPoints[i][j].Position.X,
				Y:        sector.ExitPoints[i][j].Position.Y,
				Z:        sector.ExitPoints[i][j].Position.Z,
			}
		}
	}
	return &Sector{
		ID:         [16]byte(sector.ID),
		X:          sector.Size.X,
		Y:          sector.Size.Y,
		Z:          sector.Size.Z,
		ExitPoints: eps,
	}
}
