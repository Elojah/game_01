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
		// Cancel event
		if err := s.Cancel(e); err != nil {
			return errors.Wrapf(err, "cancel event %s", e.ID.String())
		}
		// If event is a cancellation, don't set event or trigger and stop here
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

// Cancel cancels an event sending a cancel action to all events triggered by e.
// Works recursively
func (s *TriggerService) Cancel(e event.E) error {
	switch e.Action.GetValue().(type) {
	case *event.MoveSource:
		return s.CancelMoveSource(e)
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
	return gerrors.ErrNotImplementedYet
}

// CancelMoveSource cancels a MoveSource event.
func (s *TriggerService) CancelMoveSource(e event.E) error {
	ms := e.Action.MoveSource
	e = event.E{
		ID: gulid.NewTimeID(e.ID.Time()),
		Action: event.Action{
			Cancel: &event.Cancel{},
		},
		Trigger: e.ID,
	}

	var g errgroup.Group
	for _, targetID := range ms.TargetIDs {
		targetID := targetID
		g.Go(func() error {
			return errors.Wrapf(s.QStore.PublishEvent(e, targetID), "publish cancel move target event %s to target %s", e.String(), targetID.String())
		})
	}
	return g.Wait()
}

// CancelMoveTarget cancels a MoveTarget event.
func (s *TriggerService) CancelMoveTarget(e event.E) error {
	// WARNING Actually it is not a leaf but the tree would resend to all sector entities
	// Or we may should save entities who interacted with this specific spot, think about it
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
			return gerrors.ErrNotImplementedYet
		}
		for _, target := range targets.Entities {
			target := target
			g.Go(func() error {
				return errors.Wrapf(s.QStore.PublishEvent(e, target), "publish cancel move target event %s to target %s", e.String(), target.String())
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
	return errors.Wrapf(s.QStore.PublishEvent(e, pt.Source.ID), "publish cancel perform target event %s to source %s", e.String(), pt.Source.ID.String())
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
	return errors.Wrapf(s.QStore.PublishEvent(e, ls.TargetID), "publish cancel loot source event %s to target %s", e.String(), ls.TargetID.String())
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
	return errors.Wrapf(s.QStore.PublishEvent(e, lt.SourceID), "publish cancel loot target event %s to source %s", e.String(), lt.SourceID.String())
}

// CancelLootFeedback cancels a LootFeedback event.
func (s *TriggerService) CancelLootFeedback(e event.E) error {
	// Leaf
	return nil
}

// CancelConsumeSource cancels a ConsumeSource event.
func (s *TriggerService) CancelConsumeSource(e event.E) error { return gerrors.ErrNotImplementedYet }

// CancelConsumeTarget cancels a ConsumeTarget event.
func (s *TriggerService) CancelConsumeTarget(e event.E) error { return gerrors.ErrNotImplementedYet }
