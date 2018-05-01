package game

import (
	"time"
)

// Damage received.
type Damage struct {
	Source ID
	Amount int64
}

// Heal received.
type Heal struct {
	Source ID
	Amount int64
}

// Event is an entity action.
type Event struct {
	ID     ID
	TS     time.Time
	Action interface{}
}

// EventService must be implemented by a queue.
type EventService interface {
	SendEvent(event Event, target ID) error
}
