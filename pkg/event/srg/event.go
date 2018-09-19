package srg

import (
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
	return s.ZAddNX(
		eventKey+id.String(),
		redis.Z{
			Score:  0, // default key for all events, must be the same for lexico order
			Member: raw,
		},
	).Err()
}

// ListEvent retrieves event in Redis using ZRangeWithScores.
func (s *Store) ListEvent(subset event.Subset) ([]event.E, error) {
	vals, err := s.ZRangeByLex(
		eventKey+subset.Key,
		redis.ZRangeBy{
			Min: "[" + subset.Min.String(),
			Max: "+",
		},
	).Result()
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
