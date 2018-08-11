package infra

import "github.com/elojah/game_01/pkg/ulid"

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
