package game

// Action is a client action.
type Action interface{}

// CreateEntity is an entity creation.
type CreateEntity struct {
	Source   ID
	Type     EntityType
	Position Vec3
}

// Move is a new position of player.
type Move struct {
	Position Vec3
	Target   ID
}

// DamageDone is an amount of damage done by the player to Target.
type DamageDone struct {
	Source ID
	Target ID
	Amount int64
}

// DamageReceived is an amount of damage received by the player and done by Source.
type DamageReceived struct {
	Source ID
	Target ID
	Amount int64
}

// HealDone is an amount of heal done by the player to Target.
type HealDone struct {
	Source ID
	Target ID
	Amount int64
}

// HealReceived is an amount of heal received by the player and done by Source.
type HealReceived struct {
	Source ID
	Target ID
	Amount int64
}
