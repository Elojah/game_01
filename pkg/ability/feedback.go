package ability

import (
	game "github.com/elojah/game_01"
)

// Feedback represents the effects a ability had on a target.
type Feedback struct {
	ID         game.ID
	AbilityID  game.ID
	Components []FeedbackComponent
}

// FeedbackMapper is the communication interface for ability feedbacks.
type FeedbackMapper interface {
	SetAbilityFeedback(Feedback) error
	GetAbilityFeedback(FeedbackSubset) (Feedback, error)
}

// FeedbackSubset retrieves per ID.
type FeedbackSubset struct {
	ID game.ID
}
