package main

import (
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

func (a *app) LootSource(id gulid.ID, e event.E) error {

	ls := e.Action.LootSource
	ts := e.ID.Time()

	// #Check permission source/token
	if err := a.EntityPermissionService.CheckPermission(e.Token, id); err != nil {
		return errors.Wrap(err, "loot source")
	}

	// #Retrieve entity
	source, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrap(err, "loot source")
	}

	// #Check if entity is alive
	if source.Dead {
		return errors.Wrap(gerrors.ErrIsDead{EntityID: id.String()}, "loot source")
	}

	// #Check if target is lootable
	if ok, err := a.ItemLootStore.GetLoot(ls.TargetID); !ok || err != nil {
		if err != nil {
			return errors.Wrap(err, "loot source")
		}
		return errors.Wrap(gerrors.ErrNotLootableEntity{EntityID: ls.TargetID.String()}, "loot source")
	}

	// #Retrieve source inventory to check if has a free spot
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

	// #Publish loot event to target.
	return errors.Wrap(a.EventQStore.PublishEvent(
		event.E{
			ID: gulid.NewTimeID(ts + 1),
			Action: event.Action{
				LootTarget: &event.LootTarget{
					Source: source,
					ItemID: ls.ItemID,
				},
			},
			Trigger: e.ID,
		}, ls.TargetID),
		"loot source",
	)
}

func (a *app) LootTarget(id gulid.ID, e event.E) error {

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
		return errors.Wrap(gerrors.ErrMissingItem{
			ItemID:      lt.ItemID.String(),
			InventoryID: target.ID.String(),
		}, "loot target")
	}

	// #Check distance between source and target
	dist, err := a.SectorService.Segment(lt.Source.Position, target.Position)
	if err != nil {
		return errors.Wrap(err, "loot target")
	}
	if dist > a.lootRadius {
		return errors.Wrap(
			gerrors.ErrOutOfRange{
				Dist:  dist,
				Range: a.lootRadius,
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

	// Create new inventory
	targetInventory.ID = gulid.NewID()
	if err := a.EntityInventoryStore.SetInventory(targetInventory); err != nil {
		return errors.Wrap(err, "loot target")
	}

	target.InventoryID = targetInventory.ID
	// Set new inventory to target
	if err := a.EntityStore.SetEntity(target, ts); err != nil {
		return errors.Wrap(err, "loot target")
	}

	// #Publish loot event to target.
	return errors.Wrap(a.EventQStore.PublishEvent(
		event.E{
			ID: gulid.NewTimeID(ts + 1),
			Action: event.Action{
				LootFeedback: &event.LootFeedback{
					TargetID: target.ID,
					ItemID:   lt.ItemID,
				},
			},
			Trigger: e.ID,
		}, lt.Source.ID),
		"loot target",
	)
}

func (a *app) LootFeedback(id gulid.ID, e event.E) error {

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

	// #Add item in inventory
	if sourceInventory.Items == nil {
		sourceInventory.Items = map[string]uint64{
			lf.ItemID.String(): 1,
		}
	} else if _, ok := sourceInventory.Items[lf.ItemID.String()]; !ok {
		sourceInventory.Items[lf.ItemID.String()] = 1
	} else {
		sourceInventory.Items[lf.ItemID.String()]++
	}

	// #Create new inventory
	sourceInventory.ID = gulid.NewID()
	if err := a.EntityInventoryStore.SetInventory(sourceInventory); err != nil {
		return errors.Wrap(err, "loot feedback")
	}

	// Set new inventory
	source.InventoryID = sourceInventory.ID
	return errors.Wrap(a.EntityStore.SetEntity(source, ts), "loot feedback")
}
