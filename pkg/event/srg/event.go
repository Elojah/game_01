package srg

import (
	"strconv"

	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	eventKey = "event:"
)

// SetEvent implemented with redis.
func (s *Store) SetEvent(e event.E, id ulid.ID) error {
	raw, err := e.Marshal()
	if err != nil {
		return err
	}
	return redis.NewIntCmd(
		"ZADD",
		eventKey+id.String(),
		"NX",
		e.TS.UnixNano(),
		raw,
	).Err()
}

// ListEvent retrieves event in Redis using ZRangeWithScores.
func (s *Store) ListEvent(subset event.Subset) ([]event.E, error) {
	cmd := s.ZRangeByScore(
		eventKey+subset.Key,
		redis.ZRangeBy{
			Min: strconv.FormatInt(subset.Min, 10),
			Max: "+inf",
		},
	)
	vals, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	events := make([]event.E, len(vals))
	for i := range vals {
		if err := events[i].Unmarshal([]byte(vals[i])); err != nil {
			return nil, err
		}
	}
	return events, nil
}
