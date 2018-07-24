package event

import (
	"testing"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/stretchr/testify/assert"
)

func TestMarshalMove(t *testing.T) {
	ms := []Move{
		Move{},
		Move{
			Source: ulid.NewID(),
			Target: ulid.NewID(),
			Position: entity.Position{
				SectorID: ulid.NewID(),
				Coord: geometry.Vec3{
					X: 31.0613,
					Y: 18.5849,
					Z: 28.4107,
				},
			},
		},
	}

	t.Run("marshal/unmarshal", func(t *testing.T) {
		for _, m := range ms {
			raw, err := m.Marshal(nil)
			assert.NoError(t, err)
			var mu Move
			_, err = mu.Unmarshal(raw)
			assert.NoError(t, err)
			assert.Equal(t, mu, m)
		}
	})

	t.Run("unmarshal safe", func(t *testing.T) {
		for _, m := range ms {
			raw, err := m.Marshal(nil)
			assert.NoError(t, err)
			var mu Move
			for i := 0; i < len(raw); i++ {
				_, err = mu.UnmarshalSafe(raw[:i])
				assert.Error(t, err)
			}
		}
	})
}

func TestMarshalCast(t *testing.T) {
	cs := []Cast{
		Cast{},
		Cast{
			Source:  ulid.NewID(),
			Targets: []ulid.ID{ulid.NewID(), ulid.NewID(), ulid.NewID()},
			Position: entity.Position{
				SectorID: ulid.NewID(),
				Coord: geometry.Vec3{
					X: 31.0613,
					Y: 18.5849,
					Z: 28.4107,
				},
			},
		},
		Cast{
			AbilityID: ulid.NewID(),
			Source:    ulid.NewID(),
			Targets:   []ulid.ID{ulid.NewID()},
			Position:  entity.Position{},
		},
		Cast{
			AbilityID: ulid.NewID(),
			Source:    ulid.NewID(),
			Targets:   nil,
			Position:  entity.Position{},
		},
	}

	t.Run("marshal/unmarshal", func(t *testing.T) {
		for _, m := range cs {
			raw, err := m.Marshal(nil)
			assert.NoError(t, err)
			var mu Cast
			_, err = mu.Unmarshal(raw)
			assert.NoError(t, err)
			assert.Equal(t, mu, m)
			_, err = mu.UnmarshalSafe(raw)
			assert.NoError(t, err)
			assert.Equal(t, mu, m)
		}
	})

	t.Run("unmarshal safe", func(t *testing.T) {
		for _, m := range cs {
			raw, err := m.Marshal(nil)
			assert.NoError(t, err)
			var mu Cast
			for i := 0; i < len(raw); i++ {
				_, err = mu.UnmarshalSafe(raw[:i])
				assert.Error(t, err)
			}
		}
	})
}

func TestMarshalFeedback(t *testing.T) {
	fs := []Feedback{
		Feedback{},
		Feedback{
			Source: ulid.NewID(),
		},
		Feedback{
			ID:     ulid.NewID(),
			Source: ulid.NewID(),
			Target: ulid.NewID(),
		},
	}

	t.Run("marshal/unmarshal", func(t *testing.T) {
		for _, f := range fs {
			raw, err := f.Marshal(nil)
			assert.NoError(t, err)
			var fu Feedback
			_, err = fu.Unmarshal(raw)
			assert.NoError(t, err)
			assert.Equal(t, fu, f)
			_, err = fu.UnmarshalSafe(raw)
			assert.NoError(t, err)
			assert.Equal(t, fu, f)
		}
	})

	t.Run("unmarshal safe", func(t *testing.T) {
		for _, f := range fs {
			raw, err := f.Marshal(nil)
			assert.NoError(t, err)
			var fu Feedback
			for i := 0; i < len(raw); i++ {
				_, err = fu.UnmarshalSafe(raw[:i])
				assert.Error(t, err)
			}
		}
	})
}
