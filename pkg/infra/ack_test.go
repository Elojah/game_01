package infra

import (
	"testing"

	"github.com/elojah/game_01/pkg/ulid"
	"github.com/stretchr/testify/assert"
)

func TestMarshalACK(t *testing.T) {
	as := []ACK{
		ACK{},
		ACK{ID: ulid.NewID()},
	}

	t.Run("marshal/unmarshal", func(t *testing.T) {
		for _, a := range as {
			raw, err := a.Marshal(nil)
			assert.NoError(t, err)
			var au ACK
			_, err = au.Unmarshal(raw)
			assert.NoError(t, err)
			assert.Equal(t, au, a)
			_, err = au.UnmarshalSafe(raw)
			assert.NoError(t, err)
			assert.Equal(t, au, a)
		}
	})

	t.Run("unmarshal safe", func(t *testing.T) {
		for _, a := range as {
			raw, err := a.Marshal(nil)
			assert.NoError(t, err)
			var au ACK
			for i := 0; i < len(raw); i++ {
				_, err = au.UnmarshalSafe(raw[:i])
				assert.Error(t, err)
			}
		}
	})
}
