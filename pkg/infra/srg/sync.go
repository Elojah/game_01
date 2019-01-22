package srg

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	syncKey = "sync:"
)

// GetRandomSync redis implementation.
func (s *Store) GetRandomSync() (infra.Sync, error) {
	val, err := s.SRandMember(syncKey).Result()
	if err != nil {
		if err != redis.Nil {
			return infra.Sync{}, errors.Wrap(err, "get random sync")
		}
		return infra.Sync{}, errors.Wrap(
			gerrors.ErrNotFound{Store: syncKey, Index: "random"},
			"get random sync",
		)
	}

	return infra.Sync{ID: ulid.MustParse(val)}, nil
}

// SetSync redis implementation.
func (s *Store) SetSync(sync infra.Sync) error {
	return errors.Wrapf(s.SAdd(
		syncKey,
		sync.ID.String(),
	).Err(), "set sync %s", sync.ID.String())
}

// DelSync redis implementation.
func (s *Store) DelSync(id ulid.ID) error {
	return errors.Wrapf(s.SRem(
		syncKey,
		id.String(),
	).Err(), "del sync %s", id.String())
}
