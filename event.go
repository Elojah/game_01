package game

import (
	"time"
)

// Event is a game event triggered by an entity or mechanic.
type Event struct {
	ID     ID
	TS     time.Time
	Source ID
	Action Action
}

// QEventMapper must be implemented by a queue.
type QEventMapper interface {
	SendEvent(Event, ID) error
}

// EventMapper wraps action interactions.
type EventMapper interface {
	SetEvent(Event, ID) error
	ListEvent(EventSubset) ([]Event, error)
}

// EventSubset is a subset for actions. Internally usage done with ZRangeWithScores.
type EventSubset struct {
	Key string
	Min int64
}
