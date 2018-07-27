package geometry

import (
	"testing"

	"github.com/elojah/game_01/pkg/ulid"
	"github.com/stretchr/testify/assert"
)

func TestMarshalPosition(t *testing.T) {
	vs := []Position{
		Position{},
		Position{SectorID: ulid.NewID()},
		Position{
			Coord: Vec3{
				X: 5.70347,
				Y: 43.4246,
				Z: 1.07845,
			},
		},
		Position{
			SectorID: ulid.NewID(),
			Coord: Vec3{
				X: 5.70347,
				Y: 43.4246,
				Z: 1.07845,
			},
		},
		Position{
			SectorID: ulid.NewID(),
			Coord: Vec3{
				Y: 43.4246,
			},
		},
	}

	t.Run("marshal/unmarshal", func(t *testing.T) {
		for _, p := range vs {
			raw, err := p.Marshal(nil)
			assert.NoError(t, err)
			var pu Position
			_, err = pu.Unmarshal(raw)
			assert.NoError(t, err)
			assert.Equal(t, pu, p)
			_, err = pu.UnmarshalSafe(raw)
			assert.NoError(t, err)
			assert.Equal(t, pu, p)
		}
	})

	t.Run("unmarshal safe", func(t *testing.T) {
		for _, p := range vs {
			raw, err := p.Marshal(nil)
			assert.NoError(t, err)
			var pu Position
			for i := 0; i < len(raw); i++ {
				_, err = pu.UnmarshalSafe(raw[:i])
				assert.Error(t, err)
			}
		}
	})
}

func TestMarshalVec3(t *testing.T) {
	vs := []Vec3{
		Vec3{},
		Vec3{Y: 79.7871},
		Vec3{
			X: 5.70347,
			Y: 43.4246,
			Z: 1.07845,
		},
	}

	t.Run("marshal/unmarshal", func(t *testing.T) {
		for _, v := range vs {
			raw, err := v.Marshal(nil)
			assert.NoError(t, err)
			var vu Vec3
			_, err = vu.Unmarshal(raw)
			assert.NoError(t, err)
			assert.Equal(t, vu, v)
			_, err = vu.UnmarshalSafe(raw)
			assert.NoError(t, err)
			assert.Equal(t, vu, v)
		}
	})

	t.Run("unmarshal safe", func(t *testing.T) {
		for _, v := range vs {
			raw, err := v.Marshal(nil)
			assert.NoError(t, err)
			var vu Vec3
			for i := 0; i < len(raw); i++ {
				_, err = vu.UnmarshalSafe(raw[:i])
				assert.Error(t, err)
			}
		}
	})
}
