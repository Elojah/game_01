package svc

import (
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// TriggerService implements set event with trigger interactions
type TriggerService struct {
	event.TriggerStore
	event.Store
	event.QStore
}

// Set event if necessary considering trigger update or removal.
func (s *TriggerService) Set(e event.E, entityID gulid.ID) error {

	if e.Trigger.IsZero() {
		// Set event
		return errors.Wrapf(s.Store.SetEvent(e, entityID), "set event with trigger")
	}

	prevID, err := s.TriggerStore.GetTrigger(e.Trigger, entityID)
	if err != nil {
		switch errors.Cause(err).(type) {
		// Trigger doesn't exist yet so write event+trigger
		case gerrors.ErrNotFound:
			// Cancel events will just be ignored here because it can't cancel an event not triggered yet
			// If event is a cancellation, don't set event or trigger but returns a no calculate error
			if e.Action.Cancel != nil {
				return errors.Wrapf(gerrors.ErrIneffectiveCancel{
					TriggerID: e.Trigger.String(),
				}, "set event with trigger")
			}
			if err := s.Store.SetEvent(e, entityID); err != nil {
				return errors.Wrapf(err, "set event with trigger")
			}
			return errors.Wrapf(s.TriggerStore.SetTrigger(event.Trigger{
				EntityID:      entityID,
				EventSourceID: e.Trigger,
				EventTargetID: e.ID,
			}), "set event with trigger")
		default:
			return errors.Wrapf(err, "set event with trigger")
		}
	}

	// No errors when retrieving trigger means a event has already been triggered by it
	// In this case we clean previous event and previous trigger

	// Retrieve and delete previous event.
	prev, err := s.Store.GetEvent(prevID, entityID)
	if err != nil {
		return errors.Wrapf(err, "set event with trigger")
	}
	if err := s.Store.DelEvent(prevID, entityID); err != nil {
		return errors.Wrapf(err, "set event with trigger")
	}
	// Delete trigger
	if err := s.TriggerStore.DelTrigger(e.Trigger, entityID); err != nil {
		return errors.Wrapf(err, "set event with trigger")
	}
	// Cancel previous event
	if err := s.Cancel(prev); err != nil {
		return errors.Wrapf(err, "set event with trigger")
	}
	// If event is a cancellation, don't set event or trigger and stop here
	if e.Action.Cancel != nil {
		return nil
	}

	// Set event and trigger
	if err := s.Store.SetEvent(e, entityID); err != nil {
		return errors.Wrapf(err, "set event with trigger")
	}
	if err := s.TriggerStore.SetTrigger(event.Trigger{
		EntityID:      entityID,
		EventSourceID: e.Trigger,
		EventTargetID: e.ID,
	}); err != nil {
		return errors.Wrapf(err, "set event with trigger")
	}

	return nil
}

// Cancel cancels an event sending a cancel action to all events triggered by e.
// Works recursively
func (s *TriggerService) Cancel(e event.E) error {
	switch e.Action.GetValue().(type) {
	case *event.MoveTarget:
		return s.CancelMoveTarget(e)
	case *event.CastSource:
		return s.CancelCastSource(e)
	case *event.PerformSource:
		return s.CancelPerformSource(e)
	case *event.PerformTarget:
		return s.CancelPerformTarget(e)
	case *event.FeedbackTarget:
		return s.CancelFeedbackTarget(e)
	case *event.LootSource:
		return s.CancelLootSource(e)
	case *event.LootTarget:
		return s.CancelLootTarget(e)
	case *event.LootFeedback:
		return s.CancelLootFeedback(e)
	case *event.ConsumeSource:
		return s.CancelConsumeSource(e)
	case *event.ConsumeTarget:
		return s.CancelConsumeTarget(e)
	}
	return gerrors.ErrNotImplementedYet{Version: "0.2.0"}
}

// CancelMoveTarget cancels a MoveTarget event.
func (s *TriggerService) CancelMoveTarget(e event.E) error {
	// Actually move interacts with all sector entities so cancel may trigger area abilities for example
	return nil
}

// CancelCastSource cancels a CastSource event.
func (s *TriggerService) CancelCastSource(e event.E) error {
	// Cancel CastSource automatically propagates on itself
	return nil
}

// CancelPerformSource cancels a PerformSource event.
// Cancels perform target if cast source was cancelled.
func (s *TriggerService) CancelPerformSource(e event.E) error {
	ps := e.Action.PerformSource
	e = event.E{
		ID: gulid.NewTimeID(e.ID.Time()),
		Action: event.Action{
			Cancel: &event.Cancel{},
		},
		Trigger: e.ID,
	}

	var g errgroup.Group
	for _, targets := range ps.Targets {
		if len(targets.Positions) != 0 {
			return gerrors.ErrNotImplementedYet{Version: "0.2.0"}
		}
		for _, target := range targets.Entities {
			target := target
			g.Go(func() error {
				return errors.Wrapf(s.QStore.PublishEvent(e, target), "cancel move target")
			})
		}
		if err := g.Wait(); err != nil {
			return err
		}
	}
	return nil
}

// CancelPerformTarget cancels a PerformTarget event.
func (s *TriggerService) CancelPerformTarget(e event.E) error {

	pt := e.Action.PerformTarget
	e = event.E{
		ID: gulid.NewTimeID(e.ID.Time()),
		Action: event.Action{
			Cancel: &event.Cancel{},
		},
		Trigger: e.ID,
	}
	return errors.Wrapf(s.QStore.PublishEvent(e, pt.Source.ID), "cancel perform target")
}

// CancelFeedbackTarget cancels a FeedbackTarget event.
func (s *TriggerService) CancelFeedbackTarget(e event.E) error {
	// Leaf
	return nil
}

// CancelLootSource cancels a LootSource event.
func (s *TriggerService) CancelLootSource(e event.E) error {
	ls := e.Action.LootSource
	e = event.E{
		ID: gulid.NewTimeID(e.ID.Time()),
		Action: event.Action{
			Cancel: &event.Cancel{},
		},
		Trigger: e.ID,
	}
	return errors.Wrapf(s.QStore.PublishEvent(e, ls.TargetID), "cancel loot source")
}

// CancelLootTarget cancels a LootTarget event.
func (s *TriggerService) CancelLootTarget(e event.E) error {
	lt := e.Action.LootTarget
	// #Publish move event to source.
	e = event.E{
		ID: gulid.NewTimeID(e.ID.Time()),
		Action: event.Action{
			Cancel: &event.Cancel{},
		},
		Trigger: e.ID,
	}
	return errors.Wrapf(s.QStore.PublishEvent(e, lt.SourceID), "cancel loot target")
}

// CancelLootFeedback cancels a LootFeedback event.
func (s *TriggerService) CancelLootFeedback(e event.E) error {
	// Leaf
	return nil
}

// CancelConsumeSource cancels a ConsumeSource event.
func (s *TriggerService) CancelConsumeSource(e event.E) error {
	return gerrors.ErrNotImplementedYet{Version: "0.2.0"}
}

// CancelConsumeTarget cancels a ConsumeTarget event.
func (s *TriggerService) CancelConsumeTarget(e event.E) error {
	return gerrors.ErrNotImplementedYet{Version: "0.2.0"}
}
