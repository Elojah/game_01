package app

import (
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

var _ event.App = (*A)(nil)

// A implements events usecases.
type A struct {
	event.QStore
	event.Store
	event.TriggerStore
}

// Create creates an event if necessary considering trigger update or removal.
func (app *A) Create(e event.E, entityID gulid.ID) error {

	if e.Trigger.IsZero() {
		// Upsert event
		return errors.Wrapf(app.Upsert(e, entityID), "create event with trigger")
	}

	prevID, err := app.FetchTrigger(e.Trigger, entityID)
	if err != nil {
		switch errors.Cause(err).(type) {
		// Trigger doesn't exist yet so write event+trigger
		case gerrors.ErrNotFound:
			// Cancel events will just be ignored here because it can't cancel an event not triggered yet
			// If event is a cancellation, don't set event or trigger but returns a no calculate error
			if e.Action.Cancel != nil {
				return errors.Wrapf(gerrors.ErrIneffectiveCancel{
					TriggerID: e.Trigger.String(),
				}, "create event with trigger")
			}
			if err := app.Upsert(e, entityID); err != nil {
				return errors.Wrapf(err, "create event with trigger")
			}
			return errors.Wrapf(app.UpsertTrigger(event.Trigger{
				EntityID:      entityID,
				EventSourceID: e.Trigger,
				EventTargetID: e.ID,
			}), "create event with trigger")
		default:
			return errors.Wrapf(err, "create event with trigger")
		}
	}

	// No errors when retrieving trigger means a event has already been triggered by it
	// In this case we clean previous event and previous trigger

	// Retrieve and delete previous event.
	prev, err := app.Fetch(prevID, entityID)
	if err != nil {
		return errors.Wrapf(err, "create event with trigger")
	}
	if err := app.Remove(prevID, entityID); err != nil {
		return errors.Wrapf(err, "create event with trigger")
	}
	// Delete trigger
	if err := app.RemoveTrigger(e.Trigger, entityID); err != nil {
		return errors.Wrapf(err, "create event with trigger")
	}
	// Cancel previous event
	if err := app.Cancel(prev); err != nil {
		return errors.Wrapf(err, "create event with trigger")
	}
	// If event is a cancellation, don't set event or trigger and stop here
	if e.Action.Cancel != nil {
		return nil
	}

	// Upsert event and trigger
	if err := app.Upsert(e, entityID); err != nil {
		return errors.Wrapf(err, "create event with trigger")
	}
	if err := app.UpsertTrigger(event.Trigger{
		EntityID:      entityID,
		EventSourceID: e.Trigger,
		EventTargetID: e.ID,
	}); err != nil {
		return errors.Wrapf(err, "create event with trigger")
	}

	return nil
}

// Cancel cancels an event sending a cancel action to all events triggered by e.
// Works recursively
func (app *A) Cancel(e event.E) error {
	switch e.Action.GetValue().(type) {
	case *event.MoveTarget:
		return app.CancelMoveTarget(e)
	case *event.CastSource:
		return app.CancelCastSource(e)
	case *event.PerformSource:
		return app.CancelPerformSource(e)
	case *event.PerformTarget:
		return app.CancelPerformTarget(e)
	case *event.PerformFeedback:
		return app.CancelPerformFeedback(e)
	case *event.LootSource:
		return app.CancelLootSource(e)
	case *event.LootTarget:
		return app.CancelLootTarget(e)
	case *event.LootFeedback:
		return app.CancelLootFeedback(e)
	case *event.ConsumeSource:
		return app.CancelConsumeSource(e)
	case *event.ConsumeTarget:
		return app.CancelConsumeTarget(e)
	case *event.ConsumeFeedback:
		return app.CancelConsumeFeedback(e)
	}
	return gerrors.ErrNotImplementedYet{Version: "0.2.0"}
}

// CancelMoveTarget cancels a MoveTarget event.
func (app *A) CancelMoveTarget(e event.E) error {
	// Actually move interacts with all sector entities so cancel may trigger area abilities for example
	return nil
}

// CancelCastSource cancels a CastSource event.
func (app *A) CancelCastSource(e event.E) error {
	// Cancel CastSource automatically propagates on itself
	return nil
}

// CancelPerformSource cancels a PerformSource event.
// Cancels perform target if cast source was cancelled.
func (app *A) CancelPerformSource(e event.E) error {
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
				return errors.Wrapf(app.Publish(e, target), "cancel perform source")
			})
		}
		if err := g.Wait(); err != nil {
			return err
		}
	}
	return nil
}

// CancelPerformTarget cancels a PerformTarget event.
func (app *A) CancelPerformTarget(e event.E) error {

	pt := e.Action.PerformTarget
	e = event.E{
		ID: gulid.NewTimeID(e.ID.Time()),
		Action: event.Action{
			Cancel: &event.Cancel{},
		},
		Trigger: e.ID,
	}
	return errors.Wrapf(app.Publish(e, pt.Source.ID), "cancel perform target")
}

// CancelPerformFeedback cancels a PerformFeedback event.
func (app *A) CancelPerformFeedback(e event.E) error {
	// Leaf
	return nil
}

// CancelLootSource cancels a LootSource event.
func (app *A) CancelLootSource(e event.E) error {
	ls := e.Action.LootSource
	e = event.E{
		ID: gulid.NewTimeID(e.ID.Time()),
		Action: event.Action{
			Cancel: &event.Cancel{},
		},
		Trigger: e.ID,
	}
	return errors.Wrapf(app.Publish(e, ls.TargetID), "cancel loot source")
}

// CancelLootTarget cancels a LootTarget event.
func (app *A) CancelLootTarget(e event.E) error {
	lt := e.Action.LootTarget
	// #Publish move event to source.
	e = event.E{
		ID: gulid.NewTimeID(e.ID.Time()),
		Action: event.Action{
			Cancel: &event.Cancel{},
		},
		Trigger: e.ID,
	}
	return errors.Wrapf(app.Publish(e, lt.Source.ID), "cancel loot target")
}

// CancelLootFeedback cancels a LootFeedback event.
func (app *A) CancelLootFeedback(e event.E) error {
	// Leaf
	return nil
}

// CancelConsumeSource cancels a ConsumeSource event.
func (app *A) CancelConsumeSource(e event.E) error {
	return gerrors.ErrNotImplementedYet{Version: "0.2.0"}
}

// CancelConsumeTarget cancels a ConsumeTarget event.
func (app *A) CancelConsumeTarget(e event.E) error {
	return gerrors.ErrNotImplementedYet{Version: "0.2.0"}
}

// CancelConsumeFeedback cancels a ConsumeTarget event.
func (app *A) CancelConsumeFeedback(e event.E) error {
	return gerrors.ErrNotImplementedYet{Version: "0.2.0"}
}
