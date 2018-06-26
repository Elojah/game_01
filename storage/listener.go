package storage

import (
	"github.com/elojah/game_01"
	"github.com/elojah/game_01/pkg/event"
)

// NewListener converts a domain listener to a storage listener.
func NewListener(listener event.Listener) Listener {
	return Listener{
		ID: [16]byte(listener.ID),
	}
}

// Domain converts a storage user into a domain user.
func (l Listener) Domain() event.Listener {
	return event.Listener{
		ID: game.ID(l.ID),
	}
}
