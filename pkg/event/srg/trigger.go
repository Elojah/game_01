package srg

import (
	"strings"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

const (
	triggerKey = "trigger:"
)

// InsertTrigger implemented with redis.
func (s *Store) InsertTrigger(t event.Trigger) error {
	return errors.Wrapf(
		s.Set(
			triggerKey+t.EventSourceID.String()+":"+t.EntityID.String(),
			t.EventTargetID.String(),
			0,
		).Err(),
		"insert trigger %s to %s for entity %s",
		t.EventSourceID.String(),
		t.EventTargetID.String(),
		t.EntityID.String(),
	)
}

// FetchTrigger redis implementation.
func (s *Store) FetchTrigger(triggerID gulid.ID, entityID gulid.ID) (gulid.ID, error) {
	val, err := s.Get(triggerKey + triggerID.String() + ":" + entityID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return gulid.ID{}, errors.Wrapf(err, "fetch trigger %s for entity %s", triggerID.String(), entityID.String())
		}
		return gulid.ID{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: triggerKey, Index: triggerID.String() + ":" + entityID.String()},
			"fetch trigger %s for entity %s",
			triggerID.String(),
			entityID.String(),
		)
	}

	return gulid.MustParse(val), nil
}

// ListTrigger redis implementation.
func (s *Store) ListTrigger(triggerID gulid.ID) ([]event.Trigger, error) {
	vals, err := s.Keys(triggerKey + triggerID.String() + ":*").Result()
	if err != nil {
		return nil, errors.Wrapf(err, "list trigger %s", triggerID.String())
	}
	triggers := make([]event.Trigger, len(vals))
	for i, val := range vals {
		etarget, err := s.Get(val).Result()
		if err != nil {
			return nil, errors.Wrapf(err, "list trigger %s", triggerID.String())
		}
		triggers[i] = event.Trigger{
			EntityID:      gulid.MustParse(strings.Split(val, ":")[2]),
			EventSourceID: triggerID,
			EventTargetID: gulid.MustParse(etarget),
		}
	}
	return triggers, nil
}

// RemoveTrigger redis implementation.
func (s *Store) RemoveTrigger(triggerID gulid.ID, entityID gulid.ID) error {
	return errors.Wrapf(
		s.Del(triggerKey+triggerID.String()+":"+entityID.String()).Err(),
		"remove trigger %s for entity %s",
		triggerID.String(),
		entityID.String(),
	)
}
