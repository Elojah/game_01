package game

// Action is a client action.
type Action interface{}

// Move is a new position of player.
type Move Vec3

// DamageDone is an amount of damage done by the player to Target.
type DamageDone struct {
	Target ID
	Amount int64
}

// DamageReceived is an amount of damage received by the player and done by Source.
type DamageReceived struct {
	Source ID
	Amount int64
}

// HealDone is an amount of heal done by the player to Target.
type HealDone struct {
	Target ID
	Amount int64
}

// HealReceived is an amount of heal received by the player and done by Source.
type HealReceived struct {
	Source ID
	Amount int64
}
