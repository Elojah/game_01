package redis

import (
	"time"

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
func (s *Service) ListEvent(builder game.EventBuilder) ([]game.Event, error) {
	cmd := s.ZRangeWithScores(
		builder.Key,
		builder.Start,
		time.Now().UnixNano(),
	)
	vals, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	events := make([]game.Event, len(vals))
	for i := range vals {
		events[i] = vals[i].Member.(storage.Event).Domain()
	}
	return events, nil
}
