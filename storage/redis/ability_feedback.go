package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

const (
	afbKey = "afb:"
)

// GetAbilityFeedback implemented with redis.
func (s *Service) GetAbilityFeedback(subset game.AbilityFeedbackSubset) (game.AbilityFeedback, error) {
	val, err := s.Get(afbKey + subset.ID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return game.AbilityFeedback{}, err
		}
		return game.AbilityFeedback{}, storage.ErrNotFound
	}

	var afb storage.AbilityFeedback
	if _, err := afb.Unmarshal([]byte(val)); err != nil {
		return game.AbilityFeedback{}, err
	}
	return afb.Domain(), nil
}

// SetAbilityFeedback implemented with redis.
func (s *Service) SetAbilityFeedback(afb game.AbilityFeedback) error {
	raw, err := storage.NewAbilityFeedback(afb).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(afbKey+afb.ID.String(), raw, 0).Err()
}
