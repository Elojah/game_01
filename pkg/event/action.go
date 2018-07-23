package event

import (
	"github.com/elojah/game_01/pkg/entity"
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
	Position entity.Position
}

// Cast represents a ability launch.
type Cast struct {
	AbilityID ulid.ID
	Source    ulid.ID
	Targets   []ulid.ID
	Position  entity.Position
}

// Feedback represents a ability feedback of ability run by Source on target.
type Feedback struct {
	ID     ulid.ID
	Source ulid.ID
	Target ulid.ID
}
