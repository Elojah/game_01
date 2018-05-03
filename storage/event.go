package storage

import (
	"time"

	"github.com/elojah/game_01"
)

// NewEvent converts a domain event to a storage event.
func NewEvent(event game.Event) *Event {
	return &Event{
		ID:     [16]byte(event.ID),
		TS:     event.TS.UnixNano(),
		Action: event.Action,
	}
}

// Domain converts a storage user into a domain user.
func (e *Event) Domain() game.Event {
	return game.Event{
		ID:     game.ID(e.ID),
		TS:     time.Unix(0, e.TS),
		Action: e.Action,
	}
}
