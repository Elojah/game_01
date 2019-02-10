package event

import (
	"github.com/elojah/game_01/pkg/infra"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// QStore must be implemented by a queue.
type QStore interface {
	PublishEvent(E, gulid.ID) error
	SubscribeEvent(gulid.ID) *infra.Subscription
}

// Store wraps action interactions.
type Store interface {
	SetEvent(E, gulid.ID) error
	GetEvent(gulid.ID, gulid.ID) (E, error)
	ListEvent(gulid.ID, gulid.ID) ([]E, error)
	DelEvent(gulid.ID, gulid.ID) error
}
