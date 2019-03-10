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
	username7_0 = "test_AABuDcSdojJirTUl"
	password7_0 = "test_NMxa"
	pcName7_0   = "test_vUWjHGvgrPvPfOEsZh"

	username7_1 = "test_qFlZQbvOrXIDNOmvbcnF"
	password7_1 = "test_hQpmMWfMk"
	pcName7_1   = "test_TGqiqWAdNhscn"

	pcType7 = "01CE3J5ASXJSVC405QTES4M221" // mesmerist
)

// Case7 :
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
// #0
// - DisconnectPC
// - SignOut
// - Unsubscribe
// #1
// - DisconnectPC
// - SignOut
// - Unsubscribe
func Case7(as *auth.Service, cs *client.Service, ts *tool.Service) error {

	// #0
	if err := as.Subscribe(username7_0, password7_0); err != nil {
		return errors.Wrap(err, "case_7")
	}
	tok_0, err := as.SignIn(username7_0, password7_0)
	if err != nil {
		return errors.Wrap(err, "case_7")
	}
	if err := as.CreatePC(tok_0.ID, pcName7_0, pcType7); err != nil {
		return errors.Wrap(err, "case_7")
	}
	pcs_0, err := as.ListPC(tok_0.ID)
	if err != nil || len(pcs_0) != 1 {
		return errors.Wrap(err, "case_7")
	}
	ent_0, err := as.ConnectPC(tok_0.ID, pcs_0[0].ID)
	if err != nil {
		return errors.Wrap(err, "case_7")
	}

	// #1
	if err := as.Subscribe(username7_1, password7_1); err != nil {
		return errors.Wrap(err, "case_7")
	}
	tok_1, err := as.SignIn(username7_1, password7_1)
	if err != nil {
		return errors.Wrap(err, "case_7")
	}
	if err := as.CreatePC(tok_1.ID, pcName7_1, pcType7); err != nil {
		return errors.Wrap(err, "case_7")
	}
	pcs_1, err := as.ListPC(tok_1.ID)
	if err != nil || len(pcs_1) != 1 {
		return errors.Wrap(err, "case_7")
	}
	ent_1, err := as.ConnectPC(tok_1.ID, pcs_1[0].ID)
	if err != nil {
		return errors.Wrap(err, "case_7")
	}

	// Wait for sequencer/subs to be ready
	time.Sleep(70 * time.Millisecond)
	// Retrieve current entity state

	ent_0, err = cs.GetState(ent_0.ID, 50)
	if err != nil {
		return errors.Wrap(err, "case_7")
	}
	ent_1, err = cs.GetState(ent_1.ID, 50)
	if err != nil {
		return errors.Wrap(err, "case_7")
	}

	// Starter is unique: 01CF001HTBA3CDR1ERJ6RF183A (1024, 1024, 1024)
	if err := ts.EntityMove(ent_0.ID, geometry.Position{
		SectorID: ent_0.Position.SectorID,
		Coord:    geometry.Vec3{X: 500, Y: 0, Z: 0},
	}); err != nil {
		return errors.Wrap(err, "case_7")
	}

	// Move ent_1 close to caster
	if err := ts.EntityMove(ent_1.ID, geometry.Position{
		SectorID: ent_1.Position.SectorID,
		Coord:    geometry.Vec3{X: 510, Y: 10, Z: 0},
	}); err != nil {
		return errors.Wrap(err, "case_7")
	}

	// Cast from ent_0 to ent_1 with starter skill
	if err := cs.Cast(tok_0.ID, event.Cast{
		Source:    ent_0.ID,
		AbilityID: gulid.MustParse("01CP2Z4SDEWZK8YF29E07GPDVC"),
		Targets: map[string]ability.Targets{
			"01CPFBN87EESQ4QA8N820RV924": ability.Targets{
				Entities: []gulid.ID{ent_1.ID},
			},
		},
	}); err != nil {
		return errors.Wrap(err, "case_7")
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
		return errors.Wrap(err, "case_7")
	}
	if err := as.SignOut(tok_0.ID, username7_0); err != nil {
		return errors.Wrap(err, "case_7")
	}
	if err := as.Unsubscribe(username7_0, password7_0); err != nil {
		return errors.Wrap(err, "case_7")
	}
	// #1
	if err := as.DisconnectPC(tok_1.ID); err != nil {
		return errors.Wrap(err, "case_7")
	}
	if err := as.SignOut(tok_1.ID, username7_1); err != nil {
		return errors.Wrap(err, "case_7")
	}
	if err := as.Unsubscribe(username7_1, password7_1); err != nil {
		return errors.Wrap(err, "case_7")
	}
	return nil
}
