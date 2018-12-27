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
	ListEvent(ulid.ID, ulid.ID) ([]E, error)
	DelEvent(ulid.ID, ulid.ID) error
}
