package mocks

import (
	"sync/atomic"

	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// TriggerService is mock for event.TriggerService.
type TriggerService struct {
	SetFunc  func(event.E, gulid.ID) error
	SetCount int32
}

// Set mocks Set method of TriggerService.
func (s *TriggerService) Set(e event.E, id gulid.ID) error {
	atomic.AddInt32(&s.SetCount, 1)
	if s.SetFunc == nil {
		return nil
	}
	return s.SetFunc(e, id)
}

// NewTriggerService returns a TriggerService mock.
func NewTriggerService() *TriggerService {
	return &TriggerService{}
}
