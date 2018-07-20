package redis

import (
	"strconv"

	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/storage"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	eventKey = "event:"
)

// SetEvent implemented with redis.
func (s *Service) SetEvent(e event.E, id ulid.ID) error {
	raw, err := storage.NewEvent(e).Marshal(nil)
	if err != nil {
		return err
	}
	return s.ZAddNX(
		eventKey+ulid.String(id),
		redis.Z{
			Score:  float64(e.TS.UnixNano()),
			Member: raw,
		},
	).Err()
}

// ListEvent retrieves event in Redis using ZRangeWithScores.
func (s *Service) ListEvent(subset event.Subset) ([]event.E, error) {
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
		var eventS storage.Event
		if _, err := eventS.Unmarshal([]byte(vals[i])); err != nil {
			return nil, err
		}
		events[i] = eventS.Domain()
	}
	return events, nil
}
