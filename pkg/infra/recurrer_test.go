package infra

import (
	"testing"

	"github.com/elojah/game_01/pkg/ulid"
	"github.com/stretchr/testify/assert"
)

func TestMarshalRecurrer(t *testing.T) {
	rs := []Recurrer{
		Recurrer{},
		Recurrer{
			TokenID: ulid.NewID(),
		},
		Recurrer{
			TokenID:  ulid.NewID(),
			EntityID: ulid.NewID(),
			Pool:     ulid.NewID(),
			Action:   Open,
		},
		Recurrer{
			TokenID: ulid.NewID(),
			Pool:    ulid.NewID(),
			Action:  Close,
		},
	}

	t.Run("marshal/unmarshal", func(t *testing.T) {
		for _, r := range rs {
			raw, err := r.Marshal(nil)
			assert.NoError(t, err)
			var ru Recurrer
			_, err = ru.Unmarshal(raw)
			assert.NoError(t, err)
			assert.Equal(t, ru, r)
			_, err = ru.UnmarshalSafe(raw)
			assert.NoError(t, err)
			assert.Equal(t, ru, r)
		}
	})

	t.Run("unmarshal safe", func(t *testing.T) {
		for _, r := range rs {
			raw, err := r.Marshal(nil)
			assert.NoError(t, err)
			var ru Recurrer
			for i := 0; i < len(raw); i++ {
				_, err = ru.UnmarshalSafe(raw[:i])
				assert.Error(t, err)
			}
		}
	})
}
