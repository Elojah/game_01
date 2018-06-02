package redis

import (
	"github.com/elojah/game_01"
	"github.com/elojah/redis"
)

var _ game.AbilityMapper = (*Service)(nil)
var _ game.AbilityFeedbackMapper = (*Service)(nil)
var _ game.AbilityTemplateMapper = (*Service)(nil)
var _ game.AccountMapper = (*Service)(nil)
var _ game.EntityMapper = (*Service)(nil)
var _ game.EntityTemplateMapper = (*Service)(nil)
var _ game.EventMapper = (*Service)(nil)
var _ game.PCMapper = (*Service)(nil)
var _ game.PermissionMapper = (*Service)(nil)
var _ game.TokenMapper = (*Service)(nil)
var _ game.SectorMapper = (*Service)(nil)

// Service implements token and entity.
type Service struct {
	*redis.Service
}

// NewService returns a new game_01 redis Service.
func NewService(s *redis.Service) *Service {
	return &Service{
		Service: s,
	}
}
