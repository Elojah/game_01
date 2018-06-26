package event

import (
	game "github.com/elojah/game_01"
)

// Listener requires the receiver to create a new listener with subject ID.
type Listener struct {
	ID game.ID
}

// QListenerMapper handles send/receive methods for listeners.
type QListenerMapper interface {
	SendListener(Listener, game.ID) error
}
