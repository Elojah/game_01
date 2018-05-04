package storage

import (
	"time"

	"github.com/elojah/game_01"
)

// NewEvent converts a domain event to a storage event.
func NewEvent(event game.Event) *Event {
	e := &Event{
		ID: [16]byte(event.ID),
		TS: event.TS.UnixNano(),
	}
	switch event.Action.(type) {
	case game.Listener:
		e.Action = NewListener(event.Action.(game.Listener))
	}
	return e
}

// Domain converts a storage user into a domain user.
func (e Event) Domain() game.Event {
	event := game.Event{
		ID: game.ID(e.ID),
		TS: time.Unix(0, e.TS),
	}
	switch e.Action.(type) {
	case Listener:
		event.Action = e.Action.(Listener).Domain()
	}
	return event
}

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
