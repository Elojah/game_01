package storage

import (
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

// NewRecurrer converts a domain recurrer to a storage recurrer.
func NewRecurrer(recurrer event.Recurrer) *Recurrer {
	return &Recurrer{
		ID:       [16]byte(recurrer.ID),
		EntityID: [16]byte(recurrer.EntityID),
		TokenID:  [16]byte(recurrer.TokenID),
		Action:   uint8(recurrer.Action),
	}
}

// Domain converts a storage user into a domain user.
func (r Recurrer) Domain() event.Recurrer {
	return event.Recurrer{
		ID:       ulid.ID(r.ID),
		EntityID: ulid.ID(r.EntityID),
		TokenID:  ulid.ID(r.TokenID),
		Action:   event.QAction(r.Action),
	}
}
