package game

// Action is a client action.
type Action interface{}

// ActionString acts for action.String() in declarative mode.
func ActionString(a Action) string {
	switch a.(type) {
	case Move:
		return "move"
	case Cast:
		return "cast"
	case Feedback:
		return "feedback"
	default:
		return "unknown"
	}
}

// Move represents a unit move.
type Move struct {
	Source   ID
	Target   ID
	Position Vec3
}

// Cast represents a ability launch.
type Cast struct {
	AbilityID ID
	Source    ID
	Targets   []ID
	Position  Vec3
}

// Feedback represents a ability feedback of ability run by Source on target.
type Feedback struct {
	AfbID  ID
	Source ID
	Target ID
}
