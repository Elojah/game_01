package event

// Type returns the action type as string.
func (a Action) Type() string {
	if a.Move != nil {
		return "move"
	}
	if a.Cast != nil {
		return "cast"
	}
	if a.Feedback != nil {
		return "feedback"
	}
	if a.Casted != nil {
		return "casted"
	}
	return "null"
}
