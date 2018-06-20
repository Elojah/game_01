package storage

import (
	"github.com/elojah/game_01"
)

// NewRecurrer converts a domain recurrer to a storage recurrer.
func NewRecurrer(recurrer game.Recurrer) Recurrer {
	return Recurrer{
		ID:       [16]byte(recurrer.ID),
		EntityID: [16]byte(recurrer.EntityID),
		TokenID:  [16]byte(recurrer.TokenID),
		Action:   uint8(recurrer.Action),
	}
}

// Domain converts a storage user into a domain user.
func (r Recurrer) Domain() game.Recurrer {
	return game.Recurrer{
		ID:       game.ID(r.ID),
		EntityID: game.ID(r.EntityID),
		TokenID:  game.ID(r.TokenID),
		Action:   game.RecurrerAction(r.Action),
	}
}
