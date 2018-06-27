package event

import (
	"time"

	"github.com/elojah/game_01/pkg/ulid"
)

// QAction is an action required for a queue.
type QAction uint8

const (
	// Open requires the queue to open.
	Open QAction = 0
	// Close requires the queue to close.
	Close QAction = 1
)

// E is a game event triggered by an entity or mechanic.
type E struct {
	ID     ulid.ID
	TS     time.Time
	Source ulid.ID
	Action Action
}

// QMapper must be implemented by a queue.
type QMapper interface {
	SendEvent(E, ulid.ID) error
}

// Mapper wraps action interactions.
type Mapper interface {
	SetEvent(E, ulid.ID) error
	ListEvent(Subset) ([]E, error)
}

// Subset is a subset for actions. Internally usage done with ZRangeWithScores.
type Subset struct {
	Key string
	Min int64
}
