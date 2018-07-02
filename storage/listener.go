package storage

import (
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

// NewListener converts a domain listener to a storage listener.
func NewListener(listener event.Listener) *Listener {
	return &Listener{
		ID:     [16]byte(listener.ID),
		Action: uint8(listener.Action),
	}
}

// Domain converts a storage user into a domain user.
func (l Listener) Domain() event.Listener {
	return event.Listener{
		ID:     ulid.ID(l.ID),
		Action: event.QAction(l.Action),
	}
}
