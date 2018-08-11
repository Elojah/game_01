package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

// EventService mocks event.Service.
type EventService struct {
	SetEventFunc   func(event.E, ulid.ID) error
	SetEventCount  int32
	ListEventFunc  func(event.Subset) ([]event.E, error)
	ListEventCount int32
}

// SetEvent mocks event.Service.
func (s *EventService) SetEvent(e event.E, id ulid.ID) error {
	atomic.AddInt32(&s.SetEventCount, 1)
	if s.SetEventFunc == nil {
		return nil
	}
	return s.SetEventFunc(e, id)
}

// ListEvent mocks event.Service.
func (s *EventService) ListEvent(subset event.Subset) ([]event.E, error) {
	atomic.AddInt32(&s.ListEventCount, 1)
	if s.ListEventFunc == nil {
		return nil, nil
	}
	return s.ListEventFunc(subset)
}

// NewEventService returns a event service mock ready for usage.
func NewEventService() *EventService {
	return &EventService{}
}
