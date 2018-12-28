package main

import (
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

func (a *app) LootSource(id ulid.ID, e event.E) error {

	loot := e.Action.GetValue().(*event.LootSource)
	ts := e.ID.Time()

	// #Retrieve entity
	source, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrapf(err, "get entity %s", id.String())
	}

	// #Retrieve target entity
	target, err := a.EntityStore.GetEntity(loot.TargetID, ts)
	if err != nil {
		return errors.Wrapf(err, "get entity %s", loot.TargetID.String())
	}

	// #Retrieve source inventory
	sourceInventory, err := a.EntityInventoryStore.GetInventory(source.InventoryID)
	if err != nil {
		return errors.Wrapf(err, "retrieve inventory %s from source %s", source.InventoryID.String(), target.ID.String())
	}
	if len(sourceInventory.Items) > int(sourceInventory.Size_-1) {
		return gerrors.ErrFullInventory
	}

	// #Check distance between source and target
	dist, err := a.SectorService.Segment(source.Position, target.Position)
	if err != nil {
		return errors.Wrapf(err, "calculate segment between entity %s and target %s", source.ID.String(), target.ID.String())
	}
	if dist > a.lootRadius {
		return gerrors.ErrOutOfRange
	}

	// #Publish loot event to target.
	e = event.E{
		ID: ulid.NewTimeID(ts + 1),
		Action: event.Action{
			LootTarget: &event.LootTarget{
				SourceID: id,
				ItemID:   loot.ItemID,
			},
		},
		Trigger: e.ID,
	}
	if err := a.EventQStore.PublishEvent(e, target.ID); err != nil {
		return errors.Wrapf(err, "publish move target event %s to target %s", e.ID.String(), target.String())
	}

	return nil
}

func (a *app) LootTarget(id ulid.ID, e event.E) error {

	loot := e.Action.GetValue().(*event.LootTarget)
	ts := e.ID.Time()

	// #Retrieve entity
	target, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrapf(err, "get entity %s", id.String())
	}

	// #Retrieve target inventory
	targetInventory, err := a.EntityInventoryStore.GetInventory(target.InventoryID)
	if err != nil {
		return errors.Wrapf(err, "retrieve inventory %s from target %s", target.InventoryID.String(), target.ID.String())
	}

	// #Check item exists in inventory
	n, ok := targetInventory.Items[loot.ItemID.String()]
	if !ok || n < 1 {
		return errors.Wrapf(gerrors.ErrMissingItem, "retrieve item %s from inventory %s", loot.ItemID.String(), target.ID.String())
	}

	// #Remove item from inventory
	if n == 1 {
		delete(targetInventory.Items, loot.ItemID.String())
	} else {
		targetInventory.Items[loot.ItemID.String()]--
	}

	// Set new inventory
	if err := a.EntityInventoryStore.SetInventory(targetInventory); err != nil {
		return errors.Wrapf(err, "set inventory %s from target %s", targetInventory.ID.String(), target.ID.String())
	}

	// #Publish loot event to target.
	e = event.E{
		ID: ulid.NewTimeID(ts + 1),
		Action: event.Action{
			LootFeedback: &event.LootFeedback{
				SourceID: id,
				ItemID:   loot.ItemID,
			},
		},
	}
	if err := a.EventQStore.PublishEvent(e, target.ID); err != nil {
		return errors.Wrapf(err, "publish move target event %s to target %s", e.String(), target.String())
	}

	return nil
}

func (a *app) LootFeedback(id ulid.ID, e event.E) error {

	loot := e.Action.GetValue().(*event.LootTarget)
	ts := e.ID.Time()

	// #Retrieve entity
	source, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrapf(err, "get entity %s", id.String())
	}

	// #Retrieve source inventory
	sourceInventory, err := a.EntityInventoryStore.GetInventory(source.InventoryID)
	if err != nil {
		return errors.Wrapf(err, "retrieve inventory %s from source %s", source.InventoryID.String(), source.ID.String())
	}

	// #Check item exists in inventory
	_, ok := sourceInventory.Items[loot.ItemID.String()]
	if !ok {
		sourceInventory.Items[loot.ItemID.String()] = 1
	} else {
		sourceInventory.Items[loot.ItemID.String()]++
	}

	// Set new inventory
	if err := a.EntityInventoryStore.SetInventory(sourceInventory); err != nil {
		return errors.Wrapf(err, "set inventory %s from source %s", sourceInventory.ID.String(), source.ID.String())
	}

	return nil
}
