package srg

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/errors"
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
func (s *Store) DelSequencer(id ulid.ID) error {
	return s.Del(sequencerKey + id.String()).Err()
}
