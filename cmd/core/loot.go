package main

import (
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

func (a *app) LootSource(id ulid.ID, e event.E) error {

	ls := e.Action.LootSource
	ts := e.ID.Time()

	// #Check permission source/token
	if err := a.EntityPermissionService.CheckSource(id, e.Token); err != nil {
		return errors.Wrap(err, "loot source")
	}

	// #Retrieve entity
	source, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrap(err, "loot source")
	}

	// #Retrieve target entity
	target, err := a.EntityStore.GetEntity(ls.TargetID, ts)
	if err != nil {
		return errors.Wrap(err, "loot source")
	}

	// #Check if target is lootable
	if ok, err := a.ItemLootStore.GetLoot(target.ID); !ok || err != nil {
		if err != nil {
			return errors.Wrap(err, "loot source")
		}
		return errors.Wrap(gerrors.ErrNotLootableEntity{EntityID: target.ID.String()}, "loot source")
	}

	// #Retrieve source inventory
	sourceInventory, err := a.EntityInventoryStore.GetInventory(source.InventoryID)
	if err != nil {
		return errors.Wrap(err, "loot source")
	}
	if len(sourceInventory.Items) > int(sourceInventory.Size_-1) {
		return errors.Wrap(
			gerrors.ErrFullInventory{
				InventoryID: source.InventoryID.String(),
			},
			"loot source",
		)
	}

	// #Check distance between source and target
	dist, err := a.SectorService.Segment(source.Position, target.Position)
	if err != nil {
		return errors.Wrap(err, "loot source")
	}
	if dist > a.lootRadius {
		return errors.Wrap(
			gerrors.ErrOutOfRange{
				Dist:  dist,
				Range: a.lootRadius,
			},
			"loot source",
		)
	}

	// #Publish loot event to target.
	return errors.Wrap(a.EventQStore.PublishEvent(
		event.E{
			ID: ulid.NewTimeID(ts + 1),
			Action: event.Action{
				LootTarget: &event.LootTarget{
					SourceID: id,
					ItemID:   ls.ItemID,
				},
			},
			Trigger: e.ID,
		}, target.ID),
		"loot source",
	)
}

func (a *app) LootTarget(id ulid.ID, e event.E) error {

	lt := e.Action.LootTarget
	ts := e.ID.Time()

	// #Retrieve entity
	target, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrap(err, "loot target")
	}

	// #Retrieve target inventory
	targetInventory, err := a.EntityInventoryStore.GetInventory(target.InventoryID)
	if err != nil {
		return errors.Wrap(err, "loot target")
	}

	// #Check item exists in inventory
	n, ok := targetInventory.Items[lt.ItemID.String()]
	if !ok || n < 1 {
		return errors.Wrap(
			gerrors.ErrMissingItem{
				ItemID:      lt.ItemID.String(),
				InventoryID: target.ID.String(),
			},
			"loot target",
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
		return errors.Wrap(err, "loot target")
	}

	// #Publish loot event to target.
	return errors.Wrap(a.EventQStore.PublishEvent(
		event.E{
			ID: ulid.NewTimeID(ts + 1),
			Action: event.Action{
				LootFeedback: &event.LootFeedback{
					SourceID: id,
					ItemID:   lt.ItemID,
				},
			},
		}, target.ID),
		"loot target",
	)
}

func (a *app) LootFeedback(id ulid.ID, e event.E) error {

	lf := e.Action.LootFeedback
	ts := e.ID.Time()

	// #Retrieve entity
	source, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrap(err, "loot feedback")
	}

	// #Retrieve source inventory
	sourceInventory, err := a.EntityInventoryStore.GetInventory(source.InventoryID)
	if err != nil {
		return errors.Wrap(err, "loot feedback")
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
