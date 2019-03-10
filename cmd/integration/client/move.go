package client

import (
	"encoding/json"
	"fmt"
	"io"
	"math"

	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/geometry"
	gulid "github.com/elojah/game_01/pkg/ulid"
	"github.com/oklog/ulid"
)

// MoveSameSector move an entity of vec but stay in same sector if entity position + vec is going outside.
func (s *Service) MoveSameSector(tokID gulid.ID, ent entity.E, vec geometry.Vec3) (geometry.Vec3, error) {

	newCoord := geometry.Vec3{
		X: math.Max(math.Min(ent.Position.Coord.X+vec.X, 1024), 0),
		Y: math.Max(math.Min(ent.Position.Coord.Y+vec.Y, 1024), 0),
		Z: math.Max(math.Min(ent.Position.Coord.Z+vec.Z, 1024), 0),
	}

	moveSameSector := event.DTO{
		ID:    gulid.NewTimeID(ulid.Now()),
		Token: tokID,
		Query: event.Query{
			Move: &event.Move{
				Targets: []gulid.ID{ent.ID},
				Position: geometry.Position{
					SectorID: ent.Position.SectorID,
					Coord:    newCoord,
				},
			},
		},
	}
	raw, err := json.Marshal(moveSameSector)
	raw = append(raw, '\n')
	if err != nil {
		return newCoord, errors.Wrap(fmt.Errorf("failed to marshal payload"), "move same sector")
	}

	if _, err := io.WriteString(s.LA.Processes["client"].In, string(raw)); err != nil {
		return newCoord, errors.Wrap(err, "move same sector")
	}

	return newCoord, nil
}

// Move move an entity at position pos. Don't check any distance limit reach.
func (s *Service) Move(tokID gulid.ID, ent entity.E, pos geometry.Position) error {

	moveNextSector := event.DTO{
		ID:    gulid.NewTimeID(ulid.Now()),
		Token: tokID,
		Query: event.Query{
			Move: &event.Move{
				Targets:  []gulid.ID{ent.ID},
				Position: pos,
			},
		},
	}
	raw, err := json.Marshal(moveNextSector)
	raw = append(raw, '\n')
	if err != nil {
		return errors.Wrap(fmt.Errorf("failed to marshal payload"), "move same sector")
	}

	if _, err := io.WriteString(s.LA.Processes["client"].In, string(raw)); err != nil {
		return errors.Wrap(err, "move same sector")
	}

	return nil
}
