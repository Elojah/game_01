package game

// Action is a client action.
type Action interface{}

// CreatePC is a token action to create a new PC entity.
type CreatePC struct {
	Type EntityType
}

// CreateEntity is a player action to create a new target entity.
type CreateEntity struct {
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
