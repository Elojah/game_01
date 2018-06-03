package storage

import (
	"github.com/elojah/game_01"
)

// NewRecurrer converts a domain recurrer to a storage recurrer.
func NewRecurrer(recurrer game.Recurrer) Recurrer {
	return Recurrer{
		ID: [16]byte(recurrer.ID),
	}
}

// Domain converts a storage user into a domain user.
func (l Recurrer) Domain() game.Recurrer {
	return game.Recurrer{
		ID: game.ID(l.ID),
	}
}
