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
	targets := make([][16]byte, len(a.Targets))
	for i, target := range a.Targets {
		targets[i] = [16]byte(target)
	}
	return Cast{
		AbilityID: [16]byte(a.AbilityID),
		Source:    [16]byte(a.Source),
		Targets:   targets,
		Position:  Vec3(a.Position),
	}
}

// NewFeedback convert a game.Feedback into a storage Feedback.
func NewFeedback(a game.Feedback) Feedback {
	return Feedback{
		AfbID:  [16]byte(a.AfbID),
		Source: [16]byte(a.Source),
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
	targets := make([]game.ID, len(a.Targets))
	for i, target := range a.Targets {
		targets[i] = game.ID(target)
	}
	return game.Cast{
		AbilityID: game.ID(a.AbilityID),
		Source:    game.ID(a.Source),
		Targets:   targets,
		Position:  game.Vec3(a.Position),
	}
}

// Domain converts a storage Feedback into a game Feedback.
func (a Feedback) Domain() game.Feedback {
	return game.Feedback{
		AfbID:  game.ID(a.AfbID),
		Source: game.ID(a.Source),
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
	case game.Feedback:
		e.Action = NewFeedback(event.Action.(game.Feedback))
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
	case Feedback:
		event.Action = e.Action.(Feedback).Domain()
	}
	return event
}
