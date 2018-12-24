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

	// #Check distance between source and target
	dist, err := a.SectorService.Segment(source.Position, target.Position)
	if err != nil {
		return errors.Wrapf(err, "calculate segment between entity %s and target %s", source.ID.String(), target.ID.String())
	}
	if dist > a.lootRadius {
		return gerrors.ErrOutOfRange
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

	// #Publish loot event to target.
	e = event.E{
		ID: ulid.NewTimeID(ts + 1),
		Action: event.Action{
			LootTarget: &event.LootTarget{
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
