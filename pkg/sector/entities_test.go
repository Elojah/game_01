package sector

import (
	"testing"

	"github.com/elojah/game_01/pkg/ulid"
	"github.com/stretchr/testify/assert"
)

func TestMarshalEntities(t *testing.T) {
	es := []Entities{
		Entities{},
		Entities{SectorID: ulid.NewID()},
		Entities{
			SectorID: ulid.NewID(),
			EntityIDs: []ulid.ID{
				ulid.NewID(),
				ulid.NewID(),
				ulid.NewID(),
				ulid.NewID(),
				ulid.NewID(),
				ulid.NewID(),
			},
		},
	}

	t.Run("marshal/unmarshal", func(t *testing.T) {
		for _, e := range es {
			raw, err := e.Marshal(nil)
			assert.NoError(t, err)
			var eu Entities
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
			var eu Entities
			for i := 0; i < len(raw); i++ {
				_, _ = eu.UnmarshalSafe(raw[:i])
				// assert.Error(t, err)
			}
		}
	})
}
