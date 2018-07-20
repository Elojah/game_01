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
	val, err := s.Get(tokenKey + ulid.String(subset.ID)).Result()
	if err != nil {
		if err != redis.Nil {
			return account.Token{}, err
		}
		return account.Token{}, storage.ErrNotFound
	}

	var t account.Token
	if _, err := t.Unmarshal([]byte(val)); err != nil {
		return account.Token{}, err
	}
	return t, nil
}

// SetToken redis implementation.
func (s *Service) SetToken(t account.Token) error {
	raw, err := t.Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(tokenKey+ulid.String(t.ID), raw, 0).Err()
}

// DelToken redis implementation.
func (s *Service) DelToken(subset account.TokenSubset) error {
	return s.Del(tokenKey + ulid.String(subset.ID)).Err()
}

// SetTokenHC redis implementation.
func (s *Service) SetTokenHC(id ulid.ID, hc int64) error {
	return s.ZAddNX(
		tokenHCKey,
		redis.Z{
			Score:  float64(hc),
			Member: ulid.String(id),
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
