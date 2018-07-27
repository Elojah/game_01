package sector

import (
	"testing"

	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/stretchr/testify/assert"
)

func TestMarshalSector(t *testing.T) {
	ss := []S{
		S{},
		S{ID: ulid.NewID()},
		S{
			ID:  ulid.NewID(),
			Dim: geometry.Vec3{Y: 65.2222},
		},
		S{
			ID:  ulid.NewID(),
			Dim: geometry.Vec3{X: 20.9135, Y: 92.1232, Z: 23.5138},
			Connections: []Connection{
				Connection{},
			},
		},
		S{
			ID:  ulid.NewID(),
			Dim: geometry.Vec3{X: 20.9135, Y: 92.1232, Z: 23.5138},
			Connections: []Connection{
				Connection{
					Coord: geometry.Vec3{X: 48.5803, Y: 35.5452, Z: 77.9984},
				},
			},
		},
		S{
			Dim: geometry.Vec3{X: 20.9135, Y: 92.1232, Z: 23.5138},
			Connections: []Connection{
				Connection{
					Coord: geometry.Vec3{X: 30.2543, Y: 97.582, Z: 88.0111},
					External: geometry.Position{
						SectorID: ulid.NewID(),
						Coord:    geometry.Vec3{X: 80.1429, Y: 76.9509, Z: 51.8441},
					},
				},
				Connection{
					Coord: geometry.Vec3{X: 71.3435, Y: 1.3725, Z: 71.0521},
					External: geometry.Position{
						SectorID: ulid.NewID(),
						Coord:    geometry.Vec3{X: 80.1429, Y: 76.9509, Z: 51.8441},
					},
				},
				Connection{
					Coord: geometry.Vec3{X: 68.6404, Y: 8.14134, Z: 44.3749},
					External: geometry.Position{
						SectorID: ulid.NewID(),
						Coord:    geometry.Vec3{X: 80.1429, Y: 76.9509, Z: 51.8441},
					},
				},
			},
		},
		S{
			ID:  ulid.NewID(),
			Dim: geometry.Vec3{X: 20.9135, Y: 92.1232, Z: 23.5138},
			Connections: []Connection{
				Connection{
					Coord: geometry.Vec3{X: 22.3912, Y: 16.8811, Z: 6.80157},
					External: geometry.Position{
						SectorID: ulid.NewID(),
						Coord:    geometry.Vec3{X: 80.1429, Y: 76.9509, Z: 51.8441},
					},
				},
				Connection{
					Coord: geometry.Vec3{X: 96.5067, Y: 4.25029, Z: 78.0401},
					External: geometry.Position{
						SectorID: ulid.NewID(),
						Coord:    geometry.Vec3{X: 80.1429, Y: 76.9509, Z: 51.8441},
					},
				},
				Connection{
					Coord: geometry.Vec3{X: 6.02444, Y: 27.8391, Z: 8.45019},
					External: geometry.Position{
						SectorID: ulid.NewID(),
						Coord:    geometry.Vec3{X: 80.1429, Y: 76.9509, Z: 51.8441},
					},
				},
			},
		},
	}

	t.Run("marshal/unmarshal", func(t *testing.T) {
		for _, s := range ss {
			raw, err := s.Marshal(nil)
			assert.NoError(t, err)
			var su S
			_, err = su.Unmarshal(raw)
			assert.NoError(t, err)
			assert.Equal(t, su, s)
			_, err = su.UnmarshalSafe(raw)
			assert.NoError(t, err)
			assert.Equal(t, su, s)
		}
	})

	t.Run("unmarshal safe", func(t *testing.T) {
		for _, s := range ss {
			raw, err := s.Marshal(nil)
			assert.NoError(t, err)
			var su S
			for i := 0; i < len(raw); i++ {
				_, err = su.UnmarshalSafe(raw[:i])
				assert.Error(t, err)
			}
		}
	})
}
