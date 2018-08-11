package redis

import (
	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/redis"
)

var _ ability.Store = (*Service)(nil)
var _ ability.FeedbackStore = (*Service)(nil)
var _ ability.TemplateStore = (*Service)(nil)

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
