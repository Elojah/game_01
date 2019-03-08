package client

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/geometry"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// GetState retrieve an entity state sent from recurrer to client when entity reach position.
func (s *Service) GetState(id gulid.ID) (entity.E, error) {

	var actual entity.E
	if err := s.LA.Retrieve(func(s string) (bool, error) {
		// log is not a JSON
		if err := json.Unmarshal([]byte(s), &actual); err != nil {
			return false, nil
		}
		// log is not an entity
		if actual.ID.IsZero() {
			return false, nil
		}
		if actual.ID.Compare(id) == 0 {
			return true, nil
		}
		return false, nil
	}, 50); err != nil {
		return actual, errors.Wrap(err, "get state")
	}

	return actual, nil
}

// GetState retrieve an entity state sent from recurrer to client when entity reach position.
func (s *Service) GetStateAtPosition(id gulid.ID, pos geometry.Position) (entity.E, error) {

	var actual entity.E
	if err := s.LA.Retrieve(func(s string) (bool, error) {
		// log is not a JSON
		if err := json.Unmarshal([]byte(s), &actual); err != nil {
			return false, nil
		}
		// log is not an entity
		if actual.ID.IsZero() {
			return false, nil
		}
		if actual.ID.Compare(id) == 0 &&
			actual.Position.SectorID.Compare(pos.SectorID) == 0 &&
			actual.Position.Coord == pos.Coord {
			return true, nil
		}
		return false, nil
	}, 50); err != nil {
		return actual, errors.Wrap(err, "get state")
	}

	return actual, nil
}
