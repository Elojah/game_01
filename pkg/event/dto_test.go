package event

import (
	"testing"

	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/stretchr/testify/assert"
)

func TestMarshalDTO(t *testing.T) {
	dtos := []DTO{
		DTO{},
		DTO{
			ID:    ulid.NewID(),
			Token: ulid.NewID(),
			TS:    8400000,
		},
		DTO{
			ID:    ulid.NewID(),
			Token: ulid.NewID(),
			TS:    8400000,
			Action: Move{
				Source: ulid.NewID(),
				Target: ulid.NewID(),
				Position: geometry.Position{
					SectorID: ulid.NewID(),
					Coord: geometry.Vec3{
						X: 31.0613,
						Y: 18.5849,
						Z: 28.4107,
					},
				},
			},
		},
		DTO{
			ID:    ulid.NewID(),
			Token: ulid.NewID(),
			TS:    8400000,
			Action: Cast{
				Source:  ulid.NewID(),
				Targets: []ulid.ID{ulid.NewID(), ulid.NewID(), ulid.NewID()},
				Position: geometry.Position{
					SectorID: ulid.NewID(),
					Coord: geometry.Vec3{
						X: 31.0613,
						Y: 18.5849,
						Z: 28.4107,
					},
				},
			},
		},
	}

	t.Run("marshal/unmarshal", func(t *testing.T) {
		for _, d := range dtos {
			raw, err := d.Marshal(nil)
			assert.NoError(t, err)
			var du DTO
			_, err = du.Unmarshal(raw)
			assert.NoError(t, err)
			assert.Equal(t, du, d)
			_, err = du.UnmarshalSafe(raw)
			assert.NoError(t, err)
			assert.Equal(t, du, d)
		}
	})

	t.Run("unmarshal safe", func(t *testing.T) {
		for _, d := range dtos {
			raw, err := d.Marshal(nil)
			assert.NoError(t, err)
			var du DTO
			for i := 0; i < len(raw); i++ {
				_, err = du.UnmarshalSafe(raw[:i])
				assert.Error(t, err)
			}
		}
	})
}
