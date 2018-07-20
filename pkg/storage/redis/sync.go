package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/storage"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	syncKey = "sync:"
)

// GetRandomSync redis implementation.
func (s *Service) GetRandomSync(subset infra.SyncSubset) (infra.Sync, error) {
	val, err := s.SRandMember(syncKey).Result()
	if err != nil {
		if err != redis.Nil {
			return infra.Sync{}, err
		}
		return infra.Sync{}, storage.ErrNotFound
	}

	return infra.Sync{ID: ulid.MustParse(val)}, nil
}

// SetSync redis implementation.
func (s *Service) SetSync(sync infra.Sync) error {
	return s.SAdd(
		syncKey,
		ulid.String(sync.ID),
	).Err()
}

// DelSync redis implementation.
func (s *Service) DelSync(subset infra.SyncSubset) error {
	return s.SRem(
		syncKey,
		ulid.String(subset.ID),
	).Err()
}
