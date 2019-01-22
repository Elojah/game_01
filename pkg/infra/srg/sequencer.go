package srg

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	sequencerKey = "sequencer:"
)

// GetSequencer redis implementation.
func (s *Store) GetSequencer(id ulid.ID) (infra.Sequencer, error) {
	val, err := s.Get(sequencerKey + id.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return infra.Sequencer{}, errors.Wrapf(err, "get sequencer %s", id.String())
		}
		return infra.Sequencer{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: sequencerKey, Index: id.String()},
			"get sequencer %s",
			id.String(),
		)
	}

	var seq infra.Sequencer
	if err := seq.Unmarshal([]byte(val)); err != nil {
		return infra.Sequencer{}, errors.Wrapf(err, "get sequencer %s", id.String())
	}
	return seq, nil
}

// SetSequencer redis implementation.
func (s *Store) SetSequencer(seq infra.Sequencer) error {
	raw, err := seq.Marshal()
	if err != nil {
		return errors.Wrapf(err, "set sequencer %s", seq.ID.String())
	}
	return errors.Wrapf(
		s.Set(sequencerKey+seq.ID.String(), raw, 0).Err(),
		"set sequencer %s",
		seq.ID.String(),
	)
}

// DelSequencer deletes sequencer in redis.
func (s *Store) DelSequencer(id ulid.ID) error {
	return errors.Wrapf(s.Del(sequencerKey+id.String()).Err(), "del sequencer %s", id.String())
}
