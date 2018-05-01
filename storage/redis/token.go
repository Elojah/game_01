package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

const (
	tokenKey = "token:"
)

// GetToken implemented with redis.
func (s *Service) GetToken(id game.ID) (game.Token, error) {
	val, err := s.Get(tokenKey + id.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return game.Token{}, err
		}
		return game.Token{}, storage.ErrNotFound
	}

	var token storage.Token
	if _, err := token.Unmarshal([]byte(val)); err != nil {
		return game.Token{}, err
	}
	return token.Domain()
}
