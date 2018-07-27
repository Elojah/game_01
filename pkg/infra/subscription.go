package infra

import (
	"github.com/go-redis/redis"
)

// QAction is an action required for a queue.
type QAction = uint8

const (
	// Open requires the queue to open.
	Open QAction = 0
	// Close requires the queue to close.
	Close QAction = 1
)

// Subscription alias a redis subscription.
type Subscription = redis.PubSub

// Message alias a redis message.
type Message = redis.Message
