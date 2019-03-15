package main

import (
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/item"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// ConsumeSource checks if item exists and is consumable and send consume target.
func (a *app) ConsumeSource(id gulid.ID, e event.E) error {

	cs := e.Action.ConsumeSource
	ts := e.ID.Time()

	// #Check permission source/token
	if err := a.EntityPermissionService.CheckPermission(e.Token, id); err != nil {
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

	// #Check item exists in inventory
	n, ok := sourceInventory.Items[cs.ItemID.String()]
	if !ok || n < 1 {
		return errors.Wrap(gerrors.ErrMissingItem{
			ItemID:      cs.ItemID.String(),
			InventoryID: sourceInventory.ID.String(),
		}, "consume source")
	}

	// #Retrieve item
	it, err := a.ItemStore.GetItem(cs.ItemID)
	if err != nil {
		return errors.Wrap(err, "consume source")
	}

	// #Check item is consumable
	if !it.IsConsumable() {
		return errors.Wrap(gerrors.ErrNotConsumableItem{ItemID: it.ID.String()}, "consume source")
	}

	return errors.Wrap(a.EventQStore.PublishEvent(
		event.E{
			ID: gulid.NewTimeID(ts + 1),
			Action: event.Action{
				ConsumeTarget: &event.ConsumeTarget{
					Source: source,
					ItemID: cs.ItemID,
				},
			},
			Trigger: e.ID,
		}, cs.TargetID),
		"consume source",
	)
}

// ConsumeTarget checks distance and apply an item to a target from source.
func (a *app) ConsumeTarget(id gulid.ID, e event.E) error {

	ct := e.Action.ConsumeTarget
	ts := e.ID.Time()

	// #Retrieve target
	target, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrap(err, "consume target")
	}

	// #Check distance between source and target
	dist, err := a.SectorService.Segment(ct.Source.Position, target.Position)
	if err != nil {
		return errors.Wrap(err, "consume target")
	}
	if dist > a.lootRadius {
		return errors.Wrap(
			gerrors.ErrOutOfRange{
				Dist:  dist,
				Range: a.consumeRadius,
			},
			"consume target",
		)
	}

	// #Retrieve item
	it, err := a.ItemStore.GetItem(ct.ItemID)
	if err != nil {
		return errors.Wrap(err, "consume target")
	}

	// #Consume action item I on entity E.
	switch v := it.Type.GetValue().(type) {
	case *item.Orb:
		ab, err := a.AbilityTemplateStore.GetTemplate(v.AbilityID)
		if err != nil {
			return errors.Wrap(err, "consume target")
		}
		if err := a.AbilityStore.SetAbility(ab, target.ID); err != nil {
			return errors.Wrap(err, "consume target")
		}
	default:
		return errors.Wrap(gerrors.ErrNotImplementedYet{
			Version: "0.2.0",
		}, "consume target")
	}

	return errors.Wrap(a.EventQStore.PublishEvent(
		event.E{
			ID: gulid.NewTimeID(ts + 1),
			Action: event.Action{
				ConsumeFeedback: &event.ConsumeFeedback{
					TargetID: target.ID,
					ItemID:   ct.ItemID,
				},
			},
			Trigger: e.ID,
		}, ct.Source.ID),
		"consume target",
	)
}

// ConsumeFeedback removes item from source inventory.
func (a *app) ConsumeFeedback(id gulid.ID, e event.E) error {

	cf := e.Action.ConsumeFeedback
	ts := e.ID.Time()

	// #Retrieve entity
	source, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrap(err, "consume feedback")
	}

	// #Retrieve source inventory
	sourceInventory, err := a.EntityInventoryStore.GetInventory(source.InventoryID)
	if err != nil {
		return errors.Wrap(err, "consume feedback")
	}

	// #Check item exists in inventory
	n, ok := sourceInventory.Items[cf.ItemID.String()]
	if !ok || n < 1 {
		return errors.Wrap(gerrors.ErrMissingItem{
			ItemID:      cf.ItemID.String(),
			InventoryID: sourceInventory.ID.String(),
		}, "consume feedback")
	}

	// #Remove item from inventory
	if n == 1 {
		delete(sourceInventory.Items, cf.ItemID.String())
	} else {
		sourceInventory.Items[cf.ItemID.String()]--
	}

	// #Create new inventory
	sourceInventory.ID = gulid.NewID()
	if err := a.EntityInventoryStore.SetInventory(sourceInventory); err != nil {
		return errors.Wrap(err, "consume feedback")
	}

	// Set new inventory
	source.InventoryID = sourceInventory.ID
	return errors.Wrap(a.EntityStore.SetEntity(source, ts), "consume feedback")
}
