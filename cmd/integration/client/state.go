package client

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/entity"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// GetState retrieve an entity state sent from recurrer to client.
func (s *Service) GetState(id gulid.ID, max int) (entity.E, error) {

	var e entity.E
	if err := s.LA.Retrieve(func(s string) (bool, error) {

		var actual entity.E
		// log is not a JSON
		if err := json.Unmarshal([]byte(s), &actual); err != nil {
			return false, nil
		}
		// log is not an entity
		if actual.ID.IsZero() {
			return false, nil
		}
		if actual.ID.Compare(id) == 0 {
			e = actual
			return true, nil
		}
		return false, nil
	}, max); err != nil {
		return e, errors.Wrap(err, "get state")
	}

	return e, nil
}

// GetStateAt retrieve an entity state sent from recurrer to client when f is positive.
func (s *Service) GetStateAt(id gulid.ID, max int, f func(entity.E) bool) (entity.E, error) {

	var e entity.E
	if err := s.LA.Retrieve(func(s string) (bool, error) {
		var actual entity.E

		// log is not a JSON
		if err := json.Unmarshal([]byte(s), &actual); err != nil {
			return false, nil
		}
		// log is not an entity
		if actual.ID.IsZero() {
			return false, nil
		}
		if actual.ID.Compare(id) == 0 && f(actual) {
			e = actual
			return true, nil
		}
		return false, nil
	}, max); err != nil {
		return e, errors.Wrap(err, "get state at")
	}

	return e, nil
}
