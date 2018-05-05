package game

import (
	"github.com/nats-io/go-nats"

	"time"
)

// Subscription alias a nats subscription.
type Subscription nats.Subscription

// MsgChan alias a chan of nats Msg.
type MsgChan chan *nats.Msg

// Event is an entity action.
type Event struct {
	ID     ID
	TS     time.Time
	Action Action
}

// EventService must be implemented by a queue.
type EventService interface {
	SendEvent(event Event, target ID) error
	ReceiveEvent(subject string, bufsize int) (*Subscription, MsgChan, error)
}
