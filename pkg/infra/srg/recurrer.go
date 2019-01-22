package srg

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	recurrerKey = "recurrer:"
)

// GetRecurrer redis implementation.
func (s *Store) GetRecurrer(tokenID ulid.ID) (infra.Recurrer, error) {
	val, err := s.Get(recurrerKey + tokenID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return infra.Recurrer{}, errors.Wrapf(err, "get recurrer for token %s", tokenID.String())
		}
		return infra.Recurrer{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: recurrerKey, Index: tokenID.String()},
			"get recurrer for token %s",
			tokenID.String(),
		)
	}

	var recurrer infra.Recurrer
	if err := recurrer.Unmarshal([]byte(val)); err != nil {
		return infra.Recurrer{}, errors.Wrapf(err, "get recurrer for token %s", tokenID.String())
	}
	return recurrer, nil
}

// SetRecurrer redis implementation.
func (s *Store) SetRecurrer(recurrer infra.Recurrer) error {
	raw, err := recurrer.Marshal()
	if err != nil {
		return errors.Wrapf(err, "set recurrer for token %s", recurrer.TokenID.String())
	}
	return errors.Wrapf(
		s.Set(recurrerKey+recurrer.TokenID.String(), raw, 0).Err(),
		"set recurrer for token %s",
		recurrer.TokenID.String(),
	)
}

// DelRecurrer deletes recurrer in redis.
func (s *Store) DelRecurrer(tokenID ulid.ID) error {
	return errors.Wrapf(s.Del(recurrerKey+tokenID.String()).Err(), "del recurrer for token %s", tokenID.String())
}
