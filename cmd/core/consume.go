package main

import (
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/item"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

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

	// #Consume action item I on entity E.
	switch v := it.Type.GetValue().(type) {
	case item.Orb:
		ab, err := a.AbilityTemplateStore.GetTemplate(v.AbilityID)
		if err != nil {
			return errors.Wrap(err, "consume source")
		}
		if err := a.AbilityStore.SetAbility(ab, source.ID); err != nil {
			return errors.Wrap(err, "consume source")
		}
	default:
		return errors.Wrap(gerrors.ErrNotImplementedYet{
			Version: "0.2.0",
		}, "consume source")
	}

	// #Remove item from inventory
	if n == 1 {
		delete(sourceInventory.Items, cs.ItemID.String())
	} else {
		sourceInventory.Items[cs.ItemID.String()]--
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
