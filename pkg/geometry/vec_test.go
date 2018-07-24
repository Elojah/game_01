package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalVec3(t *testing.T) {
	vs := []Vec3{
		// UnixNano on 0 time is undefined
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
