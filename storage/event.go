package storage

import (
	"time"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/pkg/event"
)

// NewMove convert a event.Move into a storage Move.
func NewMove(m event.Move) Move {
	return Move{
		Source:   [16]byte(m.Source),
		Target:   [16]byte(m.Target),
		Position: Vec3(m.Position),
	}
}

// NewCast convert a event.Cast into a storage Cast.
func NewCast(c event.Cast) Cast {
	targets := make([][16]byte, len(c.Targets))
	for i, target := range c.Targets {
		targets[i] = [16]byte(target)
	}
	return Cast{
		AbilityID: [16]byte(c.AbilityID),
		Source:    [16]byte(c.Source),
		Targets:   targets,
		Position:  Vec3(c.Position),
	}
}

// NewFeedback convert a event.Feedback into a storage Feedback.
func NewFeedback(fb event.Feedback) Feedback {
	return Feedback{
		AbilityID: [16]byte(fb.AbilityID),
		Source:    [16]byte(fb.Source),
		Target:    [16]byte(fb.Target),
	}
}

// Domain converts a storage Move into a game Move.
func (m Move) Domain() event.Move {
	return event.Move{
		Source:   game.ID(m.Target),
		Target:   game.ID(m.Target),
		Position: game.Vec3(m.Position),
	}
}

// Domain converts a storage Cast into a game Cast.
func (c Cast) Domain() event.Cast {
	targets := make([]game.ID, len(c.Targets))
	for i, target := range c.Targets {
		targets[i] = game.ID(target)
	}
	return event.Cast{
		AbilityID: game.ID(c.AbilityID),
		Source:    game.ID(c.Source),
		Targets:   targets,
		Position:  game.Vec3(c.Position),
	}
}

// Domain converts a storage Feedback into a game Feedback.
func (fb Feedback) Domain() event.Feedback {
	return event.Feedback{
		AbilityID: game.ID(fb.AbilityID),
		Source:    game.ID(fb.Source),
		Target:    game.ID(fb.Target),
	}
}

// NewEvent converts a domain event to a storage event.
func NewEvent(ev event.E) *Event {
	e := &Event{
		ID:     [16]byte(ev.ID),
		Source: [16]byte(ev.Source),
		TS:     ev.TS.UnixNano(),
	}
	switch ev.Action.(type) {
	case event.Move:
		e.Action = NewMove(ev.Action.(event.Move))
	case event.Cast:
		e.Action = NewCast(ev.Action.(event.Cast))
	case event.Feedback:
		e.Action = NewFeedback(ev.Action.(event.Feedback))
	}
	return e
}

// Domain converts a storage user into a domain user.
func (e Event) Domain() event.E {
	ev := event.E{
		ID:     game.ID(e.ID),
		Source: game.ID(e.Source),
		TS:     time.Unix(0, e.TS),
	}
	switch e.Action.(type) {
	case Move:
		ev.Action = e.Action.(Move).Domain()
	case Cast:
		ev.Action = e.Action.(Cast).Domain()
	case Feedback:
		ev.Action = e.Action.(Feedback).Domain()
	}
	return ev
}
