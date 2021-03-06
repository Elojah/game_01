package main

import (
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

func (svc *service) LootSource(id gulid.ID, e event.E) error {

	ls := e.Action.LootSource
	ts := e.ID.Time()

	// #Check permission source/token
	if err := svc.entity.CheckPermission(e.Token, id); err != nil {
		return errors.Wrap(err, "loot source")
	}

	// #Retrieve entity
	source, err := svc.entity.Fetch(id, ts)
	if err != nil {
		return errors.Wrap(err, "loot source")
	}

	// #Check if entity is alive
	if source.Dead {
		return errors.Wrap(gerrors.ErrIsDead{EntityID: id.String()}, "loot source")
	}

	// #Check if target is lootable
	if ok, err := svc.item.FetchLoot(ls.TargetID); !ok || err != nil {
		if err != nil {
			return errors.Wrap(err, "loot source")
		}
		return errors.Wrap(gerrors.ErrNotLootableEntity{EntityID: ls.TargetID.String()}, "loot source")
	}

	// #Retrieve source inventory to check if has svc free spot
	sourceInventory, err := svc.entity.FetchMRInventoryFromCache(source.InventoryID, source.ID)
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
	return errors.Wrap(svc.event.Publish(
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

func (svc *service) LootTarget(id gulid.ID, e event.E) error {

	lt := e.Action.LootTarget
	ts := e.ID.Time()

	// #Retrieve entity
	target, err := svc.entity.Fetch(id, ts)
	if err != nil {
		return errors.Wrap(err, "loot target")
	}

	// #Retrieve target inventory
	targetInventory, err := svc.entity.FetchMRInventoryFromCache(target.InventoryID, target.ID)
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
	dist, err := svc.sector.Segment(lt.Source.Position, target.Position)
	if err != nil {
		return errors.Wrap(err, "loot target")
	}
	if dist > svc.lootRadius {
		return errors.Wrap(
			gerrors.ErrOutOfRange{
				Dist:  dist,
				Range: svc.lootRadius,
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
	if err := svc.entity.UpsertMRInventoryWithCache(target.ID, targetInventory); err != nil {
		return errors.Wrap(err, "loot target")
	}

	target.InventoryID = targetInventory.ID
	// Set new inventory to target
	if err := svc.entity.Upsert(target, ts+1); err != nil {
		return errors.Wrap(err, "loot target")
	}

	// #Publish loot event to target.
	return errors.Wrap(svc.event.Publish(
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

func (svc *service) LootFeedback(id gulid.ID, e event.E) error {

	lf := e.Action.LootFeedback
	ts := e.ID.Time()

	// #Retrieve entity
	source, err := svc.entity.Fetch(id, ts)
	if err != nil {
		return errors.Wrap(err, "loot feedback")
	}

	// #Retrieve source inventory
	sourceInventory, err := svc.entity.FetchMRInventoryFromCache(source.InventoryID, source.ID)
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
	if err := svc.entity.UpsertMRInventoryWithCache(source.ID, sourceInventory); err != nil {
		return errors.Wrap(err, "loot feedback")
	}

	// Set new inventory
	source.InventoryID = sourceInventory.ID
	return errors.Wrap(svc.entity.Upsert(source, ts+1), "loot feedback")
}
