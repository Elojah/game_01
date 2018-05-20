package storage

import (
	"time"

	"github.com/elojah/game_01"
)

// NewMove convert a game.Move into a storage Move.
func NewMove(a game.Move) Move {
	return Move{
		Source:   [16]byte(a.Source),
		Target:   [16]byte(a.Target),
		Position: Vec3(a.Position),
	}
}

// NewCast convert a game.Cast into a storage Cast.
func NewCast(a game.Cast) Cast {
	return Cast{
		SkillID:  [16]byte(a.SkillID),
		Source:   [16]byte(a.Source),
		Target:   [16]byte(a.Target),
		Position: Vec3(a.Position),
	}
}

// NewSetPC convert a game.SetPC into a storage SetPC.
func NewSetPC(a game.SetPC) SetPC {
	return SetPC{
		Type: [16]byte(a.Type),
	}
}

// NewConnectPC convert a game.ConnectPC into a storage ConnectPC.
func NewConnectPC(a game.ConnectPC) ConnectPC {
	return ConnectPC{
		Target: [16]byte(a.Target),
	}
}

// Domain converts a storage Move into a game Move.
func (a Move) Domain() game.Move {
	return game.Move{
		Source:   game.ID(a.Target),
		Target:   game.ID(a.Target),
		Position: game.Vec3(a.Position),
	}
}

// Domain converts a storage Cast into a game Cast.
func (a Cast) Domain() game.Cast {
	return game.Cast{
		SkillID:  game.ID(a.SkillID),
		Source:   game.ID(a.Source),
		Target:   game.ID(a.Target),
		Position: game.Vec3(a.Position),
	}
}

// Domain converts a storage SetPC into a game SetPC.
func (a SetPC) Domain() game.SetPC {
	return game.SetPC{
		Type: game.EntityType(a.Type),
	}
}

// Domain converts a storage ConnectPC into a game ConnectPC.
func (a ConnectPC) Domain() game.ConnectPC {
	return game.ConnectPC{
		Target: game.ID(a.Target),
	}
}

// NewEvent converts a domain event to a storage event.
func NewEvent(event game.Event) *Event {
	e := &Event{
		ID:     [16]byte(event.ID),
		Source: [16]byte(event.Source),
		TS:     event.TS.UnixNano(),
	}
	switch event.Action.(type) {
	case game.Move:
		e.Action = NewMove(event.Action.(game.Move))
	case game.Cast:
		e.Action = NewCast(event.Action.(game.Cast))
	case game.SetPC:
		e.Action = NewSetPC(event.Action.(game.SetPC))
	case game.ConnectPC:
		e.Action = NewConnectPC(event.Action.(game.ConnectPC))
	}
	return e
}

// Domain converts a storage user into a domain user.
func (e Event) Domain() game.Event {
	event := game.Event{
		ID:     game.ID(e.ID),
		Source: game.ID(e.Source),
		TS:     time.Unix(0, e.TS),
	}
	switch e.Action.(type) {
	case Move:
		event.Action = e.Action.(Move).Domain()
	case Cast:
		event.Action = e.Action.(Cast).Domain()
	case SetPC:
		event.Action = e.Action.(SetPC).Domain()
	case ConnectPC:
		event.Action = e.Action.(ConnectPC).Domain()
	}
	return event
}
