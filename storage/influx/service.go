package influx

import (
	"github.com/elojah/influx"
)

// Service implements InfluxDB service.
type Service struct {
	*influx.Service
}

// NewService returns a new Service.
func NewService(s *influx.Service) *Service {
	return &Service{
		Service: s,
	}
}
