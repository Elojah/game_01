package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/storage"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	afbKey = "afb:"
)

// GetAbilityFeedback implemented with redis.
func (s *Service) GetAbilityFeedback(subset ability.FeedbackSubset) (ability.Feedback, error) {
	val, err := s.Get(afbKey + ulid.String(subset.ID)).Result()
	if err != nil {
		if err != redis.Nil {
			return ability.Feedback{}, err
		}
		return ability.Feedback{}, storage.ErrNotFound
	}

	var afb ability.Feedback
	if _, err := afb.Unmarshal([]byte(val)); err != nil {
		return ability.Feedback{}, err
	}
	return afb, nil
}

// SetAbilityFeedback implemented with redis.
func (s *Service) SetAbilityFeedback(afb ability.Feedback) error {
	raw, err := afb.Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(afbKey+ulid.String(afb.ID), raw, 0).Err()
}
