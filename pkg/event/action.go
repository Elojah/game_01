package event

import (
	game "github.com/elojah/game_01"
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
	Source   game.ID
	Target   game.ID
	Position game.Vec3
}

// Cast represents a ability launch.
type Cast struct {
	AbilityID game.ID
	Source    game.ID
	Targets   []game.ID
	Position  game.Vec3
}

// Feedback represents a ability feedback of ability run by Source on target.
type Feedback struct {
	AbilityID game.ID
	Source    game.ID
	Target    game.ID
}
