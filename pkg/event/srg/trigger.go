package srg

import (
	"strings"

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

// GetTrigger redis implementation.
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

// ListTrigger redis implementation.
func (s *Store) ListTrigger(source gulid.ID) ([]event.Trigger, error) {
	vals, err := s.Keys(triggerKey + source.String() + ":*").Result()
	if err != nil {
		return nil, err
	}
	// TODO retrieve entities too
	triggers := make([]event.Trigger, len(vals))
	for i, val := range vals {
		etarget, err := s.Get(val).Result()
		if err != nil {
			return nil, err
		}
		triggers[i] = event.Trigger{
			EntityID:      gulid.MustParse(strings.Split(val, ":")[2]),
			EventSourceID: source,
			EventTargetID: gulid.MustParse(etarget),
		}
	}
	return triggers, nil
}

// DelTrigger redis implementation.
func (s *Store) DelTrigger(source gulid.ID, e gulid.ID) error {
	return s.Del(triggerKey + source.String() + ":" + e.String()).Err()
}
