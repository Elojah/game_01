package event

import (
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
)

// QStore must be implemented by a queue.
type QStore interface {
	PublishEvent(E, ulid.ID) error
	SubscribeEvent(ulid.ID) *infra.Subscription
}

// Store wraps action interactions.
type Store interface {
	SetEvent(E, ulid.ID) error
	ListEvent(Subset) ([]E, error)
}

// Subset is a subset for actions. Internally usage done with ZRangeWithScores.
type Subset struct {
	Key string
	Min ulid.ID
}
