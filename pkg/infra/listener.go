package infra

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// QListenerStore handles send/receive methods for listeners.
type QListenerStore interface {
	PublishListener(Listener, ulid.ID) error
	SubscribeListener(ulid.ID) *Subscription
}

// ListenerStore handle listener data interactions.
type ListenerStore interface {
	SetListener(Listener) error
	GetListener(ListenerSubset) (Listener, error)
	DelListener(ListenerSubset) error
}

// ListenerSubset retrieves listener per ID.
type ListenerSubset struct {
	ID ulid.ID
}

// ListenerService represents listener usecases.
type ListenerService interface {
	New(ulid.ID) (Listener, error)
	Remove(ulid.ID) error
}
