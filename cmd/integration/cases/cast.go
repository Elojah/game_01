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
	gulid "github.com/elojah/game_01/pkg/ulid"
)

const (
	usernameCast0 = "test_cast_0"
	passwordCast0 = "test_cast_0" // nolint: gosec
	pcNameCast0   = "test_cast_0"

	usernameCast1 = "test_cast_1"
	passwordCast1 = "test_cast_1" // nolint: gosec
	pcNameCast1   = "test_cast_1"

	pcTypeCast = "01CE3J5ASXJSVC405QTES4M221" // mesmerist
)

// Cast :
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
// Cast
// #0
// - DisconnectPC
// - SignOut
// - Unsubscribe
// #1
// - DisconnectPC
// - SignOut
// - Unsubscribe
func Cast(as *auth.Service, cs *client.Service, ts *tool.Service) error {

	// #0
	if err := as.Subscribe(usernameCast0, passwordCast0); err != nil {
		return errors.Wrap(err, "case_cast")
	}
	tok0, err := as.SignIn(usernameCast0, passwordCast0)
	if err != nil {
		return errors.Wrap(err, "case_cast")
	}
	if err := as.CreatePC(tok0.ID, pcNameCast0, pcTypeCast); err != nil {
		return errors.Wrap(err, "case_cast")
	}
	pcs0, err := as.ListPC(tok0.ID)
	if err != nil || len(pcs0) != 1 {
		return errors.Wrap(err, "case_cast")
	}
	ent0, err := as.ConnectPC(tok0.ID, pcs0[0].ID)
	if err != nil {
		return errors.Wrap(err, "case_cast")
	}

	// #1
	if err := as.Subscribe(usernameCast1, passwordCast1); err != nil {
		return errors.Wrap(err, "case_cast")
	}
	tok1, err := as.SignIn(usernameCast1, passwordCast1)
	if err != nil {
		return errors.Wrap(err, "case_cast")
	}
	if err := as.CreatePC(tok1.ID, pcNameCast1, pcTypeCast); err != nil {
		return errors.Wrap(err, "case_cast")
	}
	pcs1, err := as.ListPC(tok1.ID)
	if err != nil || len(pcs1) != 1 {
		return errors.Wrap(err, "case_cast")
	}
	ent1, err := as.ConnectPC(tok1.ID, pcs1[0].ID)
	if err != nil {
		return errors.Wrap(err, "case_cast")
	}

	// Wait for sequencer/subs to be ready
	time.Sleep(50 * time.Millisecond)
	// Retrieve current entity state

	ent0, err = cs.GetState(ent0.ID, 50)
	if err != nil {
		return errors.Wrap(err, "case_cast")
	}
	ent1, err = cs.GetState(ent1.ID, 50)
	if err != nil {
		return errors.Wrap(err, "case_cast")
	}

	// Starter is unique: 01CF001HTBA3CDR1ERJ6RF183A (1024, 1024, 1024)
	if err := ts.EntityMove(ent0.ID, geometry.Position{
		SectorID: ent0.Position.SectorID,
		Coord:    geometry.Vec3{X: 500, Y: 0, Z: 0},
	}); err != nil {
		return errors.Wrap(err, "case_cast")
	}
	// Move ent1 close to caster
	if err := ts.EntityMove(ent1.ID, geometry.Position{
		SectorID: ent1.Position.SectorID,
		Coord:    geometry.Vec3{X: 510, Y: 10, Z: 0},
	}); err != nil {
		return errors.Wrap(err, "case_cast")
	}
	// Wait for moves to be effective
	time.Sleep(50 * time.Millisecond)

	// Cast from ent0 to ent1 with starter skill
	if err := cs.Cast(tok0.ID, event.Cast{
		Source:    ent0.ID,
		AbilityID: gulid.MustParse("01CP2Z4SDEWZK8YF29E07GPDVC"),
		Targets: map[string]ability.Targets{
			"01CPFBN87EESQ4QA8N820RV924": {
				Entities: []gulid.ID{ent1.ID},
			},
		},
	}); err != nil {
		return errors.Wrap(err, "case_cast")
	}

	time.Sleep(1000 * time.Millisecond) // cast time last 1000 ms

	// Check entity caster used mana
	_, err = cs.GetStateAt(ent0.ID, 500, func(actual entity.E) bool {
		return actual.MP == 250-10
	})
	if err != nil {
		return errors.Wrap(err, "case_cast")
	}
	// Check entity target took damage
	_, err = cs.GetStateAt(ent1.ID, 500, func(actual entity.E) bool {
		return actual.HP == 150-30
	})
	if err != nil {
		return errors.Wrap(err, "case_cast")
	}

	// #0
	if err := as.SignOut(tok0.ID, usernameCast0); err != nil {
		return errors.Wrap(err, "case_cast")
	}
	if err := as.Unsubscribe(usernameCast0, passwordCast0); err != nil {
		return errors.Wrap(err, "case_cast")
	}
	// #1
	if err := as.SignOut(tok1.ID, usernameCast1); err != nil {
		return errors.Wrap(err, "case_cast")
	}
	if err := as.Unsubscribe(usernameCast1, passwordCast1); err != nil {
		return errors.Wrap(err, "case_cast")
	}
	return nil
}
