package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/storage"
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
		return account.A{}, storage.ErrNotFound
	}

	var a storage.Account
	if _, err := a.Unmarshal([]byte(val)); err != nil {
		return account.A{}, err
	}
	return a.Domain(subset.Username)
}

// SetAccount implemented with redis.
func (s *Service) SetAccount(a account.A) error {
	raw, err := storage.NewAccount(a).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(accountKey+a.Username, raw, 0).Err()
}
