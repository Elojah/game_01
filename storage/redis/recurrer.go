package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/storage"
)

const (
	recurrerKey = "recurrer:"
)

// GetRecurrer redis implementation.
func (s *Service) GetRecurrer(subset event.RecurrerSubset) (event.Recurrer, error) {
	val, err := s.Get(recurrerKey + subset.ID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return event.Recurrer{}, err
		}
		return event.Recurrer{}, storage.ErrNotFound
	}

	var recurrer storage.Recurrer
	if _, err := recurrer.Unmarshal([]byte(val)); err != nil {
		return event.Recurrer{}, err
	}
	return recurrer.Domain(), nil
}

// SetRecurrer redis implementation.
func (s *Service) SetRecurrer(recurrer event.Recurrer) error {
	raw, err := storage.NewRecurrer(recurrer).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(recurrerKey+recurrer.ID.String(), raw, 0).Err()
}

// DelRecurrer deletes recurrer in redis.
func (s *Service) DelRecurrer(subset event.RecurrerSubset) error {
	return s.Del(recurrerKey + subset.ID.String()).Err()
}
