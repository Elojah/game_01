package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/storage"
)

const (
	afbKey = "afb:"
)

// GetAbilityFeedback implemented with redis.
func (s *Service) GetAbilityFeedback(subset ability.FeedbackSubset) (ability.Feedback, error) {
	val, err := s.Get(afbKey + subset.ID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return ability.Feedback{}, err
		}
		return ability.Feedback{}, storage.ErrNotFound
	}

	var afb storage.AbilityFeedback
	if _, err := afb.Unmarshal([]byte(val)); err != nil {
		return ability.Feedback{}, err
	}
	return afb.Domain(), nil
}

// SetAbilityFeedback implemented with redis.
func (s *Service) SetAbilityFeedback(afb ability.Feedback) error {
	raw, err := storage.NewAbilityFeedback(afb).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(afbKey+afb.ID.String(), raw, 0).Err()
}
