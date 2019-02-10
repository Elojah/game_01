package main

import (
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

func (a *app) ConsumeSource(id ulid.ID, e event.E) error {

	ls := e.Action.ConsumeSource
	ts := e.ID.Time()

	// #Check permission source/token
	if err := a.EntityPermissionService.CheckSource(id, e.Token); err != nil {
		return errors.Wrap(err, "consume source")
	}

	// #Retrieve entity
	source, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrap(err, "consume source")
	}

	// #Retrieve source inventory
	sourceInventory, err := a.EntityInventoryStore.GetInventory(source.InventoryID)
	if err != nil {
		return errors.Wrap(err, "consume source")
	}

	// TODO check object exists in inventory

	// #Publish loot event to target.
	return errors.Wrap(a.EventQStore.PublishEvent(
		event.E{
			ID: ulid.NewTimeID(ts + 1),
			Action: event.Action{
				ConsumeTarget: &event.ConsumeTarget{
					SourceID: id,
					ItemID:   ls.ItemID,
				},
			},
			Trigger: e.ID,
		}, targetID),
		"consume source",
	)
}

func (a *app) ConsumeTarget(id ulid.ID, e event.E) error {

	lt := e.Action.ConsumeTarget
	ts := e.ID.Time()

	// #Retrieve entity
	target, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrap(err, "retrieve entity")
	}

	// #Retrieve target inventory
	targetInventory, err := a.EntityInventoryStore.GetInventory(target.InventoryID)
	if err != nil {
		return errors.Wrap(err, "retrieve inventory")
	}

	// #Check item exists in inventory
	n, ok := targetInventory.Items[lt.ItemID.String()]
	if !ok || n < 1 {
		return errors.Wrap(
			gerrors.ErrMissingItem{
				ItemID:      lt.ItemID.String(),
				InventoryID: target.ID.String(),
			},
			"check loot validity",
		)
	}

	// #Remove item from inventory
	if n == 1 {
		delete(targetInventory.Items, lt.ItemID.String())
	} else {
		targetInventory.Items[lt.ItemID.String()]--
	}

	// Set new inventory
	if err := a.EntityInventoryStore.SetInventory(targetInventory); err != nil {
		return errors.Wrap(err, "validate loot")
	}

	// #Publish loot event to target.
	return errors.Wrap(a.EventQStore.PublishEvent(
		event.E{
			ID: ulid.NewTimeID(ts + 1),
			Action: event.Action{
				ConsumeFeedback: &event.ConsumeFeedback{
					SourceID: id,
					ItemID:   lt.ItemID,
				},
			},
		}, target.ID),
		"validate loot",
	)
}

func (a *app) ConsumeFeedback(id ulid.ID, e event.E) error {

	lf := e.Action.ConsumeFeedback
	ts := e.ID.Time()

	// #Retrieve entity
	source, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrap(err, "retrieve entity")
	}

	// #Retrieve source inventory
	sourceInventory, err := a.EntityInventoryStore.GetInventory(source.InventoryID)
	if err != nil {
		return errors.Wrap(err, "retrieve inventory")
	}

	// #Check item exists in inventory
	_, ok := sourceInventory.Items[lf.ItemID.String()]
	if !ok {
		sourceInventory.Items[lf.ItemID.String()] = 1
	} else {
		sourceInventory.Items[lf.ItemID.String()]++
	}

	// Set new inventory
	return errors.Wrap(a.EntityInventoryStore.SetInventory(sourceInventory), "validate loot")
}
