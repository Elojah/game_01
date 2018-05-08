package game

import (
	"github.com/nats-io/go-nats"
)

// MsgChan alias a chan of nats Msg.
type MsgChan = chan *nats.Msg

// Subscription alias a nats subscription.
type Subscription struct {
	*nats.Subscription
	Ch MsgChan
}

// Close unsubscribe and close receiving channel.
func (s *Subscription) Close() {
	s.Unsubscribe()
	close(s.Ch)
}

// SubscriptionService creates a new subscription.
// TODO use this instead of event/listener services for receive.
type SubscriptionService interface {
	CreateSubscription(subject string, bufsize int) (Subscription, error)
}
