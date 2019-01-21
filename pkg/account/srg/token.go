package srg

import (
	"strconv"

	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	tokenKey   = "token:"
	tokenHCKey = "token_hc:"
)

// GetToken redis implementation.
func (s *Store) GetToken(id ulid.ID) (account.Token, error) {
	val, err := s.Get(tokenKey + id.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return account.Token{}, err
		}
		return account.Token{}, errors.ErrNotFound{Store: tokenKey, Index: id.String()}
	}

	var t account.Token
	if err := t.Unmarshal([]byte(val)); err != nil {
		return account.Token{}, err
	}
	return t, nil
}

// SetToken redis implementation.
func (s *Store) SetToken(t account.Token) error {
	raw, err := t.Marshal()
	if err != nil {
		return err
	}
	return s.Set(tokenKey+t.ID.String(), raw, 0).Err()
}

// DelToken redis implementation.
func (s *Store) DelToken(id ulid.ID) error {
	return s.Del(tokenKey + id.String()).Err()
}

// SetTokenHC redis implementation.
func (s *Store) SetTokenHC(id ulid.ID, hc uint64) error {
	return s.ZAddNX(
		tokenHCKey,
		redis.Z{
			Score:  float64(hc),
			Member: id.String(),
		},
	).Err()
}

// ListTokenHC redis implementation.
func (s *Store) ListTokenHC(maxTS uint64) ([]ulid.ID, error) {
	cmd := s.ZRangeByScore(
		tokenHCKey,
		redis.ZRangeBy{
			Min: "-inf",
			Max: strconv.FormatUint(maxTS, 10),
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
