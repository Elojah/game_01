package srg

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	syncKey = "sync:"
)

// GetRandomSync redis implementation.
func (s *Store) GetRandomSync(subset infra.SyncSubset) (infra.Sync, error) {
	val, err := s.SRandMember(syncKey).Result()
	if err != nil {
		if err != redis.Nil {
			return infra.Sync{}, err
		}
		return infra.Sync{}, errors.ErrNotFound
	}

	return infra.Sync{ID: ulid.MustParse(val)}, nil
}

// SetSync redis implementation.
func (s *Store) SetSync(sync infra.Sync) error {
	return s.SAdd(
		syncKey,
		sync.ID.String(),
	).Err()
}

// DelSync redis implementation.
func (s *Store) DelSync(subset infra.SyncSubset) error {
	return s.SRem(
		syncKey,
		subset.ID.String(),
	).Err()
}