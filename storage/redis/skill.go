package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

const (
	skillKey = "skill:"
)

// ListSkill implemented with redis.
func (s *Service) ListSkill(subset game.SkillSubset) ([]game.Skill, error) {
	keys, err := s.Keys(skillKey + subset.EntityID.String() + "*").Result()
	if err != nil {
		if err != redis.Nil {
			return nil, err
		}
		return nil, storage.ErrNotFound
	}

	skills := make([]game.Skill, len(keys))
	for i, key := range keys {
		val, err := s.Get(key).Result()
		if err != nil {
			return nil, err
		}

		var skill storage.Skill
		if _, err := skill.Unmarshal([]byte(val)); err != nil {
			return nil, err
		}
		skills[i] = skill.Domain()
	}
	return skills, nil
}

// GetSkill implemented with redis.
func (s *Service) GetSkill(subset game.SkillSubset) (game.Skill, error) {
	val, err := s.Get(skillKey + subset.EntityID.String() + ":" + subset.ID.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return game.Skill{}, err
		}
		return game.Skill{}, storage.ErrNotFound
	}

	var skill storage.Skill
	if _, err := skill.Unmarshal([]byte(val)); err != nil {
		return game.Skill{}, err
	}
	return skill.Domain(), nil
}

// SetSkill implemented with redis.
func (s *Service) SetSkill(skill game.Skill, entity game.ID) error {
	raw, err := storage.NewSkill(skill).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(skillKey+entity.String()+":"+skill.ID.String(), raw, 0).Err()
}
