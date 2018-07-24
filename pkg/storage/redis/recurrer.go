package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/storage"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	recurrerKey = "recurrer:"
)

// GetRecurrer redis implementation.
func (s *Service) GetRecurrer(subset infra.RecurrerSubset) (infra.Recurrer, error) {
	val, err := s.Get(recurrerKey + ulid.String(subset.TokenID)).Result()
	if err != nil {
		if err != redis.Nil {
			return infra.Recurrer{}, err
		}
		return infra.Recurrer{}, storage.ErrNotFound
	}

	var recurrer infra.Recurrer
	if _, err := recurrer.Unmarshal([]byte(val)); err != nil {
		return infra.Recurrer{}, err
	}
	return recurrer, nil
}

// SetRecurrer redis implementation.
func (s *Service) SetRecurrer(recurrer infra.Recurrer) error {
	raw, err := recurrer.Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(recurrerKey+ulid.String(recurrer.TokenID), raw, 0).Err()
}

// DelRecurrer deletes recurrer in redis.
func (s *Service) DelRecurrer(subset infra.RecurrerSubset) error {
	return s.Del(recurrerKey + ulid.String(subset.TokenID)).Err()
}
