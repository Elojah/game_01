package event

import (
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/ulid"
)

// Action is a client action.
type Action interface{}

// ActionString acts for action.String() in declarative mode.
func ActionString(a Action) string {
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
	Position geometry.Vec3
}

// Cast represents a ability launch.
type Cast struct {
	AbilityID ulid.ID
	Source    ulid.ID
	Targets   []ulid.ID
	Position  geometry.Vec3
}

// Feedback represents a ability feedback of ability run by Source on target.
type Feedback struct {
	AbilityID ulid.ID
	Source    ulid.ID
	Target    ulid.ID
}
