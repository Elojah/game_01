package main

import (
	game "github.com/elojah/game_01"
)

// Input represents a game action sent by player to be send to server.
type Input struct {
	Type   string      `json:"type"`
	Action game.Action `json:"action"`
}

// UnmarshalJSON unmarshal a game action depending on input type.
func (in *Input) UnmarshalJSON(raw []byte) error {
	return nil
}
