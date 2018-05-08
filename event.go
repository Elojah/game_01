package game

import (
	"time"
)

// Event is an entity action.
type Event struct {
	ID     ID
	TS     time.Time
	Action Action
}

// QEventService must be implemented by a queue.
type QEventService interface {
	SendEvent(Event, ID) error
}

// EventService wraps action interactions.
type EventService interface {
	CreateEvent(Event, ID) error
	ListEvent(EventBuilder) ([]Event, error)
}

// EventBuilder is a builder for actions. Internally usage done with ZRangeWithScores.
type EventBuilder struct {
	Key string
	Min int
}
