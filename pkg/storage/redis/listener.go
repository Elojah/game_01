package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/storage"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	listenerKey = "listener:"
)

// GetListener redis implementation.
func (s *Service) GetListener(subset infra.ListenerSubset) (infra.Listener, error) {
	val, err := s.Get(listenerKey + ulid.String(subset.ID)).Result()
	if err != nil {
		if err != redis.Nil {
			return infra.Listener{}, err
		}
		return infra.Listener{}, storage.ErrNotFound
	}

	var listener infra.Listener
	if _, err := listener.Unmarshal([]byte(val)); err != nil {
		return infra.Listener{}, err
	}
	return listener, nil
}

// SetListener redis implementation.
func (s *Service) SetListener(listener infra.Listener) error {
	raw, err := listener.Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(listenerKey+ulid.String(listener.ID), raw, 0).Err()
}

// DelListener deletes listener in redis.
func (s *Service) DelListener(subset infra.ListenerSubset) error {
	return s.Del(listenerKey + ulid.String(subset.ID)).Err()
}
