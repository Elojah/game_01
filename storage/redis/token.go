package redis

import (
	"strconv"

	"github.com/go-redis/redis"
	"github.com/oklog/ulid"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

const (
	tokenKey   = "token:"
	tokenHCKey = "token_hc:"
)

// GetToken redis implementation.
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
	return token.Domain(id)
}

// SetToken redis implementation.
func (s *Service) SetToken(token game.Token) error {
	raw, err := storage.NewToken(token).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(tokenKey+token.ID.String(), raw, 0).Err()
}

// SetTokenHC redis implementation.
func (s *Service) SetTokenHC(id game.ID, hc int64) error {
	return s.ZAddNX(
		tokenHCKey,
		redis.Z{
			Score:  float64(hc),
			Member: id.String(),
		},
	).Err()
}

// ListTokenHC redis implementation.
func (s *Service) ListTokenHC(subset game.TokenHCSubset) ([]game.ID, error) {
	cmd := s.ZRangeByScore(
		tokenHCKey,
		redis.ZRangeBy{
			Min: "-inf",
			Max: strconv.FormatInt(subset.MaxTS, 10),
		},
	)
	vals, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	ids := make([]game.ID, len(vals))
	for i, val := range vals {
		ids[i], err = ulid.Parse(val)
		if err != nil {
			return nil, err
		}
	}
	return ids, nil
}
