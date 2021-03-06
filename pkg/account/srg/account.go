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

// Fetch implemented with redis.
func (s *Store) Fetch(username string) (account.A, error) {
	val, err := s.Get(accountKey + username).Result()
	if err != nil {
		if err != redis.Nil {
			return account.A{}, errors.Wrapf(err, "fetch account %s", username)
		}
		return account.A{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: accountKey, Index: username},
			"fetch account %s",
			username,
		)
	}

	var a account.A
	if err := a.Unmarshal([]byte(val)); err != nil {
		return account.A{}, errors.Wrapf(err, "fetch account %s", username)
	}
	return a, nil
}

// Upsert implemented with redis.
func (s *Store) Upsert(a account.A) error {
	raw, err := a.Marshal()
	if err != nil {
		return errors.Wrapf(err, "upsert account %s", a.Username)
	}
	return errors.Wrapf(s.Set(accountKey+a.Username, raw, 0).Err(), "upsert account %s", a.Username)
}

// Remove redis implementation.
func (s *Store) Remove(username string) error {
	return errors.Wrapf(s.Del(accountKey+username).Err(), "remove account %s", username)
}
