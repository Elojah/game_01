package client

import (
	"github.com/elojah/game_01/cmd/integration/loganalyzer"
)

// Service wraps client helpers.
type Service struct {
	*loganalyzer.LA
}

// NewService returns a integration service for client.
func NewService(la *loganalyzer.LA) *Service {
	return &Service{LA: la}
}
