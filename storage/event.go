package storage

import (
	"time"

	"github.com/elojah/game_01"
)

// NewMove convert a game.Move into a storage Move.
func NewMove(a game.Move) Move {
	return Move{
		X:      a.Position.X,
		Y:      a.Position.Y,
		Z:      a.Position.Z,
		Target: [16]byte(a.Target),
	}
}

// NewDamageDone convert a game.DamageDone into a storage DamageDone.
func NewDamageDone(a game.DamageDone) DamageDone {
	return DamageDone{
		Source: [16]byte(a.Source),
		Target: [16]byte(a.Target),
		Amount: a.Amount,
	}
}

// NewDamageReceived convert a game.DamageReceived into a storage DamageReceived.
func NewDamageReceived(a game.DamageReceived) DamageReceived {
	return DamageReceived{
		Source: [16]byte(a.Source),
		Target: [16]byte(a.Target),
		Amount: a.Amount,
	}
}

// NewHealDone convert a game.HealDone into a storage HealDone.
func NewHealDone(a game.HealDone) HealDone {
	return HealDone{
		Source: [16]byte(a.Source),
		Target: [16]byte(a.Target),
		Amount: a.Amount,
	}
}

// NewHealReceived convert a game.HealReceived into a storage HealReceived.
func NewHealReceived(a game.HealReceived) HealReceived {
	return HealReceived{
		Source: [16]byte(a.Source),
		Target: [16]byte(a.Target),
		Amount: a.Amount,
	}
}

// Domain converts a storage Move into a game Move.
func (a Move) Domain() game.Move {
	return game.Move{
		Position: game.Vec3{
			X: a.X,
			Y: a.Y,
			Z: a.Z,
		},
		Target: [16]byte(a.Target),
	}
}

// Domain converts a storage DamageDone into a game DamageDone.
func (a DamageDone) Domain() game.DamageDone {
	return game.DamageDone{
		Source: game.ID(a.Source),
		Target: game.ID(a.Target),
		Amount: a.Amount,
	}
}

// Domain converts a storage DamageReceived into a game DamageReceived.
func (a DamageReceived) Domain() game.DamageReceived {
	return game.DamageReceived{
		Source: game.ID(a.Source),
		Target: game.ID(a.Target),
		Amount: a.Amount,
	}
}

// Domain converts a storage HealDone into a game HealDone.
func (a HealDone) Domain() game.HealDone {
	return game.HealDone{
		Source: game.ID(a.Source),
		Target: game.ID(a.Target),
		Amount: a.Amount,
	}
}

// Domain converts a storage HealReceived into a game HealReceived.
func (a HealReceived) Domain() game.HealReceived {
	return game.HealReceived{
		Source: game.ID(a.Source),
		Target: game.ID(a.Target),
		Amount: a.Amount,
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
	case game.DamageDone:
		e.Action = NewDamageDone(event.Action.(game.DamageDone))
	case game.DamageReceived:
		e.Action = NewDamageReceived(event.Action.(game.DamageReceived))
	case game.HealDone:
		e.Action = NewHealDone(event.Action.(game.HealDone))
	case game.HealReceived:
		e.Action = NewHealReceived(event.Action.(game.HealReceived))
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
	case DamageDone:
		event.Action = e.Action.(DamageDone).Domain()
	case DamageReceived:
		event.Action = e.Action.(DamageReceived).Domain()
	case HealDone:
		event.Action = e.Action.(HealDone).Domain()
	case HealReceived:
		event.Action = e.Action.(HealReceived).Domain()
	}
	return event
}
