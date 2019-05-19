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

// FetchPC implemented with redis.
func (s *Store) FetchPC(accountID ulid.ID, id ulid.ID) (entity.PC, error) {
	val, err := s.Get(pcKey + accountID.String() + ":" + id.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.PC{}, errors.Wrapf(err, "fetch pc %s for account %s", id.String(), accountID.String())
		}
		return entity.PC{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: pcKey, Index: accountID.String() + ":" + id.String()},
			"fetch pc %s for account %s",
			id.String(),
			accountID.String(),
		)
	}

	var e entity.E
	if err := e.Unmarshal([]byte(val)); err != nil {
		return entity.PC{}, errors.Wrapf(err, "fetch pc %s for account %s", id.String(), accountID.String())
	}
	return e, nil
}

// InsertPC implemented with redis.
func (s *Store) InsertPC(pc entity.PC, accountID ulid.ID) error {
	raw, err := pc.Marshal()
	if err != nil {
		return errors.Wrapf(err, "insert pc %s for account %s", pc.ID.String(), accountID.String())
	}
	return errors.Wrapf(
		s.Set(pcKey+accountID.String()+":"+pc.ID.String(), raw, 0).Err(),
		"insert pc %s for account %s",
		pc.ID.String(),
		accountID.String(),
	)
}

// RemovePC implemented with redis.
func (s *Store) RemovePC(accountID ulid.ID, id ulid.ID) error {
	return errors.Wrapf(
		s.Del(pcKey+accountID.String()+":"+id.String()).Err(),
		"remove pc %s for account %s",
		id.String(),
		accountID.String(),
	)
}

// InsertPCLeft implemented with redis.
func (s *Store) InsertPCLeft(pc entity.PCLeft, accountID ulid.ID) error {
	return errors.Wrapf(
		s.Set(pcLeftKey+accountID.String(), int(pc), 0).Err(),
		"insert pc left at %d for %s",
		pc,
		accountID.String(),
	)
}

// FetchPCLeft implemented with redis.
func (s *Store) FetchPCLeft(accountID ulid.ID) (entity.PCLeft, error) {
	val, err := s.Get(pcLeftKey + accountID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.PCLeft(0), errors.Wrapf(err, "fetch pc left for account %s", accountID.String())
		}
		return entity.PCLeft(0), errors.Wrapf(
			gerrors.ErrNotFound{Store: pcLeftKey, Index: accountID.String()},
			"fetch pc left for account %s",
			accountID.String(),
		)
	}

	pcLeft, err := strconv.Atoi(val)
	return entity.PCLeft(pcLeft), errors.Wrapf(err, "fetch pc left for account %s", accountID.String())
}

// RemovePCLeft implemented with redis.
func (s *Store) RemovePCLeft(accountID ulid.ID) error {
	return errors.Wrapf(
		s.Del(pcLeftKey+accountID.String()).Err(),
		"remove pc left for account %s",
		accountID.String(),
	)
}
