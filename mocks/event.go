package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01"
)

// EventService mocks game.EventService.
type EventService struct {
	CreateEventFunc  func(game.Event, game.ID) error
	CreateEventCount int32
	ListEventFunc    func(game.EventSubset) ([]game.Event, error)
	ListEventCount   int32
}

// CreateEvent mocks game.EventService.
func (s *EventService) CreateEvent(event game.Event, id game.ID) error {
	atomic.AddInt32(&s.CreateEventCount, 1)
	if s.CreateEventFunc == nil {
		return nil
	}
	return s.CreateEventFunc(event, id)
}

// ListEvent mocks game.EventService.
func (s *EventService) ListEvent(subset game.EventSubset) ([]game.Event, error) {
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
