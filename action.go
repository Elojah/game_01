package game

// Action is a client action.
type Action interface{}

// Damage received.
type Damage struct {
	Source ID
	Amount int64
}

// DamageInflict inflicted.
type DamageInflict struct {
	Target ID
	Amount int64
}

// Heal received.
type Heal struct {
	Source ID
	Amount int64
}

// HealInflict inflicted.
type HealInflict struct {
	Target ID
	Amount int64
}
