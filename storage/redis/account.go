package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

const (
	accountKey = "account:"
)

// GetAccount implemented with redis.
func (s *Service) GetAccount(subset game.AccountSubset) (game.Account, error) {
	val, err := s.Get(accountKey + subset.Username).Result()
	if err != nil {
		if err != redis.Nil {
			return game.Account{}, err
		}
		return game.Account{}, storage.ErrNotFound
	}

	var account storage.Account
	if _, err := account.Unmarshal([]byte(val)); err != nil {
		return game.Account{}, err
	}
	return account.Domain(subset.Username)
}

// SetAccount implemented with redis.
func (s *Service) SetAccount(account game.Account) error {
	raw, err := storage.NewAccount(account).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(accountKey+account.Username, raw, 0).Err()
}
