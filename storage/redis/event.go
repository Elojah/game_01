package redis

import (
	"strconv"

	"github.com/go-redis/redis"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

const (
	eventKey = "event:"
)

// CreateEvent implemented with redis.
func (s *Service) CreateEvent(event game.Event, id game.ID) error {
	raw, err := storage.NewEvent(event).Marshal(nil)
	if err != nil {
		return err
	}
	return s.ZAddNX(
		eventKey+id.String(),
		redis.Z{
			Score:  float64(event.TS.UnixNano()),
			Member: raw,
		},
	).Err()
}

// ListEvent retrieves event in Redis using ZRangeWithScores.
func (s *Service) ListEvent(subset game.EventSubset) ([]game.Event, error) {
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
	events := make([]game.Event, len(vals))
	for i := range vals {
		var eventS storage.Event
		if _, err := eventS.Unmarshal([]byte(vals[i])); err != nil {
			return nil, err
		}
		events[i] = eventS.Domain()
	}
	return events, nil
}
