package redis

import (
	"github.com/go-redis/redis"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

const (
	skillTemplateKey = "s_template:"
)

// GetSkillTemplate implemented with redis.
func (s *Service) GetSkillTemplate(subset game.SkillTemplateSubset) (game.SkillTemplate, error) {
	val, err := s.Get(skillTemplateKey + subset.Type.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return game.SkillTemplate{}, err
		}
		return game.SkillTemplate{}, storage.ErrNotFound
	}

	var skill storage.Skill
	if _, err := skill.Unmarshal([]byte(val)); err != nil {
		return game.SkillTemplate{}, err
	}
	return game.SkillTemplate(skill.Domain()), nil
}

// SetSkillTemplate implemented with redis.
func (s *Service) SetSkillTemplate(template game.SkillTemplate) error {
	raw, err := storage.NewSkill(game.Skill(template)).Marshal(nil)
	if err != nil {
		return err
	}
	return s.Set(skillTemplateKey+template.Type.String(), raw, 0).Err()
}
