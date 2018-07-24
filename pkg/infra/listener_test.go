package infra

import (
	"testing"

	"github.com/elojah/game_01/pkg/ulid"
	"github.com/stretchr/testify/assert"
)

func TestMarshalListener(t *testing.T) {
	ls := []Listener{
		Listener{},
		Listener{
			ID: ulid.NewID(),
		},
		Listener{
			ID:   ulid.NewID(),
			Pool: ulid.NewID(),
		},
		Listener{
			ID:     ulid.NewID(),
			Action: Open,
			Pool:   ulid.NewID(),
		},
		Listener{
			Action: Close,
			Pool:   ulid.NewID(),
		},
	}

	t.Run("marshal/unmarshal", func(t *testing.T) {
		for _, l := range ls {
			raw, err := l.Marshal(nil)
			assert.NoError(t, err)
			var lu Listener
			_, err = lu.Unmarshal(raw)
			assert.NoError(t, err)
			assert.Equal(t, lu, l)
			_, err = lu.UnmarshalSafe(raw)
			assert.NoError(t, err)
			assert.Equal(t, lu, l)
		}
	})

	t.Run("unmarshal safe", func(t *testing.T) {
		for _, l := range ls {
			raw, err := l.Marshal(nil)
			assert.NoError(t, err)
			var lu Listener
			for i := 0; i < len(raw); i++ {
				_, err = lu.UnmarshalSafe(raw[:i])
				assert.Error(t, err)
			}
		}
	})
}
