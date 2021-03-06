package main

import (
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/item"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// ConsumeSource checks if item exists and is consumable and send consume target.
func (svc *service) ConsumeSource(id gulid.ID, e event.E) error {

	cs := e.Action.ConsumeSource
	ts := e.ID.Time()

	// #Check permission source/token
	if err := svc.entity.CheckPermission(e.Token, id); err != nil {
		return errors.Wrap(err, "consume source")
	}

	// #Retrieve entity
	source, err := svc.entity.Fetch(id, ts)
	if err != nil {
		return errors.Wrap(err, "consume source")
	}

	// #Check if entity is alive
	if source.Dead {
		return errors.Wrap(gerrors.ErrIsDead{EntityID: id.String()}, "consume source")
	}

	// #Retrieve source inventory
	sourceInventory, err := svc.entity.FetchMRInventoryFromCache(source.InventoryID, source.ID)
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
	it, err := svc.item.Fetch(cs.ItemID)
	if err != nil {
		return errors.Wrap(err, "consume source")
	}

	// #Check item is consumable
	if !it.IsConsumable() {
		return errors.Wrap(gerrors.ErrNotConsumableItem{ItemID: it.ID.String()}, "consume source")
	}

	return errors.Wrap(svc.event.Publish(
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

// ConsumeTarget checks distance and apply an item to svc target from source.
func (svc *service) ConsumeTarget(id gulid.ID, e event.E) error {

	ct := e.Action.ConsumeTarget
	ts := e.ID.Time()

	// #Retrieve target
	target, err := svc.entity.Fetch(id, ts)
	if err != nil {
		return errors.Wrap(err, "consume target")
	}

	// #Check distance between source and target
	dist, err := svc.sector.Segment(ct.Source.Position, target.Position)
	if err != nil {
		return errors.Wrap(err, "consume target")
	}
	if dist > svc.lootRadius {
		return errors.Wrap(
			gerrors.ErrOutOfRange{
				Dist:  dist,
				Range: svc.consumeRadius,
			},
			"consume target",
		)
	}

	// #Retrieve item
	it, err := svc.item.Fetch(ct.ItemID)
	if err != nil {
		return errors.Wrap(err, "consume target")
	}

	// #Consume action item I on entity E.
	switch v := it.Type.GetValue().(type) {
	case *item.Orb:
		ab, err := svc.ability.FetchTemplate(v.AbilityID)
		if err != nil {
			return errors.Wrap(err, "consume target")
		}
		if err := svc.ability.Upsert(ab, target.ID); err != nil {
			return errors.Wrap(err, "consume target")
		}
	default:
		return errors.Wrap(gerrors.ErrNotImplementedYet{
			Version: "0.2.0",
		}, "consume target")
	}

	return errors.Wrap(svc.event.Publish(
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
func (svc *service) ConsumeFeedback(id gulid.ID, e event.E) error {

	cf := e.Action.ConsumeFeedback
	ts := e.ID.Time()

	// #Retrieve entity
	source, err := svc.entity.Fetch(id, ts)
	if err != nil {
		return errors.Wrap(err, "consume feedback")
	}

	// #Retrieve source inventory
	sourceInventory, err := svc.entity.FetchMRInventoryFromCache(source.InventoryID, source.ID)
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
	if err := svc.entity.UpsertMRInventoryWithCache(source.ID, sourceInventory); err != nil {
		return errors.Wrap(err, "consume feedback")
	}

	// Set new inventory
	source.InventoryID = sourceInventory.ID
	return errors.Wrap(svc.entity.Upsert(source, ts+1), "consume feedback")
}
