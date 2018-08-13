package storage

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/infra"
)

const (
	recurrerKey = "recurrer:"
)

// GetRecurrer redis implementation.
func (s *Store) GetRecurrer(subset infra.RecurrerSubset) (infra.Recurrer, error) {
	val, err := s.Get(recurrerKey + subset.TokenID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return infra.Recurrer{}, err
		}
		return infra.Recurrer{}, errors.ErrNotFound
	}

	var recurrer infra.Recurrer
	if err := recurrer.Unmarshal([]byte(val)); err != nil {
		return infra.Recurrer{}, err
	}
	return recurrer, nil
}

// SetRecurrer redis implementation.
func (s *Store) SetRecurrer(recurrer infra.Recurrer) error {
	raw, err := recurrer.Marshal()
	if err != nil {
		return err
	}
	return s.Set(recurrerKey+recurrer.TokenID.String(), raw, 0).Err()
}

// DelRecurrer deletes recurrer in redis.
func (s *Store) DelRecurrer(subset infra.RecurrerSubset) error {
	return s.Del(recurrerKey + subset.TokenID.String()).Err()
}
