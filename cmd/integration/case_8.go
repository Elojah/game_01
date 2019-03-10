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
	username8_0 = "test_hzOADnuxelP"
	password8_0 = "test_FRKsnfbDFFd"
	pcName8_0   = "test_uxbQmPNwdyXxJ"

	username8_1 = "test_yNRkOPVlGbWod"
	password8_1 = "test_hSNMmjIuDCpONZAx"
	pcName8_1   = "test_qAp"

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
	tok_0, err := as.SignIn(username8_0, password8_0)
	if err != nil {
		return errors.Wrap(err, "case_8")
	}
	if err := as.CreatePC(tok_0.ID, pcName8_0, pcType8); err != nil {
		return errors.Wrap(err, "case_8")
	}
	pcs_0, err := as.ListPC(tok_0.ID)
	if err != nil || len(pcs_0) != 1 {
		return errors.Wrap(err, "case_8")
	}
	ent_0, err := as.ConnectPC(tok_0.ID, pcs_0[0].ID)
	if err != nil {
		return errors.Wrap(err, "case_8")
	}

	// #1
	if err := as.Subscribe(username8_1, password8_1); err != nil {
		return errors.Wrap(err, "case_8")
	}
	tok_1, err := as.SignIn(username8_1, password8_1)
	if err != nil {
		return errors.Wrap(err, "case_8")
	}
	if err := as.CreatePC(tok_1.ID, pcName8_1, pcType8); err != nil {
		return errors.Wrap(err, "case_8")
	}
	pcs_1, err := as.ListPC(tok_1.ID)
	if err != nil || len(pcs_1) != 1 {
		return errors.Wrap(err, "case_8")
	}
	ent_1, err := as.ConnectPC(tok_1.ID, pcs_1[0].ID)
	if err != nil {
		return errors.Wrap(err, "case_8")
	}

	// Wait for sequencer/subs to be ready
	time.Sleep(80 * time.Millisecond)
	// Retrieve current entity state

	ent_0, err = cs.GetState(ent_0.ID, 50)
	if err != nil {
		return errors.Wrap(err, "case_8")
	}
	ent_1, err = cs.GetState(ent_1.ID, 50)
	if err != nil {
		return errors.Wrap(err, "case_8")
	}

	// Starter is unique: 01CF001HTBA3CDR1ERJ6RF183A (1024, 1024, 1024)
	if err := ts.EntityMove(ent_0.ID, geometry.Position{
		SectorID: ent_0.Position.SectorID,
		Coord:    geometry.Vec3{X: 1020, Y: 0, Z: 0},
	}); err != nil {
		return errors.Wrap(err, "case_8")
	}

	// Move ent_1 close to caster
	if err := ts.EntityMove(ent_1.ID, geometry.Position{
		SectorID: ent_1.Position.SectorID,
		Coord:    geometry.Vec3{X: 10, Y: 10, Z: 0},
	}); err != nil {
		return errors.Wrap(err, "case_8")
	}

	// Cast from ent_0 to ent_1 with starter skill
	if err := cs.Cast(tok_0.ID, event.Cast{
		Source:    ent_0.ID,
		AbilityID: gulid.MustParse("01CP2Z4SDEWZK8YF29E08GPDVC"),
		Targets: map[string]ability.Targets{
			"01CPFBN88EESQ4QA8N820RV924": ability.Targets{
				Entities: []gulid.ID{ent_1.ID},
			},
		},
	}); err != nil {
		return errors.Wrap(err, "case_8")
	}

	time.Sleep(1000 * time.Millisecond) // cast time last 1000 ms

	// Check entity caster used mana
	_, err = cs.GetStateAt(ent_0.ID, 500, func(actual entity.E) bool {
		return actual.MP == 250-10
	})
	// Check entity target took damage
	_, err = cs.GetStateAt(ent_1.ID, 50, func(actual entity.E) bool {
		return actual.HP == 150-30
	})

	// #0
	if err := as.DisconnectPC(tok_0.ID); err != nil {
		return errors.Wrap(err, "case_8")
	}
	if err := as.SignOut(tok_0.ID, username8_0); err != nil {
		return errors.Wrap(err, "case_8")
	}
	if err := as.Unsubscribe(username8_0, password8_0); err != nil {
		return errors.Wrap(err, "case_8")
	}
	// #1
	if err := as.DisconnectPC(tok_1.ID); err != nil {
		return errors.Wrap(err, "case_8")
	}
	if err := as.SignOut(tok_1.ID, username8_1); err != nil {
		return errors.Wrap(err, "case_8")
	}
	if err := as.Unsubscribe(username8_1, password8_1); err != nil {
		return errors.Wrap(err, "case_8")
	}
	return nil
}
