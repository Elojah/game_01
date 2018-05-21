package game

import (
	"github.com/nats-io/go-nats"
)

// MsgHandler is a callback function when receiving messge from nats.
type MsgHandler = nats.MsgHandler

// Subscription alias a nats subscription.
type Subscription = nats.Subscription

// SubscriptionMapper creates a new subscription.
type SubscriptionMapper interface {
	SetSubscription(subject string, consumer MsgHandler) (*Subscription, error)
}
