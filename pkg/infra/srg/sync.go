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

// FetchRandomSync redis implementation.
func (s *Store) FetchRandomSync() (infra.Sync, error) {
	val, err := s.SRandMember(syncKey).Result()
	if err != nil {
		if err != redis.Nil {
			return infra.Sync{}, errors.Wrap(err, "fetch random sync")
		}
		return infra.Sync{}, errors.Wrap(
			gerrors.ErrNotFound{Store: syncKey, Index: "random"},
			"fetch random sync",
		)
	}

	return infra.Sync{ID: ulid.MustParse(val)}, nil
}

// UpsertSync redis implementation.
func (s *Store) UpsertSync(sync infra.Sync) error {
	return errors.Wrapf(s.SAdd(
		syncKey,
		sync.ID.String(),
	).Err(), "upsert sync %s", sync.ID.String())
}

// RemoveSync redis implementation.
func (s *Store) RemoveSync(id ulid.ID) error {
	return errors.Wrapf(s.SRem(
		syncKey,
		id.String(),
	).Err(), "remove sync %s", id.String())
}
