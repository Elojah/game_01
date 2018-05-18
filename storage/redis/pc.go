package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

const (
	pcKey = "pc:"
)

// ListPC implemented with redis.
func (s *Service) ListPC(subset game.PCSubset) ([]game.PC, error) {
	keys, err := s.Keys(pcKey + subset.AccountID.String() + "*").Result()
	if err != nil {
		if err != redis.Nil {
			return nil, err
		}
		return nil, storage.ErrNotFound
	}

	pcs := make([]game.PC, len(keys))
	for i, key := range keys {
		val, err := s.Get(key).Result()
		if err != nil {
			return nil, err
		}

		var entity storage.Entity
		if _, err := entity.Unmarshal([]byte(val)); err != nil {
			return nil, err
		}
		pcs[i] = game.PC(entity.Domain())
	}
	return pcs, nil
}

// SetPC implemented with redis.
func (s *Service) SetPC(pc game.PC, account game.ID) error {
	raw, err := storage.NewEntity(game.Entity(pc)).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(pcKey+account.String()+":"+pc.ID.String(), raw, 0).Err()
}
