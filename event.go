package game

import (
	"github.com/nats-io/go-nats"

	"time"
)

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

// Event is an entity action.
type Event struct {
	ID     ID
	TS     time.Time
	Action interface{}
}

// EventService must be implemented by a queue.
type EventService interface {
	SendEvent(event Event, target ID) error
	ReceiveEvent(subject string, bufsize int) (*nats.Subscription, chan *nats.Msg, error)
}
