package srg

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	feedbackKey = "afb:"
)

// GetFeedback implemented with redis.
func (s *Store) GetFeedback(id ulid.ID) (ability.Feedback, error) {
	val, err := s.Get(feedbackKey + id.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return ability.Feedback{}, err
		}
		return ability.Feedback{}, errors.ErrNotFound
	}

	var fb ability.Feedback
	if err := fb.Unmarshal([]byte(val)); err != nil {
		return ability.Feedback{}, err
	}
	return fb, nil
}

// SetFeedback implemented with redis.
func (s *Store) SetFeedback(fb ability.Feedback) error {
	raw, err := fb.Marshal()
	if err != nil {
		return err
	}
	return s.Set(feedbackKey+fb.ID.String(), raw, 0).Err()
}
