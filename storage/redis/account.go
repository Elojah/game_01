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
func (s *Service) GetAccount(builder game.AccountBuilder) (game.Account, error) {
	val, err := s.Get(accountKey + builder.Username).Result()
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
	return account.Domain()
}
