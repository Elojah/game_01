package event

import (
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/ulid"
)

// Action is a client action.
type Action interface{}

// String serialize an action.
func String(a Action) string {
	switch a.(type) {
	case Move:
		return "move"
	case Cast:
		return "cast"
	case Casted:
		return "casted"
	case Feedback:
		return "feedback"
	default:
		return "unknown"
	}
}

// Move represents a unit move.
type Move struct {
	Source   ulid.ID
	Target   ulid.ID
	Position geometry.Position
}

// Cast represents a ability start cast.
type Cast struct {
	AbilityID ulid.ID
	Source    ulid.ID
	Targets   []ulid.ID
	Position  geometry.Position
}

// Feedback represents a ability feedback of ability run by Source on target.
type Feedback struct {
	ID     ulid.ID
	Source ulid.ID
	Target ulid.ID
}

// Casted represents a end of cast ability.
type Casted Cast
