package app

import (
	"github.com/oklog/ulid"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/sector"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// A implementation of entity applications.
type A struct {
	entity.InventoryStore
	entity.MRInventoryStore
	entity.PCLeftStore
	entity.PCStore
	entity.PermissionStore
	entity.SpawnStore
	entity.Store
	entity.TemplateStore

	AbilityStore ability.Store

	SectorEntitiesStore sector.EntitiesStore

	SequencerService infra.SequencerService
}

// Disconnect disconnects an entity.
func (app A) Disconnect(id gulid.ID) error {

	e, err := app.Store.Fetch(id, ulid.Now())
	if err != nil {
		return errors.Wrap(err, "disconnect entity")
	}

	// #Close entity sequencer
	if err := app.SequencerService.Remove(id); err != nil {
		return errors.Wrap(err, "disconnect entity")
	}

	// #Delete pc entity position
	if err := app.SectorEntitiesStore.RemoveEntityFromSector(id, e.Position.SectorID); err != nil {
		return errors.Wrap(err, "disconnect entity")
	}

	// #Delete pc abilities
	abs, err := app.AbilityStore.List(id)
	if err != nil {
		return errors.Wrap(err, "disconnect entity")
	}
	for _, ab := range abs {
		if err := app.AbilityStore.Remove(id, ab.ID); err != nil {
			return errors.Wrap(err, "disconnect entity")
		}
	}

	// #Delete pc entity
	if err := app.Store.Remove(id); err != nil {
		return errors.Wrap(err, "disconnect entity")
	}

	return nil
}

// FetchMRInventoryFromCache retrieve an inventory and check most recent store if not found.
func (app A) FetchMRInventoryFromCache(id gulid.ID, entityID gulid.ID) (entity.Inventory, error) {
	inv, err := app.InventoryStore.FetchInventory(id)
	if err != nil {
		switch errors.Cause(err).(type) {
		case gerrors.ErrNotFound:
			inv, err := app.MRInventoryStore.FetchMRInventory(entityID)
			if err != nil {
				return entity.Inventory{}, errors.Wrap(err, "fetch mr inventory from cache")
			}
			return inv, nil
		}
		return entity.Inventory{}, errors.Wrap(err, "fetch mr inventory from cache")
	}
	return inv, nil
}

// SetMRInventory set an inventory as most recent and traditional store.
func (app A) SetMRInventory(entityID gulid.ID, inv entity.Inventory) error {
	if err := app.InventoryStore.InsertInventory(inv); err != nil {
		return errors.Wrap(err, "set mr inventory")
	}
	if err := app.MRInventoryStore.InsertMRInventory(entityID, inv); err != nil {
		return errors.Wrap(err, "set mr inventory")
	}
	return nil
}

// ErasePC remove a pc and clean associated abilities.
func (app A) ErasePC(accountID gulid.ID, id gulid.ID) error {

	// #Delete pc abilities
	abs, err := app.AbilityStore.List(id)
	if err != nil {
		return errors.Wrap(err, "remove pc")
	}
	for _, ab := range abs {
		if err := app.AbilityStore.Remove(id, ab.ID); err != nil {
			return errors.Wrap(err, "remove pc")
		}
	}

	// #Delete pc inventory
	if err := app.MRInventoryStore.RemoveMRInventory(id); err != nil {
		return errors.Wrap(err, "remove pc")
	}

	// #Delete pc
	if err := app.PCStore.RemovePC(accountID, id); err != nil {
		return errors.Wrap(err, "remove pc")
	}

	// #Add 1 to pc left
	pcLeft, err := app.PCLeftStore.FetchPCLeft(accountID)
	if err != nil {
		return errors.Wrap(err, "remove pc")
	}
	pcLeft = pcLeft - 1
	if err := app.PCLeftStore.InsertPCLeft(pcLeft, accountID); err != nil {
		return errors.Wrap(err, "remove pc")
	}

	return nil
}

// CheckPermission check if token has owner permission on source.
func (app A) CheckPermission(tok gulid.ID, id gulid.ID) error {

	// #Check permission token/source.
	permission, err := app.PermissionStore.FetchPermission(tok.String(), id.String())
	switch errors.Cause(err).(type) {
	case gerrors.ErrNotFound:
		return errors.Wrap(err, "check permission token")
	}
	if err == nil && account.ACL(permission.Value) != account.Owner {
		return errors.Wrap(
			gerrors.ErrInsufficientACLs{
				Value:  permission.Value,
				Source: tok.String(),
				Target: id.String(),
			},
			"check permission token",
		)
	}
	return errors.Wrap(err, "check permission token")
}
