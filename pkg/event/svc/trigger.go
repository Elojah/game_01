package svc

import (
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// TriggerService implements set event with trigger interactions
type TriggerService struct {
	event.TriggerStore
	event.Store
}

// Set event if necessary considering trigger update or removal.
func (s *TriggerService) Set(e event.E, entityID gulid.ID) error {
	var t gulid.ID
	err := gerrors.ErrNotFound

	// This trick is to avoid huge if clauses. By default err is ErrNotFound to jump to final set
	if !e.Trigger.IsZero() {
		t, err = s.TriggerStore.GetTrigger(e.Trigger, entityID)
	}

	// err checking of above statement
	if err != nil && err != gerrors.ErrNotFound {
		return errors.Wrapf(err, "get trigger %s from event %s", e.Trigger.String(), e.ID.String())
	}
	// No errors when retrieving trigger means a event has already been triggered by it
	// In this case we clean previous event and previous trigger
	if err == nil {
		// Delete previous event.
		if err := s.Store.DelEvent(t, entityID); err != nil {
			return errors.Wrapf(err, "delete previous event %s", e.ID.String())
		}
		// Delete trigger
		if err := s.TriggerStore.DelTrigger(e.Trigger, entityID); err != nil {
			return errors.Wrapf(err, "delete trigger %s", e.Trigger.String())
		}
		// If event is a cancellation, don't set event or trigger
		if e.Action.Cancel != nil {
			return nil
		}
	}
	// If event is a cancellation, don't set event or trigger but returns a no calculate error
	if e.Action.Cancel != nil {
		return gerrors.ErrIneffectiveCancel
	}

	// Set event and trigger
	if err := s.Store.SetEvent(e, entityID); err != nil {
		return errors.Wrapf(err, "create event %s", e.ID.String())
	}
	if err := s.TriggerStore.SetTrigger(event.Trigger{
		EntityID:      entityID,
		EventSourceID: e.Trigger,
		EventTargetID: e.ID,
	}); err != nil {
		return errors.Wrapf(err, "create trigger %s", e.ID.String())
	}

	return nil
}
