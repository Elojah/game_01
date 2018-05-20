package redis

import (
	"strconv"

	"github.com/go-redis/redis"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

const (
	pcKey     = "pc:"
	pcLeftKey = "pc_left:"
)

// ListPC implemented with redis.
func (s *Service) ListPC(subset game.PCSubset) ([]game.PC, error) {
	keys, err := s.Keys(pcKey + subset.AccountID.String() + "*").Result()
	if err != nil {
		return nil, err
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

// GetPC implemented with redis.
func (s *Service) GetPC(subset game.PCSubset) (game.PC, error) {
	val, err := s.Get(pcKey + subset.AccountID.String() + ":" + subset.ID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return game.PC{}, err
		}
		return game.PC{}, storage.ErrNotFound
	}

	var entity storage.Entity
	if _, err := entity.Unmarshal([]byte(val)); err != nil {
		return game.PC{}, err
	}
	return game.PC(entity.Domain()), nil
}

// SetPC implemented with redis.
func (s *Service) SetPC(pc game.PC, account game.ID) error {
	raw, err := storage.NewEntity(game.Entity(pc)).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(pcKey+account.String()+":"+pc.ID.String(), raw, 0).Err()
}

// SetPCLeft implemented with redis.
func (s *Service) SetPCLeft(pc game.PCLeft, account game.ID) error {
	return s.Set(pcLeftKey+account.String(), int(pc), 0).Err()
}

// GetPCLeft implemented with redis.
func (s *Service) GetPCLeft(subset game.PCLeftSubset) (game.PCLeft, error) {
	val, err := s.Get(pcLeftKey + subset.AccountID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return game.PCLeft(0), err
		}
		return game.PCLeft(0), storage.ErrNotFound
	}

	pcLeft, err := strconv.Atoi(val)
	return game.PCLeft(pcLeft), err
}
