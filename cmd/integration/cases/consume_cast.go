package cases

import (
	"time"

	"github.com/pkg/errors"

	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/elojah/game_01/cmd/integration/client"
	"github.com/elojah/game_01/cmd/integration/tool"
	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/item"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

const (
	usernameConsumeCast0 = "test_consume_cast_0"
	passwordConsumeCast0 = "test_consume_cast_0" // nolint: gosec
	pcNameConsumeCast0   = "test_consume_cast_0"

	usernameConsumeCast1 = "test_consume_cast_1"
	passwordConsumeCast1 = "test_consume_cast_1" // nolint: gosec
	pcNameConsumeCast1   = "test_consume_cast_1"

	itemNameConsumeCast = "test_consume_cast"

	pcTypeConsumeCast = "01CE3J5ASXJSVC405QTES4M221" // mesmerist

	pcSpawnConsumeCast = "01D6WJF3XF8ADHAGASDR6PW12P"
)

// ConsumeCast :
// #0
// - Subscribe
// - SignIn
// - CreatePC
// - ListPC
// - ConnectPC
// #1
// - Subscribe
// - SignIn
// - CreatePC
// - ListPC
// - ConnectPC
// #
// EntityMove0
// EntityMove1
// CreateItem
// Consume
// Cast
// #0
// - SignOut
// - Unsubscribe
// #1
// - SignOut
// - Unsubscribe
func ConsumeCast(as *auth.Service, cs *client.Service, ts *tool.Service) error {

	// #0
	if err := as.Subscribe(usernameConsumeCast0, passwordConsumeCast0); err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}
	tok0, err := as.SignIn(usernameConsumeCast0, passwordConsumeCast0)
	if err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}
	if err := as.CreatePC(tok0.ID, pcNameConsumeCast0, pcTypeConsumeCast, pcSpawnConsumeCast); err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}
	pcs0, err := as.ListPC(tok0.ID)
	if err != nil || len(pcs0) != 1 {
		return errors.Wrap(err, "case_consume_cast")
	}
	ent0, err := as.ConnectPC(tok0.ID, pcs0[0].ID)
	if err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}

	// #1
	if err := as.Subscribe(usernameConsumeCast1, passwordConsumeCast1); err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}
	tok1, err := as.SignIn(usernameConsumeCast1, passwordConsumeCast1)
	if err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}
	if err := as.CreatePC(tok1.ID, pcNameConsumeCast1, pcTypeConsumeCast, pcSpawnConsumeCast); err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}
	pcs1, err := as.ListPC(tok1.ID)
	if err != nil || len(pcs1) != 1 {
		return errors.Wrap(err, "case_consume_cast")
	}
	ent1, err := as.ConnectPC(tok1.ID, pcs1[0].ID)
	if err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}

	// Wait for sequencer/subs to be ready
	time.Sleep(50 * time.Millisecond)
	// Retrieve current entity state

	ent0, err = cs.GetState(ent0.ID, 50)
	if err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}
	ent1, err = cs.GetState(ent1.ID, 50)
	if err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}

	// Starter is unique: 01CF001HTBA3CDR1ERJ6RF183A (1024, 1024, 1024)
	if err := ts.EntityMove(ent0.ID, geometry.Position{
		SectorID: ent0.Position.SectorID,
		Coord:    geometry.Vec3{X: 500, Y: 0, Z: 0},
	}); err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}
	// Move ent1 close to caster
	if err := ts.EntityMove(ent1.ID, geometry.Position{
		SectorID: ent1.Position.SectorID,
		Coord:    geometry.Vec3{X: 510, Y: 10, Z: 0},
	}); err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}
	// Wait for moves to be effective
	time.Sleep(50 * time.Millisecond)

	// Create item
	it := item.I{
		ID:   gulid.NewID(),
		Name: itemNameConsumeCast,
		Icon: gulid.NewID(),
		Mesh: gulid.NewID(),
		Type: item.Type{
			Orb: &item.Orb{AbilityID: gulid.MustParse("01D614CA6ZJZTDQ7B54ZDH9WN7")},
		},
	}
	if err := ts.Item([]item.I{it}); err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}
	consumeableInventory := entity.Inventory{
		ID:    ent0.InventoryID,
		Size_: 42,
		Items: map[string]uint64{
			it.ID.String(): 1,
		},
	}
	if err := ts.Inventory([]entity.Inventory{consumeableInventory}); err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}
	// Wait for moves to be effective
	time.Sleep(50 * time.Millisecond)

	// Consume item from consumeableEntity
	if err := cs.Consume(tok0.ID, event.Consume{
		Source:   ent0.ID,
		TargetID: ent0.ID,
		ItemID:   it.ID,
	}); err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}

	// Check inventory changed because item was consumed
	_, err = cs.GetStateAt(ent0.ID, 500, func(actual entity.E) bool {
		return actual.InventoryID.Compare(ent0.InventoryID) != 0
	})
	if err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}

	// Cast from ent0 to ent1 with starter skill
	if err := cs.Cast(tok0.ID, event.Cast{
		Source:    ent0.ID,
		AbilityID: gulid.MustParse("01D614CA6ZJZTDQ7B54ZDH9WN7"),
		Targets: map[string]ability.Targets{
			"01D614E5FW3ZK463YD7M2DE6Q6": {
				Entities: []gulid.ID{ent1.ID},
			},
		},
	}); err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}

	time.Sleep(500 * time.Millisecond) // cast time last 500 ms

	// Check entity caster used mana
	_, err = cs.GetStateAt(ent0.ID, 500, func(actual entity.E) bool {
		return actual.MP == 250-15
	})
	if err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}
	// Check entity target took damage
	_, err = cs.GetStateAt(ent1.ID, 500, func(actual entity.E) bool {
		return actual.HP == 150-50
	})
	if err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}

	// #0
	if err := as.SignOut(tok0.ID, usernameConsumeCast0); err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}
	if err := as.Unsubscribe(usernameConsumeCast0, passwordConsumeCast0); err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}
	// #1
	if err := as.SignOut(tok1.ID, usernameConsumeCast1); err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}
	if err := as.Unsubscribe(usernameConsumeCast1, passwordConsumeCast1); err != nil {
		return errors.Wrap(err, "case_consume_cast")
	}
	return nil
}
