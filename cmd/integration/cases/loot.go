package cases

import (
	"time"

	"github.com/pkg/errors"

	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/elojah/game_01/cmd/integration/client"
	"github.com/elojah/game_01/cmd/integration/tool"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/item"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

const (
	usernameLoot = "test_loot"
	passwordLoot = "test_loot" // nolint: gosec
	pcNameLoot   = "test_loot"

	itemNameLoot = "test_loot"
	entNameLoot  = "test_loot"

	pcTypeLoot = "01CE3J5ASXJSVC405QTES4M221" // mesmerist

	pcSpawnLoot = "01D6WJF3XF8ADHAGASDR6PW12P"
)

// Loot :
// #
// - Subscribe
// - SignIn
// - CreatePC
// - ListPC
// - ConnectPC
// #
// MoveEntity
// CreateItem
// CreateInventory
// CreateLootableEntity
// Loot
// #
// - SignOut
// - Unsubscribe
func Loot(as *auth.Service, cs *client.Service, ts *tool.Service) error {

	if err := as.Subscribe(usernameLoot, passwordLoot); err != nil {
		return errors.Wrap(err, "case_loot")
	}
	tok, err := as.SignIn(usernameLoot, passwordLoot)
	if err != nil {
		return errors.Wrap(err, "case_loot")
	}
	if err := as.CreatePC(tok.ID, pcNameLoot, pcTypeLoot, pcSpawnLoot); err != nil {
		return errors.Wrap(err, "case_loot")
	}
	pcs, err := as.ListPC(tok.ID)
	if err != nil || len(pcs) != 1 {
		return errors.Wrap(err, "case_loot")
	}
	ent, err := as.ConnectPC(tok.ID, pcs[0].ID)
	if err != nil {
		return errors.Wrap(err, "case_loot")
	}

	// Wait for sequencer/subs to be ready
	time.Sleep(50 * time.Millisecond)
	// Retrieve current entity state

	ent, err = cs.GetState(ent.ID, 50)
	if err != nil {
		return errors.Wrap(err, "case_loot")
	}

	// Starter is unique: 01CF001HTBA3CDR1ERJ6RF183A (1024, 1024, 1024)
	if err := ts.EntityMove(ent.ID, geometry.Position{
		SectorID: ent.Position.SectorID,
		Coord:    geometry.Vec3{X: 500, Y: 0, Z: 0},
	}); err != nil {
		return errors.Wrap(err, "case_loot")
	}
	// Create item
	it := item.I{
		ID:   gulid.NewID(),
		Name: itemNameLoot,
		Icon: gulid.NewID(),
		Mesh: gulid.NewID(),
		Type: item.Type{
			Orb: &item.Orb{AbilityID: gulid.MustParse("01CP2Z4SDEWZK8YF29E07GPDVC")},
		},
	}
	if err := ts.Item([]item.I{it}); err != nil {
		return errors.Wrap(err, "case_loot")
	}
	lootableInventory := entity.Inventory{
		ID:    gulid.NewID(),
		Size_: 1,
		Items: map[string]uint64{
			it.ID.String(): 1,
		},
	}
	if err := ts.Inventory([]entity.Inventory{lootableInventory}); err != nil {
		return errors.Wrap(err, "case_loot")
	}

	// Create lootable entity
	target := entity.E{
		ID:        gulid.NewID(),
		Name:      entNameLoot,
		Direction: geometry.Vec3{},
		Position: geometry.Position{
			SectorID: ent.Position.SectorID,
			Coord:    geometry.Vec3{X: 501, Y: 0, Z: 0},
		},
		InventoryID: lootableInventory.ID,
	}
	if err := ts.Entity([]entity.E{target}); err != nil {
		return errors.Wrap(err, "case_loot")
	}
	if err := ts.Loot([]gulid.ID{target.ID}); err != nil {
		return errors.Wrap(err, "case_loot")
	}
	// Wait for moves to be effective
	time.Sleep(50 * time.Millisecond)

	// Loot item from lootableEntity
	if err := cs.Loot(tok.ID, event.Loot{
		Source:   ent.ID,
		TargetID: target.ID,
		ItemID:   it.ID,
	}); err != nil {
		return errors.Wrap(err, "case_loot")
	}

	// Check entity looter used mana
	_, err = cs.GetStateAt(ent.ID, 500, func(actual entity.E) bool {
		return actual.InventoryID.Compare(ent.InventoryID) != 0
	})
	if err != nil {
		return errors.Wrap(err, "case_loot")
	}

	// #0
	if err := as.SignOut(tok.ID, usernameLoot); err != nil {
		return errors.Wrap(err, "case_loot")
	}
	if err := as.Unsubscribe(usernameLoot, passwordLoot); err != nil {
		return errors.Wrap(err, "case_loot")
	}
	return nil
}
