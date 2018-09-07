package srg

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/infra"
)

const (
	sequencerKey = "sequencer:"
)

// GetSequencer redis implementation.
func (s *Store) GetSequencer(subset infra.SequencerSubset) (infra.Sequencer, error) {
	val, err := s.Get(sequencerKey + subset.ID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return infra.Sequencer{}, err
		}
		return infra.Sequencer{}, errors.ErrNotFound
	}

	var sequencer infra.Sequencer
	if err := sequencer.Unmarshal([]byte(val)); err != nil {
		return infra.Sequencer{}, err
	}
	return sequencer, nil
}

// SetSequencer redis implementation.
func (s *Store) SetSequencer(sequencer infra.Sequencer) error {
	raw, err := sequencer.Marshal()
	if err != nil {
		return err
	}
	return s.Set(sequencerKey+sequencer.ID.String(), raw, 0).Err()
}

// DelSequencer deletes sequencer in redis.
func (s *Store) DelSequencer(subset infra.SequencerSubset) error {
	return s.Del(sequencerKey + subset.ID.String()).Err()
}
