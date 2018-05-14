package storage

import (
	"time"

	"github.com/elojah/game_01"
)

// NewDamageDone convert a game.DamageDone into a storage DamageDone.
func NewDamageDone(a game.DamageDone) DamageDone {
	return DamageDone{
		Target: [16]byte(a.Target),
		Amount: a.Amount,
	}
}

// NewDamageReceived convert a game.DamageReceived into a storage DamageReceived.
func NewDamageReceived(a game.DamageReceived) DamageReceived {
	return DamageReceived{
		Source: [16]byte(a.Source),
		Amount: a.Amount,
	}
}

// NewHealDone convert a game.HealDone into a storage HealDone.
func NewHealDone(a game.HealDone) HealDone {
	return HealDone{
		Target: [16]byte(a.Target),
		Amount: a.Amount,
	}
}

// NewHealReceived convert a game.HealReceived into a storage HealReceived.
func NewHealReceived(a game.HealReceived) HealReceived {
	return HealReceived{
		Source: [16]byte(a.Source),
		Amount: a.Amount,
	}
}

// Domain converts a storage DamageDone into a game DamageDone.
func (a DamageDone) Domain() game.DamageDone {
	return game.DamageDone{
		Target: game.ID(a.Target),
		Amount: a.Amount,
	}
}

// Domain converts a storage DamageReceived into a game DamageReceived.
func (a DamageReceived) Domain() game.DamageReceived {
	return game.DamageReceived{
		Source: game.ID(a.Source),
		Amount: a.Amount,
	}
}

// Domain converts a storage HealDone into a game HealDone.
func (a HealDone) Domain() game.HealDone {
	return game.HealDone{
		Target: game.ID(a.Target),
		Amount: a.Amount,
	}
}

// Domain converts a storage HealReceived into a game HealReceived.
func (a HealReceived) Domain() game.HealReceived {
	return game.HealReceived{
		Source: game.ID(a.Source),
		Amount: a.Amount,
	}
}

// NewEvent converts a domain event to a storage event.
func NewEvent(event game.Event) *Event {
	e := &Event{
		ID: [16]byte(event.ID),
		TS: event.TS.UnixNano(),
	}
	switch event.Action.(type) {
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
		ID: game.ID(e.ID),
		TS: time.Unix(0, e.TS),
	}
	switch e.Action.(type) {
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
