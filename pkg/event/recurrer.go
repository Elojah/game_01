package event

import (
	game "github.com/elojah/game_01"
)

// RecurrerAction is an action required for a recurrer.
type RecurrerAction uint8

const (
	// OpenRec requires the recurrer to open.
	OpenRec RecurrerAction = 0
	// CloseRec requires the recurrer to close.
	CloseRec RecurrerAction = 1
)

// Recurrer requires the receiver to create a new recurrer with subject ID.
type Recurrer struct {
	ID       game.ID
	EntityID game.ID
	TokenID  game.ID
	Action   RecurrerAction
}

// QRecurrerMapper handles send/receive methods for recurrers.
type QRecurrerMapper interface {
	SendRecurrer(Recurrer, game.ID) error
}
