package sector

import (
	"testing"

	"github.com/elojah/game_01/pkg/entity"
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
			BondPoints: []BondPoint{
				BondPoint{},
			},
		},
		S{
			ID:  ulid.NewID(),
			Dim: geometry.Vec3{X: 20.9135, Y: 92.1232, Z: 23.5138},
			BondPoints: []BondPoint{
				BondPoint{
					ID: ulid.NewID(),
				},
			},
		},
		S{
			Dim: geometry.Vec3{X: 20.9135, Y: 92.1232, Z: 23.5138},
			BondPoints: []BondPoint{
				BondPoint{
					ID: ulid.NewID(),
					Position: entity.Position{
						SectorID: ulid.NewID(),
						Coord:    geometry.Vec3{X: 80.1429, Y: 76.9509, Z: 51.8441},
					},
				},
				BondPoint{
					ID: ulid.NewID(),
					Position: entity.Position{
						SectorID: ulid.NewID(),
						Coord:    geometry.Vec3{X: 80.1429, Y: 76.9509, Z: 51.8441},
					},
				},
				BondPoint{
					ID: ulid.NewID(),
					Position: entity.Position{
						SectorID: ulid.NewID(),
						Coord:    geometry.Vec3{X: 80.1429, Y: 76.9509, Z: 51.8441},
					},
				},
			},
		},
		S{
			ID:  ulid.NewID(),
			Dim: geometry.Vec3{X: 20.9135, Y: 92.1232, Z: 23.5138},
			BondPoints: []BondPoint{
				BondPoint{
					ID: ulid.NewID(),
					Position: entity.Position{
						SectorID: ulid.NewID(),
						Coord:    geometry.Vec3{X: 80.1429, Y: 76.9509, Z: 51.8441},
					},
				},
				BondPoint{
					ID: ulid.NewID(),
					Position: entity.Position{
						SectorID: ulid.NewID(),
						Coord:    geometry.Vec3{X: 80.1429, Y: 76.9509, Z: 51.8441},
					},
				},
				BondPoint{
					ID: ulid.NewID(),
					Position: entity.Position{
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
