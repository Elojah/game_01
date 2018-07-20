package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/storage"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	listenerKey = "listener:"
)

// GetListener redis implementation.
func (s *Service) GetListener(subset event.ListenerSubset) (event.Listener, error) {
	val, err := s.Get(listenerKey + ulid.String(subset.ID)).Result()
	if err != nil {
		if err != redis.Nil {
			return event.Listener{}, err
		}
		return event.Listener{}, storage.ErrNotFound
	}

	var listener event.Listener
	if _, err := listener.Unmarshal([]byte(val)); err != nil {
		return event.Listener{}, err
	}
	return listener, nil
}

// SetListener redis implementation.
func (s *Service) SetListener(listener event.Listener) error {
	raw, err := listener.Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(listenerKey+ulid.String(listener.ID), raw, 0).Err()
}

// DelListener deletes listener in redis.
func (s *Service) DelListener(subset event.ListenerSubset) error {
	return s.Del(listenerKey + ulid.String(subset.ID)).Err()
}
