package event

import (
	"github.com/go-redis/redis"
)

// Subscription alias a redis subscription.
type Subscription = redis.PubSub

// Message alias a redis message.
type Message = redis.Message
