package client

import (
	"encoding/json"
	"fmt"
	"io"
	"math"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/geometry"
	gulid "github.com/elojah/game_01/pkg/ulid"
	"github.com/oklog/ulid"
)

// MoveSameSector move an entity of vec but stay in same sector if entity position + vec is going outside.
func (s *Service) MoveSameSector(tokID gulid.ID, ent entity.E, vec geometry.Vec3) error {

	newCoord := geometry.Vec3{
		X: math.Min(ent.Position.Coord.X+33, 1024),
		Y: math.Min(ent.Position.Coord.Y+33, 1024),
		Z: math.Min(ent.Position.Coord.Z+33, 1024),
	}

	now := ulid.Now()
	moveSameSector := event.DTO{
		ID:    gulid.NewTimeID(now),
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
		return fmt.Errorf("failed to marshal payload")
	}

	if _, err := io.WriteString(s.in, string(raw)); err != nil {
		return err
	}

	return nil
}
