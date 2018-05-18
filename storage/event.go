package storage

import (
	"time"

	"github.com/elojah/game_01"
)

// NewMoveDone convert a game.MoveDone into a storage MoveDone.
func NewMoveDone(a game.MoveDone) MoveDone {
	return MoveDone{
		Source: [16]byte(a.Source),
		Target: [16]byte(a.Target),
		X:      a.Position.X,
		Y:      a.Position.Y,
		Z:      a.Position.Z,
	}
}

// NewMoveReceived convert a game.MoveReceived into a storage MoveReceived.
func NewMoveReceived(a game.MoveReceived) MoveReceived {
	return MoveReceived{
		Source: [16]byte(a.Source),
		Target: [16]byte(a.Target),
		X:      a.Position.X,
		Y:      a.Position.Y,
		Z:      a.Position.Z,
	}
}

// NewAttackDone convert a game.AttackDone into a storage AttackDone.
func NewAttackDone(a game.AttackDone) AttackDone {
	return AttackDone{
		Source: [16]byte(a.Source),
		Target: [16]byte(a.Target),
	}
}

// NewAttackReceived convert a game.AttackReceived into a storage AttackReceived.
func NewAttackReceived(a game.AttackReceived) AttackReceived {
	return AttackReceived{
		Source: [16]byte(a.Source),
		Target: [16]byte(a.Target),
	}
}

// NewHealDone convert a game.HealDone into a storage HealDone.
func NewHealDone(a game.HealDone) HealDone {
	return HealDone{
		Source: [16]byte(a.Source),
		Target: [16]byte(a.Target),
	}
}

// NewHealReceived convert a game.HealReceived into a storage HealReceived.
func NewHealReceived(a game.HealReceived) HealReceived {
	return HealReceived{
		Source: [16]byte(a.Source),
		Target: [16]byte(a.Target),
	}
}

// NewSetEntity convert a game.SetEntity into a storage SetEntity.
func NewSetEntity(a game.SetEntity) SetEntity {
	return SetEntity{
		Source: [16]byte(a.Source),
		Type:   uint8(a.Type),
		X:      a.Position.X,
		Y:      a.Position.Y,
		Z:      a.Position.Z,
	}
}

// NewSetPC convert a game.SetPC into a storage SetPC.
func NewSetPC(a game.SetPC) SetPC {
	return SetPC{
		Type: uint8(a.Type),
	}
}

// Domain converts a storage MoveDone into a game MoveDone.
func (a MoveDone) Domain() game.MoveDone {
	return game.MoveDone{
		Source: game.ID(a.Target),
		Target: game.ID(a.Target),
		Position: game.Vec3{
			X: a.X,
			Y: a.Y,
			Z: a.Z,
		},
	}
}

// Domain converts a storage MoveReceived into a game MoveReceived.
func (a MoveReceived) Domain() game.MoveReceived {
	return game.MoveReceived{
		Source: game.ID(a.Target),
		Target: game.ID(a.Target),
		Position: game.Vec3{
			X: a.X,
			Y: a.Y,
			Z: a.Z,
		},
	}
}

// Domain converts a storage AttackDone into a game AttackDone.
func (a AttackDone) Domain() game.AttackDone {
	return game.AttackDone{
		Source: game.ID(a.Source),
		Target: game.ID(a.Target),
	}
}

// Domain converts a storage AttackReceived into a game AttackReceived.
func (a AttackReceived) Domain() game.AttackReceived {
	return game.AttackReceived{
		Source: game.ID(a.Source),
		Target: game.ID(a.Target),
	}
}

// Domain converts a storage HealDone into a game HealDone.
func (a HealDone) Domain() game.HealDone {
	return game.HealDone{
		Source: game.ID(a.Source),
		Target: game.ID(a.Target),
	}
}

// Domain converts a storage HealReceived into a game HealReceived.
func (a HealReceived) Domain() game.HealReceived {
	return game.HealReceived{
		Source: game.ID(a.Source),
		Target: game.ID(a.Target),
	}
}

// Domain converts a storage SetEntity into a game SetEntity.
func (a SetEntity) Domain() game.SetEntity {
	return game.SetEntity{
		Source: game.ID(a.Source),
		Type:   game.EntityType(a.Type),
		Position: game.Vec3{
			X: a.X,
			Y: a.Y,
			Z: a.Z,
		},
	}
}

// Domain converts a storage SetPC into a game SetPC.
func (a SetPC) Domain() game.SetPC {
	return game.SetPC{
		Type: game.EntityType(a.Type),
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
	case game.MoveDone:
		e.Action = NewMoveDone(event.Action.(game.MoveDone))
	case game.MoveReceived:
		e.Action = NewMoveReceived(event.Action.(game.MoveReceived))
	case game.AttackDone:
		e.Action = NewAttackDone(event.Action.(game.AttackDone))
	case game.AttackReceived:
		e.Action = NewAttackReceived(event.Action.(game.AttackReceived))
	case game.HealDone:
		e.Action = NewHealDone(event.Action.(game.HealDone))
	case game.HealReceived:
		e.Action = NewHealReceived(event.Action.(game.HealReceived))
	case game.SetEntity:
		e.Action = NewSetEntity(event.Action.(game.SetEntity))
	case game.SetPC:
		e.Action = NewSetPC(event.Action.(game.SetPC))
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
	case MoveDone:
		event.Action = e.Action.(MoveDone).Domain()
	case MoveReceived:
		event.Action = e.Action.(MoveReceived).Domain()
	case AttackDone:
		event.Action = e.Action.(AttackDone).Domain()
	case AttackReceived:
		event.Action = e.Action.(AttackReceived).Domain()
	case HealDone:
		event.Action = e.Action.(HealDone).Domain()
	case HealReceived:
		event.Action = e.Action.(HealReceived).Domain()
	case SetEntity:
		event.Action = e.Action.(SetEntity).Domain()
	case SetPC:
		event.Action = e.Action.(SetPC).Domain()
	}
	return event
}
