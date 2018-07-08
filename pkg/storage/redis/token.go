package redis

import (
	"strconv"

	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/storage"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	tokenKey   = "token:"
	tokenHCKey = "token_hc:"
)

// GetToken redis implementation.
func (s *Service) GetToken(subset account.TokenSubset) (account.Token, error) {
	val, err := s.Get(tokenKey + subset.ID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return account.Token{}, err
		}
		return account.Token{}, storage.ErrNotFound
	}

	var token storage.Token
	if _, err := token.Unmarshal([]byte(val)); err != nil {
		return account.Token{}, err
	}
	return token.Domain()
}

// SetToken redis implementation.
func (s *Service) SetToken(token account.Token) error {
	raw, err := storage.NewToken(token).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(tokenKey+token.ID.String(), raw, 0).Err()
}

// DelToken redis implementation.
func (s *Service) DelToken(subset account.TokenSubset) error {
	return s.Del(tokenKey + subset.ID.String()).Err()
}

// SetTokenHC redis implementation.
func (s *Service) SetTokenHC(id ulid.ID, hc int64) error {
	return s.ZAddNX(
		tokenHCKey,
		redis.Z{
			Score:  float64(hc),
			Member: id.String(),
		},
	).Err()
}

// ListTokenHC redis implementation.
func (s *Service) ListTokenHC(subset account.TokenHCSubset) ([]ulid.ID, error) {
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
	ids := make([]ulid.ID, len(vals))
	for i, val := range vals {
		ids[i] = ulid.MustParse(val)
	}
	return ids, nil
}
