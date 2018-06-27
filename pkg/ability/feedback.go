package ability

import "github.com/elojah/game_01/pkg/ulid"

// Feedback represents the effects a ability had on a target.
type Feedback struct {
	ID         ulid.ID
	AbilityID  ulid.ID
	Components []FeedbackComponent
}

// FeedbackMapper is the communication interface for ability feedbacks.
type FeedbackMapper interface {
	SetAbilityFeedback(Feedback) error
	GetAbilityFeedback(FeedbackSubset) (Feedback, error)
}

// FeedbackSubset retrieves per ID.
type FeedbackSubset struct {
	ID ulid.ID
}
