package ability

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// FeedbackStore is the communication interface for ability feedbacks.
type FeedbackStore interface {
	SetFeedback(Feedback) error
	GetFeedback(ulid.ID) (Feedback, error)
}
