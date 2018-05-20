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

// NewSkill convert a game.Skill into a storage Skill.
func NewSkill(a game.Skill) Skill {
	return Skill{
		ID:       [16]byte(a.ID),
		Source:   [16]byte(a.Source),
		Target:   [16]byte(a.Target),
		Position: Vec3(a.Position),
	}
}

// NewSetPC convert a game.SetPC into a storage SetPC.
func NewSetPC(a game.SetPC) SetPC {
	return SetPC{
		Type: uint8(a.Type),
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

// Domain converts a storage Skill into a game Skill.
func (a Skill) Domain() game.Skill {
	return game.Skill{
		ID:       game.ID(a.ID),
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
	case game.Skill:
		e.Action = NewSkill(event.Action.(game.Skill))
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
	case Skill:
		event.Action = e.Action.(Skill).Domain()
	case SetPC:
		event.Action = e.Action.(SetPC).Domain()
	case ConnectPC:
		event.Action = e.Action.(ConnectPC).Domain()
	}
	return event
}
