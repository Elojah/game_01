package redis

import (
	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/perm"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/redis"
)

var _ ability.Mapper = (*Service)(nil)
var _ ability.FeedbackMapper = (*Service)(nil)
var _ ability.TemplateMapper = (*Service)(nil)
var _ account.Mapper = (*Service)(nil)
var _ entity.Mapper = (*Service)(nil)
var _ entity.TemplateMapper = (*Service)(nil)
var _ event.Mapper = (*Service)(nil)
var _ entity.PCMapper = (*Service)(nil)
var _ perm.Mapper = (*Service)(nil)
var _ sector.Mapper = (*Service)(nil)
var _ sector.EntitiesMapper = (*Service)(nil)
var _ account.TokenMapper = (*Service)(nil)
var _ account.TokenHCMapper = (*Service)(nil)

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
