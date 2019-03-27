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
	usernameSpawn0 = "test_spawn_0"
	passwordSpawn0 = "test_spawn_0" // nolint: gosec
	pcNameSpawn0   = "test_spawn_0"

	usernameSpawn1 = "test_spawn_1"
	passwordSpawn1 = "test_spawn_1" // nolint: gosec
	pcNameSpawn1   = "test_spawn_1"

	pcTypeSpawn0 = "01CE3J622E95AQ2QS4XECMRFCV" // scavenger
	pcTypeSpawn1 = "01CE3J5ASXJSVC405QTES4M221" // mesmerist
)

// Spawn :
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
func Spawn(as *auth.Service, cs *client.Service, ts *tool.Service) error {

	// #0
	if err := as.Subscribe(usernameSpawn0, passwordSpawn0); err != nil {
		return errors.Wrap(err, "case_spawn")
	}
	tok0, err := as.SignIn(usernameSpawn0, passwordSpawn0)
	if err != nil {
		return errors.Wrap(err, "case_spawn")
	}
	if err := as.CreatePC(tok0.ID, pcNameSpawn0, pcTypeSpawn0); err != nil {
		return errors.Wrap(err, "case_spawn")
	}
	pcs0, err := as.ListPC(tok0.ID)
	if err != nil || len(pcs0) != 1 {
		return errors.Wrap(err, "case_spawn")
	}
	ent0, err := as.ConnectPC(tok0.ID, pcs0[0].ID)
	if err != nil {
		return errors.Wrap(err, "case_spawn")
	}

	// #1
	if err := as.Subscribe(usernameSpawn1, passwordSpawn1); err != nil {
		return errors.Wrap(err, "case_spawn")
	}
	tok1, err := as.SignIn(usernameSpawn1, passwordSpawn1)
	if err != nil {
		return errors.Wrap(err, "case_spawn")
	}
	if err := as.CreatePC(tok1.ID, pcNameSpawn1, pcTypeSpawn1); err != nil {
		return errors.Wrap(err, "case_spawn")
	}
	pcs1, err := as.ListPC(tok1.ID)
	if err != nil || len(pcs1) != 1 {
		return errors.Wrap(err, "case_spawn")
	}
	ent1, err := as.ConnectPC(tok1.ID, pcs1[0].ID)
	if err != nil {
		return errors.Wrap(err, "case_spawn")
	}

	// Wait for sequencer/subs to be ready
	time.Sleep(50 * time.Millisecond)
	// Retrieve current entity state

	ent0, err = cs.GetState(ent0.ID, 50)
	if err != nil {
		return errors.Wrap(err, "case_spawn")
	}
	ent1, err = cs.GetState(ent1.ID, 50)
	if err != nil {
		return errors.Wrap(err, "case_spawn")
	}

	// Starter is unique: 01CF001HTBA3CDR1ERJ6RF183A (1024, 1024, 1024)
	if err := ts.EntityMove(ent0.ID, geometry.Position{
		SectorID: ent0.Position.SectorID,
		Coord:    geometry.Vec3{X: 500, Y: 0, Z: 0},
	}); err != nil {
		return errors.Wrap(err, "case_spawn")
	}
	// Move ent1 close to caster
	if err := ts.EntityMove(ent1.ID, geometry.Position{
		SectorID: ent1.Position.SectorID,
		Coord:    geometry.Vec3{X: 500, Y: 5, Z: 0},
	}); err != nil {
		return errors.Wrap(err, "case_spawn")
	}
	// Wait for moves to be effective
	time.Sleep(50 * time.Millisecond)

	// Cast from ent0 to ent1 with starter skill
	if err := cs.Cast(tok0.ID, event.Cast{
		Source:    ent0.ID,
		AbilityID: gulid.MustParse("01D6WRF5KQJFHZFKQGFVVJVM7P"),
		Targets: map[string]ability.Targets{
			"01D6WS5MRXRG7KGM1190SBJFHA": {
				Entities: []gulid.ID{ent1.ID},
			},
		},
	}); err != nil {
		return errors.Wrap(err, "case_spawn")
	}

	time.Sleep(1000 * time.Millisecond) // cast time last 1000 ms

	// Check entity caster used mana
	_, err = cs.GetStateAt(ent0.ID, 500, func(actual entity.E) bool {
		return actual.MP == 250-200
	})
	if err != nil {
		return errors.Wrap(err, "case_spawn")
	}
	// // Check entity target took damage
	// _, err = cs.GetStateAt(ent1.ID, 500, func(actual entity.E) bool {
	// 	return actual.Dead == true
	// })
	// if err != nil {
	// 	return errors.Wrap(err, "case_spawn")
	// }

	// #0
	if err := as.SignOut(tok0.ID, usernameSpawn0); err != nil {
		return errors.Wrap(err, "case_spawn")
	}
	if err := as.Unsubscribe(usernameSpawn0, passwordSpawn0); err != nil {
		return errors.Wrap(err, "case_spawn")
	}
	// #1
	if err := as.SignOut(tok1.ID, usernameSpawn1); err != nil {
		return errors.Wrap(err, "case_spawn")
	}
	if err := as.Unsubscribe(usernameSpawn1, passwordSpawn1); err != nil {
		return errors.Wrap(err, "case_spawn")
	}
	return nil
}
