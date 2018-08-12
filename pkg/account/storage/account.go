package storage

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/errors"
)

const (
	accountKey = "account:"
)

// GetAccount implemented with redis.
func (s *Service) GetAccount(subset account.Subset) (account.A, error) {
	val, err := s.Get(accountKey + subset.Username).Result()
	if err != nil {
		if err != redis.Nil {
			return account.A{}, err
		}
		return account.A{}, errors.ErrNotFound
	}

	var a account.A
	if err := a.Unmarshal([]byte(val)); err != nil {
		return account.A{}, err
	}
	return a, nil
}

// SetAccount implemented with redis.
func (s *Service) SetAccount(a account.A) error {
	raw, err := a.Marshal()
	if err != nil {
		return err
	}
	return s.Set(accountKey+a.Username, raw, 0).Err()
}

// DelAccount redis implementation.
func (s *Service) DelAccount(subset account.Subset) error {
	return s.Del(accountKey + subset.Username).Err()
}
