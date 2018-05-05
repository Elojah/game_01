package game

import (
	"time"
)

// Action is a client action.
type Action interface{}

// Listener requires the receiver to create a new listener with subject ID.
type Listener struct {
	ID ID
}

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

// ActionService wraps action interactions.
type ActionService interface {
	CreateAction(Action, ID, time.Time) error
}
