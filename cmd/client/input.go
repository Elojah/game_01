package main

import (
	"encoding/json"

	game "github.com/elojah/game_01"
)

// Input represents a game action sent by player to be send to server.
type Input struct {
	game.Action
}

// UnmarshalJSON unmarshal a game action depending on input type.
func (in *Input) UnmarshalJSON(raw []byte) error {
	var alias struct {
		Type   string          `json:"type"`
		Action json.RawMessage `json:"action"`
	}
	if err := json.Unmarshal(raw, &alias); err != nil {
		return err
	}
	switch alias.Type {
	case "move":
		var action game.Move
		if err := json.Unmarshal(alias.Action, &action); err != nil {
			return err
		}
		in.Action = action
	case "cast":
		var action game.Cast
		if err := json.Unmarshal(alias.Action, &action); err != nil {
			return err
		}
		in.Action = action
	case "feedback":
		var action game.Feedback
		if err := json.Unmarshal(alias.Action, &action); err != nil {
			return err
		}
		in.Action = action
	case "connect_pc":
		var action game.ConnectPC
		if err := json.Unmarshal(alias.Action, &action); err != nil {
			return err
		}
		in.Action = action
	case "set_pc":
		var action game.SetPC
		if err := json.Unmarshal(alias.Action, &action); err != nil {
			return err
		}
		in.Action = action
	default:
		return &json.UnsupportedValueError{Str: alias.Type}
	}
	return nil
}
