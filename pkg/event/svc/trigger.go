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
	if !e.Trigger.IsZero() {
		t, err = s.TriggerStore.GetTrigger(e.Trigger, entityID)
	}

	if err != nil && err != gerrors.ErrNotFound {
		return errors.Wrapf(err, "get trigger %s from event %s", e.Trigger.String(), e.ID.String())
	}
	if err == nil {
		// Delete previous event.
		if err := s.Store.DelEvent(t, entityID); err != nil {
			return errors.Wrapf(err, "delete previous event %s", e.ID.String())
		}
		if err := s.TriggerStore.DelTrigger(entityID, e.Trigger); err != nil {
			return errors.Wrapf(err, "delete trigger %s", e.Trigger.String())
		}
		if e.Action.Cancel != nil {
			return nil
		}
	}

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
