package game

import (
	"time"
)

// Event is an entity action.
type Event struct {
	ID     ID
	TS     time.Time
	Source ID
	Action Action
}

// QEventService must be implemented by a queue.
type QEventService interface {
	SendEvent(Event, ID) error
}

// EventService wraps action interactions.
type EventService interface {
	SetEvent(Event, ID) error
	ListEvent(EventSubset) ([]Event, error)
}

// EventSubset is a subset for actions. Internally usage done with ZRangeWithScores.
type EventSubset struct {
	Key string
	Min int64
}
