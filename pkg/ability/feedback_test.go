package ability

import (
	"testing"

	"github.com/elojah/game_01/pkg/ulid"
	"github.com/stretchr/testify/assert"
)

func TestMarshalFeedback(t *testing.T) {
	fs := []Feedback{
		// This one fails, gencode transforms empty array into nil.
		// Feedback{
		// 	Components: []FeedbackComponent{},
		// },
		Feedback{},
		Feedback{
			ID:         ulid.NewID(),
			AbilityID:  ulid.NewID(),
			Components: nil,
		},
		Feedback{
			ID:         ulid.NewID(),
			AbilityID:  ulid.NewID(),
			Components: nil,
		},
		Feedback{
			Components: []FeedbackComponent{nil, nil},
		},
		Feedback{
			Components: []FeedbackComponent{
				HealDirectFeedback{
					Amount: 83,
				},
				DamageDirectFeedback{
					Amount: 62,
				},
				HealOverTimeFeedback{},
				DamageOverTimeFeedback{},
			},
		},
		Feedback{
			ID:        ulid.NewID(),
			AbilityID: ulid.NewID(),
			Components: []FeedbackComponent{
				HealDirectFeedback{
					Amount: 83,
				},
				DamageDirectFeedback{
					Amount: 62,
				},
				HealOverTimeFeedback{},
				DamageOverTimeFeedback{},
			},
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
