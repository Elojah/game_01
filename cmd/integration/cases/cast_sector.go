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
	usernameCastSector0 = "test_case_cast_sector_0"
	passwordCastSector0 = "test_case_cast_sector_0" // nolint: gosec
	pcNameCastSector0   = "test_case_cast_sector_0"

	usernameCastSector1 = "test_case_cast_sector_1"
	passwordCastSector1 = "test_case_cast_sector_1" // nolint: gosec
	pcNameCastSector1   = "test_case_cast_sector_1"

	pcTypeCastSector = "01CE3J5ASXJSVC405QTES4M221" // mesmerist
)

// CastSector :
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
func CastSector(as *auth.Service, cs *client.Service, ts *tool.Service) error {

	// #0
	if err := as.Subscribe(usernameCastSector0, passwordCastSector0); err != nil {
		return errors.Wrap(err, "case_cast_sector")
	}
	tok0, err := as.SignIn(usernameCastSector0, passwordCastSector0)
	if err != nil {
		return errors.Wrap(err, "case_cast_sector")
	}
	if err := as.CreatePC(tok0.ID, pcNameCastSector0, pcTypeCastSector); err != nil {
		return errors.Wrap(err, "case_cast_sector")
	}
	pcs0, err := as.ListPC(tok0.ID)
	if err != nil || len(pcs0) != 1 {
		return errors.Wrap(err, "case_cast_sector")
	}
	ent0, err := as.ConnectPC(tok0.ID, pcs0[0].ID)
	if err != nil {
		return errors.Wrap(err, "case_cast_sector")
	}

	// #1
	if err := as.Subscribe(usernameCastSector1, passwordCastSector1); err != nil {
		return errors.Wrap(err, "case_cast_sector")
	}
	tok1, err := as.SignIn(usernameCastSector1, passwordCastSector1)
	if err != nil {
		return errors.Wrap(err, "case_cast_sector")
	}
	if err := as.CreatePC(tok1.ID, pcNameCastSector1, pcTypeCastSector); err != nil {
		return errors.Wrap(err, "case_cast_sector")
	}
	pcs1, err := as.ListPC(tok1.ID)
	if err != nil || len(pcs1) != 1 {
		return errors.Wrap(err, "case_cast_sector")
	}
	ent1, err := as.ConnectPC(tok1.ID, pcs1[0].ID)
	if err != nil {
		return errors.Wrap(err, "case_cast_sector")
	}

	// Wait for sequencer/subs to be ready
	time.Sleep(50 * time.Millisecond)
	// Retrieve current entity state

	ent0, err = cs.GetState(ent0.ID, 50)
	if err != nil {
		return errors.Wrap(err, "case_cast_sector")
	}
	ent1, err = cs.GetState(ent1.ID, 50)
	if err != nil {
		return errors.Wrap(err, "case_cast_sector")
	}

	// Starter is unique: 01CF001HTBA3CDR1ERJ6RF183A (1024, 1024, 1024)
	if err := ts.EntityMove(ent0.ID, geometry.Position{
		SectorID: ent0.Position.SectorID,
		Coord:    geometry.Vec3{X: 1020, Y: 0, Z: 0},
	}); err != nil {
		return errors.Wrap(err, "case_cast_sector")
	}

	// Move ent1 close to caster
	if err := ts.EntityMove(ent1.ID, geometry.Position{
		SectorID: gulid.MustParse("01CKQQPVZN5KQC8XC9Q9NK8YXQ"),
		Coord:    geometry.Vec3{X: 10, Y: 0, Z: 0},
	}); err != nil {
		return errors.Wrap(err, "case_cast_sector")
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
		return errors.Wrap(err, "case_cast_sector")
	}

	time.Sleep(1000 * time.Millisecond) // cast time last 1000 ms

	// Check entity caster used mana
	_, err = cs.GetStateAt(ent0.ID, 50, func(actual entity.E) bool {
		return actual.MP == 250-10
	})
	if err != nil {
		return errors.Wrap(err, "case_cast_sector")
	}

	// Check entity target took damage
	_, err = cs.GetStateAt(ent1.ID, 100, func(actual entity.E) bool {
		return actual.HP == 150-30
	})
	if err != nil {
		return errors.Wrap(err, "case_cast_sector")
	}

	// #0
	if err := as.SignOut(tok0.ID, usernameCastSector0); err != nil {
		return errors.Wrap(err, "case_cast_sector")
	}
	if err := as.Unsubscribe(usernameCastSector0, passwordCastSector0); err != nil {
		return errors.Wrap(err, "case_cast_sector")
	}
	// #1
	if err := as.SignOut(tok1.ID, usernameCastSector1); err != nil {
		return errors.Wrap(err, "case_cast_sector")
	}
	if err := as.Unsubscribe(usernameCastSector1, passwordCastSector1); err != nil {
		return errors.Wrap(err, "case_cast_sector")
	}
	return nil
}
