package srg

import (
	"strconv"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

const (
	eventKey = "event:"
)

// Upsert implemented with redis.
func (s *Store) Upsert(e event.E, id gulid.ID) error {
	raw, err := e.Marshal()
	if err != nil {
		return errors.Wrapf(err, "upsert event %s for entity %s", e.ID.String(), id.String())
	}
	return errors.Wrapf(s.ZAddNX(
		eventKey+id.String(),
		redis.Z{
			Score:  float64(e.ID.Time()),
			Member: raw,
		},
	).Err(), "upsert event %s for entity %s", e.ID.String(), id.String())
}

// Fetch implemented with redis.
func (s *Store) Fetch(id gulid.ID, entityID gulid.ID) (event.E, error) {
	vals, err := s.ZRangeByScore(
		eventKey+entityID.String(),
		redis.ZRangeBy{
			Min:   strconv.FormatUint(id.Time(), 10),
			Max:   "+inf",
			Count: 1,
		},
	).Result()
	if err != nil {
		return event.E{}, errors.Wrapf(err, "fetch event for entity %s at %d", entityID.String(), id.Time())
	}
	if len(vals) != 0 {
		return event.E{}, errors.Wrapf(gerrors.ErrNotFound{
			Store: eventKey,
			Index: eventKey + entityID.String(),
		}, "fetch event at %d", id.Time())
	}

	var e event.E
	if err := e.Unmarshal([]byte(vals[0])); err != nil {
		return event.E{}, errors.Wrapf(err, "fetch event for entity %s at %d", entityID.String(), id.Time())
	}
	return e, nil
}

// List list events in redis upsert key from min (included).
func (s *Store) List(id gulid.ID, min gulid.ID) ([]event.E, error) {
	vals, err := s.ZRangeByScore(
		eventKey+id.String(),
		redis.ZRangeBy{
			Min: strconv.FormatUint(min.Time(), 10),
			Max: "+inf",
		},
	).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "list events for entity %s at %d", id.String(), min.Time())
	}
	events := make([]event.E, len(vals))
	for i := range vals {
		if err := events[i].Unmarshal([]byte(vals[i])); err != nil {
			return nil, errors.Wrapf(err, "list events for entity %s at %d", id.String(), min.Time())
		}
	}
	return events, nil
}

// Remove implemented with redis.
func (s *Store) Remove(eventID gulid.ID, id gulid.ID) error {
	return errors.Wrapf(s.ZRem(
		eventKey+id.String(),
		eventID.String(),
	).Err(), "remove event %s for entity %s", eventID.String(), id.String())
}
