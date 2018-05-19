package game

// Action is a client action.
type Action interface{}

// ActionString acts for action.String() in declarative mode.
func ActionString(a Action) string {
	switch a.(type) {
	case MoveDone:
		return "move_done"
	case MoveReceived:
		return "move_received"
	case AttackDone:
		return "attack_done"
	case AttackReceived:
		return "attack_received"
	case HealDone:
		return "heal_done"
	case HealReceived:
		return "heal_received"
	case SetEntity:
		return "set_entity"
	case SetPC:
		return "set_pc"
	case ConnectPC:
		return "connect_pc"
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

// SetEntity is a player action to create a new target entity.
type SetEntity struct {
	Source   ID
	Type     EntityType
	Position Vec3
}

// MoveDone is an order to move Target.
type MoveDone struct {
	Source   ID
	Target   ID
	Position Vec3
}

// MoveReceived is an ordered to move by Source.
type MoveReceived struct {
	Source   ID
	Target   ID
	Position Vec3
}

// AttackDone is a basic attack done by Source.
type AttackDone struct {
	Source ID
	Target ID
}

// AttackReceived is a basic attack received by Target.
type AttackReceived struct {
	Source ID
	Target ID
}

// HealDone is an amount of heal done by the player to Target.
type HealDone struct {
	Source ID
	Target ID
}

// HealReceived is an amount of heal received by the player and done by Source.
type HealReceived struct {
	Source ID
	Target ID
}
