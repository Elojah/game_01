package srg

import (
	"strconv"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/account"
	gerrors "github.com/elojah/game_01/pkg/errors"
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
			return account.Token{}, errors.Wrapf(err, "get token %s", id.String())
		}
		return account.Token{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: tokenKey, Index: id.String()},
			"get token %s",
			id.String(),
		)
	}

	var t account.Token
	if err := t.Unmarshal([]byte(val)); err != nil {
		return account.Token{}, errors.Wrapf(err, "get token %s", id.String())
	}
	return t, nil
}

// SetToken redis implementation.
func (s *Store) SetToken(t account.Token) error {
	raw, err := t.Marshal()
	if err != nil {
		return errors.Wrapf(err, "set token %s", t.ID.String())
	}
	return errors.Wrapf(s.Set(tokenKey+t.ID.String(), raw, 0).Err(), "set token %s", t.ID.String())
}

// DelToken redis implementation.
func (s *Store) DelToken(id ulid.ID) error {
	return errors.Wrapf(s.Del(tokenKey+id.String()).Err(), "del token %s", id.String())
}

// SetTokenHC redis implementation.
func (s *Store) SetTokenHC(id ulid.ID, hc uint64) error {
	return errors.Wrapf(s.ZAddNX(
		tokenHCKey,
		redis.Z{
			Score:  float64(hc),
			Member: id.String(),
		},
	).Err(), "set token hc %s at %d", id.String(), hc)
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
		return nil, errors.Wrapf(err, "list token hc at %d", maxTS)
	}
	ids := make([]ulid.ID, len(vals))
	for i, val := range vals {
		ids[i] = ulid.MustParse(val)
	}
	return ids, nil
}
