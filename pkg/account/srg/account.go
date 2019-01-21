package srg

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/account"
	gerrors "github.com/elojah/game_01/pkg/errors"
)

const (
	accountKey = "account:"
)

// GetAccount implemented with redis.
func (s *Store) GetAccount(username string) (account.A, error) {
	val, err := s.Get(accountKey + username).Result()
	if err != nil {
		if err != redis.Nil {
			return account.A{}, errors.Wrapf(err, "get account %s", username)
		}
		return account.A{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: accountKey, Index: username},
			"get account %s",
			username,
		)
	}

	var a account.A
	if err := a.Unmarshal([]byte(val)); err != nil {
		return account.A{}, errors.Wrapf(err, "get account %s", username)
	}
	return a, nil
}

// SetAccount implemented with redis.
func (s *Store) SetAccount(a account.A) error {
	raw, err := a.Marshal()
	if err != nil {
		return errors.Wrapf(err, "set account %s", a.Username)
	}
	return errors.Wrapf(s.Set(accountKey+a.Username, raw, 0).Err(), "set account %s", a.Username)
}

// DelAccount redis implementation.
func (s *Store) DelAccount(username string) error {
	return errors.Wrapf(s.Del(accountKey+username).Err(), "del account %s", username)
}
