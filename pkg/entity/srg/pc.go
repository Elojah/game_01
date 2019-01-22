package srg

import (
	"strconv"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	pcKey     = "pc:"
	pcLeftKey = "pc_left:"
)

// ListPC implemented with redis.
func (s *Store) ListPC(accountID ulid.ID) ([]entity.PC, error) {
	keys, err := s.Keys(pcKey + accountID.String() + "*").Result()
	if err != nil {
		return nil, errors.Wrapf(err, "list pc for account %s", accountID.String())
	}

	pcs := make([]entity.PC, len(keys))
	for i, key := range keys {
		val, err := s.Get(key).Result()
		if err != nil {
			return nil, errors.Wrapf(err, "list pc for account %s", accountID.String())
		}

		if err := pcs[i].Unmarshal([]byte(val)); err != nil {
			return nil, errors.Wrapf(err, "list pc for account %s", accountID.String())
		}
	}
	return pcs, nil
}

// GetPC implemented with redis.
func (s *Store) GetPC(accountID ulid.ID, id ulid.ID) (entity.PC, error) {
	val, err := s.Get(pcKey + accountID.String() + ":" + id.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.PC{}, errors.Wrapf(err, "get pc %s for account %s", id.String(), accountID.String())
		}
		return entity.PC{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: pcKey, Index: accountID.String() + ":" + id.String()},
			"get pc %s for account %s",
			id.String(),
			accountID.String(),
		)
	}

	var e entity.E
	if err := e.Unmarshal([]byte(val)); err != nil {
		return entity.PC{}, errors.Wrapf(err, "get pc %s for account %s", id.String(), accountID.String())
	}
	return entity.PC(e), nil
}

// SetPC implemented with redis.
func (s *Store) SetPC(pc entity.PC, accountID ulid.ID) error {
	raw, err := pc.Marshal()
	if err != nil {
		return errors.Wrapf(err, "set pc %s for account", pc.ID.String(), accountID.String())
	}
	return errors.Wrapf(
		s.Set(pcKey+accountID.String()+":"+pc.ID.String(), raw, 0).Err(),
		"set pc %s for account",
		pc.ID.String(),
		accountID.String(),
	)
}

// DelPC implemented with redis.
func (s *Store) DelPC(accountID ulid.ID, id ulid.ID) error {
	return errors.Wrapf(
		s.Del(pcKey+accountID.String()+":"+id.String()).Err(),
		"del pc %s for account %s",
		id.String(),
		accountID.String(),
	)
}

// SetPCLeft implemented with redis.
func (s *Store) SetPCLeft(pc entity.PCLeft, accountID ulid.ID) error {
	return errors.Wrapf(
		s.Set(pcLeftKey+accountID.String(), int(pc), 0).Err(),
		"set pc left at %d for %s",
		pc,
		accountID.String(),
	)
}

// GetPCLeft implemented with redis.
func (s *Store) GetPCLeft(accountID ulid.ID) (entity.PCLeft, error) {
	val, err := s.Get(pcLeftKey + accountID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.PCLeft(0), errors.Wrapf(err, "get pc left for account %s", accountID.String())
		}
		return entity.PCLeft(0), errors.Wrapf(
			gerrors.ErrNotFound{Store: pcLeftKey, Index: accountID.String()},
			"get pc left for account %s",
			accountID.String(),
		)
	}

	pcLeft, err := strconv.Atoi(val)
	return entity.PCLeft(pcLeft), errors.Wrapf(err, "get pc left for account %s", accountID.String())
}

// DelPCLeft implemented with redis.
func (s *Store) DelPCLeft(accountID ulid.ID) error {
	return errors.Wrapf(
		s.Del(pcLeftKey+accountID.String()).Err(),
		"del pc left for account %s",
		accountID.String(),
	)
}
