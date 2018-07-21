package ability

import (
	"testing"

	"github.com/elojah/game_01/pkg/ulid"
	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	as := []A{
		A{},
		A{
			ID:            ulid.NewID(),
			Type:          ulid.NewID(),
			Name:          "some random quite long fucking name",
			MPConsumption: 23,
			CD:            97,
			CurrentCD:     24,
			Components:    nil,
		},
		// This one fails, gencode transforms empty array into nil.
		// A{
		// 	Components: []Component{},
		// },
		A{
			Components: []Component{nil, nil},
		},
		A{
			Components: []Component{
				HealDirect{
					Amount: 83,
					Type:   44,
				},
				DamageDirect{
					Amount: 62,
					Type:   53,
				},
				HealOverTime{
					Amount:    52,
					Type:      96,
					Frequency: 9,
					Duration:  57,
				},
				DamageOverTime{
					Amount:    71,
					Type:      48,
					Frequency: 48,
					Duration:  89,
				},
			},
		},
	}
	t.Run("marshal/unmarshal", func(t *testing.T) {
		for _, a := range as {
			raw, err := a.Marshal(nil)
			assert.NoError(t, err)
			var au A
			_, err = au.Unmarshal(raw)
			assert.NoError(t, err)
			assert.Equal(t, au, a)
		}
	})
	t.Run("unmarshal safe", func(t *testing.T) {
		for _, a := range as {
			raw, err := a.Marshal(nil)
			assert.NoError(t, err)
			var au A
			_, err = au.UnmarshalSafe(raw[:len(raw)-4])
			assert.Error(t, err)
			_, err = au.UnmarshalSafe(raw[:len(raw)-8])
			assert.Error(t, err)
			_, err = au.UnmarshalSafe(raw[:len(raw)-12])
			assert.Error(t, err)
			_, err = au.UnmarshalSafe(raw[:len(raw)-16])
			assert.Error(t, err)
			_, err = au.UnmarshalSafe(raw[:len(raw)-20])
			assert.Error(t, err)
			_, err = au.UnmarshalSafe(raw[:len(raw)-24])
			assert.Error(t, err)
			_, err = au.UnmarshalSafe(raw[:len(raw)-28])
			assert.Error(t, err)
			_, err = au.UnmarshalSafe(raw[:len(raw)-32])
			assert.Error(t, err)
			_, err = au.UnmarshalSafe(raw[:len(raw)-36])
			assert.Error(t, err)
			_, err = au.UnmarshalSafe(raw[:len(raw)-40])
			assert.Error(t, err)
			_, err = au.UnmarshalSafe(raw[:len(raw)-44])
			assert.Error(t, err)
			_, err = au.UnmarshalSafe(raw[:len(raw)-48])
			assert.Error(t, err)
		}
	})
}
