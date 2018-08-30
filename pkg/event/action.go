package event

// Type returns the action type as string.
func (a Action) Type() string {

	if a.MoveSource != nil {
		return "move_source"
	}
	if a.MoveTarget != nil {
		return "move_target"
	}
	if a.CastSource != nil {
		return "cast_source"
	}
	if a.PerformSource != nil {
		return "perform_source"
	}
	if a.PerformTarget != nil {
		return "perform_target"
	}
	if a.FeedbackTarget != nil {
		return "feedback_target"
	}
	return "unknown"
}
