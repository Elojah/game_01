package game

import (
	"github.com/nats-io/go-nats"

	"time"
)

// Subscription alias a nats subscription.
type Subscription = nats.Subscription

// MsgChan alias a chan of nats Msg.
type MsgChan = chan *nats.Msg

// Event is an entity action.
type Event struct {
	ID     ID
	TS     time.Time
	Action Action
}

// QEventService must be implemented by a queue.
type QEventService interface {
	SendEvent(Event, ID) error
	ReceiveEvent(string, int) (*Subscription, MsgChan, error)
}

// EventService wraps action interactions.
type EventService interface {
	CreateEvent(Event, ID) error
	ListEvent(EventBuilder) ([]Event, error)
}

// EventBuilder is a builder for actions. Internally usage done with ZRangeWithScores.
type EventBuilder struct {
	Key   string
	Start int64
}
