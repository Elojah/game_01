package srg

import (
	"strconv"

	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	pcKey     = "pc:"
	pcLeftKey = "pc_left:"
)

// ListPC implemented with redis.
func (s *Store) ListPC(subset entity.PCSubset) ([]entity.PC, error) {
	keys, err := s.Keys(pcKey + subset.AccountID.String() + "*").Result()
	if err != nil {
		return nil, err
	}

	pcs := make([]entity.PC, len(keys))
	for i, key := range keys {
		val, err := s.Get(key).Result()
		if err != nil {
			return nil, err
		}

		if err := pcs[i].Unmarshal([]byte(val)); err != nil {
			return nil, err
		}
	}
	return pcs, nil
}

// GetPC implemented with redis.
func (s *Store) GetPC(subset entity.PCSubset) (entity.PC, error) {
	val, err := s.Get(pcKey + subset.AccountID.String() + ":" + subset.ID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.PC{}, err
		}
		return entity.PC{}, errors.ErrNotFound
	}

	var e entity.E
	if err := e.Unmarshal([]byte(val)); err != nil {
		return entity.PC{}, err
	}
	return entity.PC(e), nil
}

// SetPC implemented with redis.
func (s *Store) SetPC(pc entity.PC, accountID ulid.ID) error {
	raw, err := pc.Marshal()
	if err != nil {
		return err
	}
	return s.Set(pcKey+accountID.String()+":"+pc.ID.String(), raw, 0).Err()
}

// DelPC implemented with redis.
func (s *Store) DelPC(subset entity.PCSubset) error {
	return s.Del(pcKey + subset.AccountID.String() + ":" + subset.ID.String()).Err()
}

// SetPCLeft implemented with redis.
func (s *Store) SetPCLeft(pc entity.PCLeft, accountID ulid.ID) error {
	return s.Set(pcLeftKey+accountID.String(), int(pc), 0).Err()
}

// GetPCLeft implemented with redis.
func (s *Store) GetPCLeft(subset entity.PCLeftSubset) (entity.PCLeft, error) {
	val, err := s.Get(pcLeftKey + subset.AccountID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.PCLeft(0), err
		}
		return entity.PCLeft(0), errors.ErrNotFound
	}

	pcLeft, err := strconv.Atoi(val)
	return entity.PCLeft(pcLeft), err
}

// DelPCLeft implemented with redis.
func (s *Store) DelPCLeft(subset entity.PCLeftSubset) error {
	return s.Del(pcLeftKey + subset.AccountID.String()).Err()
}