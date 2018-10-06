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
	return s.ZAddNX(
		eventKey+id.String(),
		redis.Z{
			Score:  float64(e.ID.Time()),
			Member: raw,
		},
	).Err()
}

// ListEvent list events in redis set key from min (included).
func (s *Store) ListEvent(id ulid.ID, min ulid.ID) ([]event.E, error) {
	vals, err := s.ZRangeByScore(
		eventKey+id.String(),
		redis.ZRangeBy{
			Min: strconv.FormatUint(min.Time(), 10),
			Max: "+inf",
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
