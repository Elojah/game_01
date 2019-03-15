package cases

import (
	"time"

	"github.com/pkg/errors"

	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/elojah/game_01/cmd/integration/client"
	"github.com/elojah/game_01/cmd/integration/tool"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/item"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

const (
	usernameConsume = "test_consume"
	passwordConsume = "test_consume" // nolint: gosec
	pcNameConsume   = "test_consume"

	itemNameConsume = "test_consume"

	pcTypeConsume = "01CE3J5ASXJSVC405QTES4M221" // mesmerist
)

// Consume :
// #
// - Subscribe
// - SignIn
// - CreatePC
// - ListPC
// - ConnectPC
// #
// MoveEntity
// CreateItem
// Consume
// #
// - SignOut
// - Unsubscribe
func Consume(as *auth.Service, cs *client.Service, ts *tool.Service) error {

	if err := as.Subscribe(usernameConsume, passwordConsume); err != nil {
		return errors.Wrap(err, "case_consume")
	}
	tok, err := as.SignIn(usernameConsume, passwordConsume)
	if err != nil {
		return errors.Wrap(err, "case_consume")
	}
	if err := as.CreatePC(tok.ID, pcNameConsume, pcTypeConsume); err != nil {
		return errors.Wrap(err, "case_consume")
	}
	pcs, err := as.ListPC(tok.ID)
	if err != nil || len(pcs) != 1 {
		return errors.Wrap(err, "case_consume")
	}
	ent, err := as.ConnectPC(tok.ID, pcs[0].ID)
	if err != nil {
		return errors.Wrap(err, "case_consume")
	}

	// Wait for sequencer/subs to be ready
	time.Sleep(50 * time.Millisecond)
	// Retrieve current entity state

	ent, err = cs.GetState(ent.ID, 50)
	if err != nil {
		return errors.Wrap(err, "case_consume")
	}
	// Create item
	it := item.I{
		ID:   gulid.NewID(),
		Name: itemNameConsume,
		Icon: gulid.NewID(),
		Mesh: gulid.NewID(),
		Type: item.Type{
			Orb: &item.Orb{AbilityID: gulid.MustParse("01CP2Z4SDEWZK8YF29E07GPDVC")},
		},
	}
	if err := ts.Item([]item.I{it}); err != nil {
		return errors.Wrap(err, "case_consume")
	}
	consumeableInventory := entity.Inventory{
		ID:    ent.InventoryID,
		Size_: 42,
		Items: map[string]uint64{
			it.ID.String(): 1,
		},
	}
	if err := ts.Inventory([]entity.Inventory{consumeableInventory}); err != nil {
		return errors.Wrap(err, "case_consume")
	}
	// Wait for moves to be effective
	time.Sleep(50 * time.Millisecond)

	// Consume item from consumeableEntity
	if err := cs.Consume(tok.ID, event.Consume{
		Source:   ent.ID,
		TargetID: ent.ID,
		ItemID:   it.ID,
	}); err != nil {
		return errors.Wrap(err, "case_consume")
	}

	// Check inventory changed because item was consumed
	_, err = cs.GetStateAt(ent.ID, 500, func(actual entity.E) bool {
		return actual.InventoryID.Compare(ent.InventoryID) != 0
	})
	if err != nil {
		return errors.Wrap(err, "case_consume")
	}

	// #0
	if err := as.SignOut(tok.ID, usernameConsume); err != nil {
		return errors.Wrap(err, "case_consume")
	}
	if err := as.Unsubscribe(usernameConsume, passwordConsume); err != nil {
		return errors.Wrap(err, "case_consume")
	}
	return nil
}
