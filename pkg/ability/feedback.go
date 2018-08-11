package ability

import "github.com/elojah/game_01/pkg/ulid"

// FeedbackMapper is the communication interface for ability feedbacks.
type FeedbackMapper interface {
	SetAbilityFeedback(Feedback) error
	GetAbilityFeedback(FeedbackSubset) (Feedback, error)
}

// FeedbackSubset retrieves per ID.
type FeedbackSubset struct {
	ID ulid.ID
}
