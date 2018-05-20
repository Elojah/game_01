package game

// Action is a client action.
type Action interface{}

// ActionString acts for action.String() in declarative mode.
func ActionString(a Action) string {
	switch a.(type) {
	case Move:
		return "move"
	case Skill:
		return "skill"
	case ConnectPC:
		return "connect_pc"
	case SetPC:
		return "set_pc"
	default:
		return "unknown"
	}
}

// SetPC is a token action to create a new PC entity.
type SetPC struct {
	Type EntityType
}

// ConnectPC is a token action to connect to a previously created PC.
type ConnectPC struct {
	Target ID
}
