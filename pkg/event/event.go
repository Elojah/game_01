package event

import (
	"time"

	game "github.com/elojah/game_01"
)

// E is a game event triggered by an entity or mechanic.
type E struct {
	ID     game.ID
	TS     time.Time
	Source game.ID
	Action Action
}

// QMapper must be implemented by a queue.
type QMapper interface {
	SendEvent(E, game.ID) error
}

// Mapper wraps action interactions.
type Mapper interface {
	SetEvent(E, game.ID) error
	ListEvent(Subset) ([]E, error)
}

// Subset is a subset for actions. Internally usage done with ZRangeWithScores.
type Subset struct {
	Key string
	Min int64
}
