package srg

import (
	"github.com/go-redis/redis"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

const (
	triggerKey = "trigger:"
)

// SetTrigger implemented with redis.
func (s *Store) SetTrigger(t event.Trigger) error {
	return s.Set(
		triggerKey+t.EventSourceID.String()+":"+t.EntityID.String(),
		t.EventTargetID.String(),
		0,
	).Err()
}

// GetTrigger list triggers in redis set key from min (included).
func (s *Store) GetTrigger(source gulid.ID, e gulid.ID) (gulid.ID, error) {
	val, err := s.Get(triggerKey + source.String() + ":" + e.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return gulid.ID{}, err
		}
		return gulid.ID{}, gerrors.ErrNotFound
	}

	return gulid.MustParse(val), nil
}

// ListTrigger list triggers in redis set key from min (included).
func (s *Store) ListTrigger(source gulid.ID) ([]event.Trigger, error) {
	vals, err := s.Keys(triggerKey + source.String() + ":*").Result()
	if err != nil {
		return nil, err
	}
	// TODO retrieve entities too
	triggers := make([]gulid.ID, len(vals))
	for i, val := range vals {
		triggers[i] = gulid.MustParse(val)
	}
	return nil, gerrors.ErrNotImplementedYet
	// return triggers, nil
}
