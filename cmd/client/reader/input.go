package reader

import (
	"encoding/json"

	"github.com/oklog/ulid"

	"github.com/elojah/game_01/pkg/dto"
)

// Input represents a game action sent by player to be send to server.
type Input struct {
	dto.Event
}

// UnmarshalJSON unmarshal a game action depending on input type.
func (in *Input) UnmarshalJSON(raw []byte) error {
	var alias struct {
		Type   string
		Action json.RawMessage
	}
	if err := json.Unmarshal(raw, &alias); err != nil {
		return err
	}
	switch alias.Type {
	case "move":
		var actionAlias struct {
			Source   string
			Target   string
			Position dto.Vec3
		}
		if err := json.Unmarshal(alias.Action, &actionAlias); err != nil {
			return err
		}
		source, err := ulid.Parse(actionAlias.Source)
		if err != nil {
			return err
		}
		target, err := ulid.Parse(actionAlias.Target)
		if err != nil {
			return err
		}
		in.Action = dto.Move{
			Source:   [16]byte(source),
			Target:   [16]byte(target),
			Position: actionAlias.Position,
		}
	case "cast":
		var actionAlias struct {
			AbilityID string
			Source    string
			Targets   []string
			Position  dto.Vec3
		}
		if err := json.Unmarshal(alias.Action, &actionAlias); err != nil {
			return err
		}
		source, err := ulid.Parse(actionAlias.Source)
		if err != nil {
			return err
		}
		abilityID, err := ulid.Parse(actionAlias.AbilityID)
		if err != nil {
			return err
		}
		targets := make([][16]byte, len(actionAlias.Targets))
		for i, target := range actionAlias.Targets {
			id, err := ulid.Parse(target)
			if err != nil {
				return err
			}
			targets[i] = [16]byte(id)
		}
		in.Action = dto.Cast{
			AbilityID: abilityID,
			Source:    source,
			Targets:   targets,
			Position:  actionAlias.Position,
		}
	default:
		return &json.UnsupportedValueError{Str: alias.Type}
	}
	return nil
}
