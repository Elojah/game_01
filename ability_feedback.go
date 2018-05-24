package game

// AbilityFeedback represents the effects a ability had on a target.
type AbilityFeedback struct {
	ID         ID
	AbilityID  ID
	Components []AbilityFeedbackComponent
}

// AbilityFeedbackMapper is the communication interface for ability feedbacks.
type AbilityFeedbackMapper interface {
	SetAbilityFeedback(AbilityFeedback) error
	GetAbilityFeedback(AbilityFeedbackSubset) (AbilityFeedback, error)
}

// AbilityFeedbackSubset retrieves per ID.
type AbilityFeedbackSubset struct {
	ID ID
}
