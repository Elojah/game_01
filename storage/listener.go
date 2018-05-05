package storage

import (
	"github.com/elojah/game_01"
)

// NewListener converts a domain listener to a storage listener.
func NewListener(listener game.Listener) Listener {
	return Listener{
		ID: [16]byte(listener.ID),
	}
}

// Domain converts a storage user into a domain user.
func (l Listener) Domain() game.Listener {
	return game.Listener{
		ID: game.ID(l.ID),
	}
}
