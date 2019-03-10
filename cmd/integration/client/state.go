package client

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/entity"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// GetState retrieve an entity state sent from recurrer to client.
func (s *Service) GetState(id gulid.ID, max int) (entity.E, error) {

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
	}, max); err != nil {
		return actual, errors.Wrap(err, "get state")
	}

	return actual, nil
}

// GetStateAt retrieve an entity state sent from recurrer to client when f is positive.
func (s *Service) GetStateAt(id gulid.ID, max int, f func(entity.E) bool) (entity.E, error) {

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
		if actual.ID.Compare(id) == 0 && f(actual) {
			return true, nil
		}
		return false, nil
	}, max); err != nil {
		return actual, errors.Wrap(err, "get state")
	}

	return actual, nil
}
