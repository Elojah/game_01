package storage

import (
	"github.com/elojah/game_01"
)

// NewEvent converts a domain event to a storage event.
func NewEvent(event game.Event) *Event {
	return &Event{
		ID:     [16]byte(event.ID),
		Action: event.Action,
	}
}
