package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01"
)

// EventMapper mocks game.EventMapper.
type EventMapper struct {
	SetEventFunc   func(game.Event, game.ID) error
	SetEventCount  int32
	ListEventFunc  func(game.EventSubset) ([]game.Event, error)
	ListEventCount int32
}

// SetEvent mocks game.EventMapper.
func (s *EventMapper) SetEvent(event game.Event, id game.ID) error {
	atomic.AddInt32(&s.SetEventCount, 1)
	if s.SetEventFunc == nil {
		return nil
	}
	return s.SetEventFunc(event, id)
}

// ListEvent mocks game.EventMapper.
func (s *EventMapper) ListEvent(subset game.EventSubset) ([]game.Event, error) {
	atomic.AddInt32(&s.ListEventCount, 1)
	if s.ListEventFunc == nil {
		return nil, nil
	}
	return s.ListEventFunc(subset)
}

// NewEventMapper returns a event service mock ready for usage.
func NewEventMapper() *EventMapper {
	return &EventMapper{}
}
