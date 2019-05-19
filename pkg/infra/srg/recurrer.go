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

// FetchRecurrer redis implementation.
func (s *Store) FetchRecurrer(tokenID ulid.ID) (infra.Recurrer, error) {
	val, err := s.Get(recurrerKey + tokenID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return infra.Recurrer{}, errors.Wrapf(err, "fetch recurrer for token %s", tokenID.String())
		}
		return infra.Recurrer{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: recurrerKey, Index: tokenID.String()},
			"fetch recurrer for token %s",
			tokenID.String(),
		)
	}

	var recurrer infra.Recurrer
	if err := recurrer.Unmarshal([]byte(val)); err != nil {
		return infra.Recurrer{}, errors.Wrapf(err, "fetch recurrer for token %s", tokenID.String())
	}
	return recurrer, nil
}

// InsertRecurrer redis implementation.
func (s *Store) InsertRecurrer(recurrer infra.Recurrer) error {
	raw, err := recurrer.Marshal()
	if err != nil {
		return errors.Wrapf(err, "insert recurrer for token %s", recurrer.TokenID.String())
	}
	return errors.Wrapf(
		s.Set(recurrerKey+recurrer.TokenID.String(), raw, 0).Err(),
		"insert recurrer for token %s",
		recurrer.TokenID.String(),
	)
}

// RemoveRecurrer deletes recurrer in redis.
func (s *Store) RemoveRecurrer(tokenID ulid.ID) error {
	return errors.Wrapf(s.Del(recurrerKey+tokenID.String()).Err(), "remove recurrer for token %s", tokenID.String())
}
