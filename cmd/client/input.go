package main

import (
	"github.com/elojah/game_01/pkg/event"
)

// Input represents a game action sent by player to be send to server.
type Input struct {
	event.DTO
}
