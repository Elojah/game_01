package tile38

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/oklog/ulid"

	"github.com/elojah/game_01"
)

type t38response struct {
	Ok      bool
	Fields  []string
	Objects []struct {
		ID     string
		Object struct {
			Type_       string `json:"type"`
			Coordinates []float64
		}
		Fields []string
	}
	Count   int
	Cursor  int
	Elapsed string
}

// SetEntityPosition is the tile38 implementation of EntityPosition service,
func (s *Service) SetEntityPosition(entity game.Entity, ts int64) error {
	cmd := redis.NewStringCmd(
		"SET",
		"entity",
		entity.ID.String(),
		"FIELD",
		"ts",
		ts,
		"POINT",
		entity.Position.X,
		entity.Position.Y,
	)
	if err := s.Process(cmd); err != nil {
		return err
	}
	_, err := cmd.Result()
	return err
}

// ListEntityPosition is the tile38 implementation of EntityPosition service,
func (s *Service) ListEntityPosition(subset game.EntityPositionSubset) ([]game.Entity, error) {
	cmd := redis.NewStringCmd(
		"NEARBY",
		"entity",
		"LIMIT",
		subset.Limit,
		"POINT",
		subset.Position.X,
		subset.Position.Y,
		subset.Radius,
	)
	if err := s.Process(cmd); err != nil {
		return nil, err
	}
	fmt.Println(cmd)
	val, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	var resp t38response
	if err := json.Unmarshal([]byte(val), &resp); err != nil {
		return nil, err
	}
	entities := make([]game.Entity, len(resp.Objects))
	for i, object := range resp.Objects {
		id, err := ulid.Parse(object.ID)
		if err != nil {
			return nil, err
		}
		entities[i] = game.Entity{
			ID: id,
			Position: game.Vec3{
				X: object.Object.Coordinates[0],
				Y: object.Object.Coordinates[1],
				Z: 0,
			},
		}
	}
	return entities, nil
}
