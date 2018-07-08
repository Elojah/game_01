package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/storage"
)

const (
	listenerKey = "listener:"
)

// GetListener redis implementation.
func (s *Service) GetListener(subset event.ListenerSubset) (event.Listener, error) {
	val, err := s.Get(listenerKey + subset.ID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return event.Listener{}, err
		}
		return event.Listener{}, storage.ErrNotFound
	}

	var listener storage.Listener
	if _, err := listener.Unmarshal([]byte(val)); err != nil {
		return event.Listener{}, err
	}
	return listener.Domain(), nil
}

// SetListener redis implementation.
func (s *Service) SetListener(listener event.Listener) error {
	raw, err := storage.NewListener(listener).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(listenerKey+listener.ID.String(), raw, 0).Err()
}

// DelListener deletes listener in redis.
func (s *Service) DelListener(subset event.ListenerSubset) error {
	return s.Del(listenerKey + subset.ID.String()).Err()
}
