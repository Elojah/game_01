package ability

import "github.com/elojah/game_01/pkg/ulid"

// FeedbackStore is the communication interface for ability feedbacks.
type FeedbackStore interface {
	SetFeedback(Feedback) error
	GetFeedback(FeedbackSubset) (Feedback, error)
}

// FeedbackSubset retrieves per ID.
type FeedbackSubset struct {
	ID ulid.ID
}
