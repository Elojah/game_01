package scyllax

import (
	"github.com/elojah/scylla"
)

// Service implements ScyllaDB service.
type Service struct {
	*scylla.Service
}

// NewService returns a new Service.
func NewService(s *scylla.Service) *Service {
	return &Service{
		Service: s,
	}
}
