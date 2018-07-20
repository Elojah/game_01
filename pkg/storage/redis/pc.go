package redis

import (
	"strconv"

	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/storage"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	pcKey     = "pc:"
	pcLeftKey = "pc_left:"
)

// ListPC implemented with redis.
func (s *Service) ListPC(subset entity.PCSubset) ([]entity.PC, error) {
	keys, err := s.Keys(pcKey + ulid.String(subset.AccountID) + "*").Result()
	if err != nil {
		return nil, err
	}

	pcs := make([]entity.PC, len(keys))
	for i, key := range keys {
		val, err := s.Get(key).Result()
		if err != nil {
			return nil, err
		}

		if _, err := pcs[i].Unmarshal([]byte(val)); err != nil {
			return nil, err
		}
	}
	return pcs, nil
}

// GetPC implemented with redis.
func (s *Service) GetPC(subset entity.PCSubset) (entity.PC, error) {
	val, err := s.Get(pcKey + ulid.String(subset.AccountID) + ":" + ulid.String(subset.ID)).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.PC{}, err
		}
		return entity.PC{}, storage.ErrNotFound
	}

	var e entity.E
	if _, err := e.Unmarshal([]byte(val)); err != nil {
		return entity.PC{}, err
	}
	return entity.PC(e), nil
}

// SetPC implemented with redis.
func (s *Service) SetPC(pc entity.PC, account ulid.ID) error {
	raw, err := pc.Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(pcKey+ulid.String(account)+":"+ulid.String(pc.ID), raw, 0).Err()
}

// SetPCLeft implemented with redis.
func (s *Service) SetPCLeft(pc entity.PCLeft, account ulid.ID) error {
	return s.Set(pcLeftKey+ulid.String(account), int(pc), 0).Err()
}

// GetPCLeft implemented with redis.
func (s *Service) GetPCLeft(subset entity.PCLeftSubset) (entity.PCLeft, error) {
	val, err := s.Get(pcLeftKey + ulid.String(subset.AccountID)).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.PCLeft(0), err
		}
		return entity.PCLeft(0), storage.ErrNotFound
	}

	pcLeft, err := strconv.Atoi(val)
	return entity.PCLeft(pcLeft), err
}
