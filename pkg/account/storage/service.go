package storage

import (
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/redis"
)

var _ account.Store = (*Service)(nil)
var _ account.TokenStore = (*Service)(nil)
var _ account.TokenHCStore = (*Service)(nil)

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
