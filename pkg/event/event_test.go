package event

import (
	"testing"
	"time"

	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/stretchr/testify/assert"
)

func TestMarshalEvent(t *testing.T) {
	es := []E{
		// UnixNano on 0 time is undefined
		E{TS: time.Now().Round(-1)},
		E{
			ID:     ulid.NewID(),
			Source: ulid.NewID(),
			TS:     time.Now().Round(-1),
		},
		E{
			ID:     ulid.NewID(),
			Source: ulid.NewID(),
			TS:     time.Now().Round(-1),
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
		E{
			ID:     ulid.NewID(),
			Source: ulid.NewID(),
			TS:     time.Now().Round(-1),
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
		E{
			ID:     ulid.NewID(),
			Source: ulid.NewID(),
			TS:     time.Now().Round(-1),
			Action: Feedback{
				ID:     ulid.NewID(),
				Source: ulid.NewID(),
				Target: ulid.NewID(),
			},
		},
	}

	t.Run("marshal/unmarshal", func(t *testing.T) {
		for _, e := range es {
			raw, err := e.Marshal(nil)
			assert.NoError(t, err)
			var eu E
			_, err = eu.Unmarshal(raw)
			assert.NoError(t, err)
			assert.Equal(t, eu, e)
			_, err = eu.UnmarshalSafe(raw)
			assert.NoError(t, err)
			assert.Equal(t, eu, e)
		}
	})

	t.Run("unmarshal safe", func(t *testing.T) {
		for _, e := range es {
			raw, err := e.Marshal(nil)
			assert.NoError(t, err)
			var eu E
			for i := 0; i < len(raw); i++ {
				_, err = eu.UnmarshalSafe(raw[:i])
				assert.Error(t, err)
			}
		}
	})
}
