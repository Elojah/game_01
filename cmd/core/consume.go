package main

import (
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/item"
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

	// #Check item exists in inventory
	if n, ok := sourceInventory.Items[ls.ItemID.String()]; !ok || n == 0 {
		return errors.Wrap(gerrors.ErrMissingItem{
			ItemID:      ls.ItemID.String(),
			InventoryID: sourceInventory.ID.String(),
		}, "consume source")
	}

	// #Retrieve item
	it, err := a.ItemStore.GetItem(ls.ItemID)
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
		return a.AbilityStore.SetAbility(ab, source.ID)
	}

	return errors.Wrap(gerrors.ErrNotImplementedYet{
		Version: "0.2.0",
	}, "consume source")
}
