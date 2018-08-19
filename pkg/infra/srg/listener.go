package srg

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/infra"
)

const (
	listenerKey = "listener:"
)

// GetListener redis implementation.
func (s *Store) GetListener(subset infra.ListenerSubset) (infra.Listener, error) {
	val, err := s.Get(listenerKey + subset.ID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return infra.Listener{}, err
		}
		return infra.Listener{}, errors.ErrNotFound
	}

	var listener infra.Listener
	if err := listener.Unmarshal([]byte(val)); err != nil {
		return infra.Listener{}, err
	}
	return listener, nil
}

// SetListener redis implementation.
func (s *Store) SetListener(listener infra.Listener) error {
	raw, err := listener.Marshal()
	if err != nil {
		return err
	}
	return s.Set(listenerKey+listener.ID.String(), raw, 0).Err()
}

// DelListener deletes listener in redis.
func (s *Store) DelListener(subset infra.ListenerSubset) error {
	return s.Del(listenerKey + subset.ID.String()).Err()
}
