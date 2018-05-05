package game

import (
	"github.com/nats-io/go-nats"

	"time"
)

// Event is an entity action.
type Event struct {
	ID     ID
	TS     time.Time
	Action Action
}

// EventService must be implemented by a queue.
type EventService interface {
	SendEvent(event Event, target ID) error
	ReceiveEvent(subject string, bufsize int) (*nats.Subscription, chan *nats.Msg, error)
}
