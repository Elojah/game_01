package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// TriggerService is mock for event.TriggerService.
type TriggerService struct {
	SetFunc     func(event.E, gulid.ID) error
	SetCount    int32
	CancelFunc  func(event.E) error
	CancelCount int32
}

// Set mocks Set method of TriggerService.
func (s *TriggerService) Set(e event.E, id gulid.ID) error {
	atomic.AddInt32(&s.SetCount, 1)
	if s.SetFunc == nil {
		return nil
	}
	return s.SetFunc(e, id)
}

// Cancel mocks Cancel method of TriggerService.
func (s *TriggerService) Cancel(e event.E) error {
	atomic.AddInt32(&s.CancelCount, 1)
	if s.CancelFunc == nil {
		return nil
	}
	return s.CancelFunc(e)
}

// NewTriggerService returns a TriggerService mock.
func NewTriggerService() *TriggerService {
	return &TriggerService{}
}
