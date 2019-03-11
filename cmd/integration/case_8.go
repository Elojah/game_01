package main

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
	gulid "github.com/elojah/game_01/pkg/ulid"
)

const (
	username8_0 = "test_case8_0"
	password8_0 = "test_case8_0" // nolint: gosec
	pcName8_0   = "test_case8_0"

	username8_1 = "test_case8_1"
	password8_1 = "test_case8_1" // nolint: gosec
	pcName8_1   = "test_case8_1"

	pcType8 = "01CE3J5ASXJSVC405QTES4M221" // mesmerist
)

// Case8 :
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
// EntityMove1 (different sector)
// Cast
// #0
// - DisconnectPC
// - SignOut
// - Unsubscribe
// #1
// - DisconnectPC
// - SignOut
// - Unsubscribe
func Case8(as *auth.Service, cs *client.Service, ts *tool.Service) error {

	// #0
	if err := as.Subscribe(username8_0, password8_0); err != nil {
		return errors.Wrap(err, "case_8")
	}
	tok0, err := as.SignIn(username8_0, password8_0)
	if err != nil {
		return errors.Wrap(err, "case_8")
	}
	if err := as.CreatePC(tok0.ID, pcName8_0, pcType8); err != nil {
		return errors.Wrap(err, "case_8")
	}
	pcs0, err := as.ListPC(tok0.ID)
	if err != nil || len(pcs0) != 1 {
		return errors.Wrap(err, "case_8")
	}
	ent0, err := as.ConnectPC(tok0.ID, pcs0[0].ID)
	if err != nil {
		return errors.Wrap(err, "case_8")
	}

	// #1
	if err := as.Subscribe(username8_1, password8_1); err != nil {
		return errors.Wrap(err, "case_8")
	}
	tok1, err := as.SignIn(username8_1, password8_1)
	if err != nil {
		return errors.Wrap(err, "case_8")
	}
	if err := as.CreatePC(tok1.ID, pcName8_1, pcType8); err != nil {
		return errors.Wrap(err, "case_8")
	}
	pcs1, err := as.ListPC(tok1.ID)
	if err != nil || len(pcs1) != 1 {
		return errors.Wrap(err, "case_8")
	}
	ent1, err := as.ConnectPC(tok1.ID, pcs1[0].ID)
	if err != nil {
		return errors.Wrap(err, "case_8")
	}

	// Wait for sequencer/subs to be ready
	time.Sleep(80 * time.Millisecond)
	// Retrieve current entity state

	ent0, err = cs.GetState(ent0.ID, 50)
	if err != nil {
		return errors.Wrap(err, "case_8")
	}
	ent1, err = cs.GetState(ent1.ID, 50)
	if err != nil {
		return errors.Wrap(err, "case_8")
	}

	// Starter is unique: 01CF001HTBA3CDR1ERJ6RF183A (1024, 1024, 1024)
	if err := ts.EntityMove(ent0.ID, geometry.Position{
		SectorID: ent0.Position.SectorID,
		Coord:    geometry.Vec3{X: 1020, Y: 0, Z: 0},
	}); err != nil {
		return errors.Wrap(err, "case_8")
	}

	// Move ent1 close to caster
	if err := ts.EntityMove(ent1.ID, geometry.Position{
		SectorID: ent1.Position.SectorID,
		Coord:    geometry.Vec3{X: 10, Y: 10, Z: 0},
	}); err != nil {
		return errors.Wrap(err, "case_8")
	}

	// Cast from ent0 to ent1 with starter skill
	if err := cs.Cast(tok0.ID, event.Cast{
		Source:    ent0.ID,
		AbilityID: gulid.MustParse("01CP2Z4SDEWZK8YF29E08GPDVC"),
		Targets: map[string]ability.Targets{
			"01CPFBN88EESQ4QA8N820RV924": {
				Entities: []gulid.ID{ent1.ID},
			},
		},
	}); err != nil {
		return errors.Wrap(err, "case_8")
	}

	time.Sleep(1000 * time.Millisecond) // cast time last 1000 ms

	// Check entity caster used mana
	_, err = cs.GetStateAt(ent0.ID, 500, func(actual entity.E) bool {
		return actual.MP == 250-10
	})
	if err != nil {
		return errors.Wrap(err, "case_8")
	}

	// Check entity target took damage
	_, err = cs.GetStateAt(ent1.ID, 50, func(actual entity.E) bool {
		return actual.HP == 150-30
	})
	if err != nil {
		return errors.Wrap(err, "case_8")
	}

	// #0
	if err := as.DisconnectPC(tok0.ID); err != nil {
		return errors.Wrap(err, "case_8")
	}
	if err := as.SignOut(tok0.ID, username8_0); err != nil {
		return errors.Wrap(err, "case_8")
	}
	if err := as.Unsubscribe(username8_0, password8_0); err != nil {
		return errors.Wrap(err, "case_8")
	}
	// #1
	if err := as.DisconnectPC(tok1.ID); err != nil {
		return errors.Wrap(err, "case_8")
	}
	if err := as.SignOut(tok1.ID, username8_1); err != nil {
		return errors.Wrap(err, "case_8")
	}
	if err := as.Unsubscribe(username8_1, password8_1); err != nil {
		return errors.Wrap(err, "case_8")
	}
	return nil
}
