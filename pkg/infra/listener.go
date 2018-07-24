package infra

import "github.com/elojah/game_01/pkg/ulid"

// Listener requires the receiver to create a new listener with subject ID.
type Listener struct {
	ID     ulid.ID // ID aliases a token ID OR an entity ID
	Action QAction
	Pool   ulid.ID
}

// QListenerMapper handles send/receive methods for listeners.
type QListenerMapper interface {
	PublishListener(Listener, ulid.ID) error
	SubscribeListener(ulid.ID) *Subscription
}

// ListenerMapper handle listener data interactions.
type ListenerMapper interface {
	SetListener(Listener) error
	GetListener(ListenerSubset) (Listener, error)
	DelListener(ListenerSubset) error
}

// ListenerSubset retrieves listener per ID.
type ListenerSubset struct {
	ID ulid.ID
}
