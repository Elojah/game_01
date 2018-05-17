package game

import (
	"github.com/nats-io/go-nats"
)

// MsgHandler is a callback function when receiving messge from nats.
type MsgHandler = nats.MsgHandler

// Subscription alias a nats subscription.
type Subscription = nats.Subscription

// SubscriptionService creates a new subscription.
type SubscriptionService interface {
	SetSubscription(subject string, consumer MsgHandler) (*Subscription, error)
}
