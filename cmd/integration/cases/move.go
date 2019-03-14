package cases

import (
	"time"

	"github.com/pkg/errors"

	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/elojah/game_01/cmd/integration/client"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/geometry"
)

const (
	usernameMove = "test_move"
	passwordMove = "test_move" // nolint: gosec

	pcNameMove = "test_move"
	pcTypeMove = "01CE3J5ASXJSVC405QTES4M221" // mesmerist
)

// Move :
// - Subscribe
// - SignIn
// - CreatePC
// - ListPC
// - ConnectPC
// - MoveSameSector
// - DisconnectPC
// - SignOut
// - Unsubscribe
func Move(as *auth.Service, cs *client.Service) error {
	if err := as.Subscribe(usernameMove, passwordMove); err != nil {
		return errors.Wrap(err, "case_move")
	}
	tok, err := as.SignIn(usernameMove, passwordMove)
	if err != nil {
		return errors.Wrap(err, "case_move")
	}
	if err := as.CreatePC(tok.ID, pcNameMove, pcTypeMove); err != nil {
		return errors.Wrap(err, "case_move")
	}
	pcs, err := as.ListPC(tok.ID)
	if err != nil || len(pcs) != 1 {
		return errors.Wrap(err, "case_move")
	}
	ent, err := as.ConnectPC(tok.ID, pcs[0].ID)
	if err != nil {
		return errors.Wrap(err, "case_move")
	}
	// Wait for sequencer/subs to be ready
	time.Sleep(50 * time.Millisecond)
	// Retrieve current entity state
	ent, err = cs.GetState(ent.ID, 50)
	if err != nil {
		return errors.Wrap(err, "case_move")
	}
	// Move entity in 10/10/10 direction
	newCoord, err := cs.MoveSameSector(tok.ID, ent, geometry.Vec3{X: 10, Y: 10, Z: 10})
	if err != nil {
		return errors.Wrap(err, "case_move")
	}
	// Check entity moved at correct position
	_, err = cs.GetStateAt(ent.ID, 500, func(actual entity.E) bool {
		return actual.Position.SectorID.Compare(ent.Position.SectorID) == 0 &&
			actual.Position.Coord == newCoord
	})
	if err != nil {
		return errors.Wrap(err, "case_move")
	}
	if err := as.SignOut(tok.ID, usernameMove); err != nil {
		return errors.Wrap(err, "case_move")
	}
	if err := as.Unsubscribe(usernameMove, passwordMove); err != nil {
		return errors.Wrap(err, "case_move")
	}
	return nil
}
