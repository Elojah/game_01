package event

import "github.com/elojah/game_01/pkg/ulid"

// Listener requires the receiver to create a new listener with subject ID.
type Listener struct {
	ID     ulid.ID
	Action QAction
	Pool   ulid.ID
}

// QListenerMapper handles send/receive methods for listeners.
type QListenerMapper interface {
	SendListener(Listener) error
}
