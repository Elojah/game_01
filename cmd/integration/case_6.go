package main

import (
	"time"

	"github.com/pkg/errors"

	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/elojah/game_01/cmd/integration/client"
	"github.com/elojah/game_01/cmd/integration/tool"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/geometry"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

const (
	username6 = "test_case6"
	password6 = "test_case6" // nolint: gosec

	pcName6 = "test_case6"
	pcType6 = "01CE3J5ASXJSVC405QTES4M221" // mesmerist
)

// Case5 :
// - Subscribe
// - SignIn
// - CreatePC
// - ListPC
// - ConnectPC
// - EntityMove
// - Move
// - DisconnectPC
// - SignOut
// - Unsubscribe
func Case6(as *auth.Service, cs *client.Service, ts *tool.Service) error {
	if err := as.Subscribe(username6, password6); err != nil {
		return errors.Wrap(err, "case_6")
	}
	tok, err := as.SignIn(username6, password6)
	if err != nil {
		return errors.Wrap(err, "case_6")
	}
	if err := as.CreatePC(tok.ID, pcName6, pcType6); err != nil {
		return errors.Wrap(err, "case_6")
	}
	pcs, err := as.ListPC(tok.ID)
	if err != nil || len(pcs) != 1 {
		return errors.Wrap(err, "case_6")
	}
	ent, err := as.ConnectPC(tok.ID, pcs[0].ID)
	if err != nil {
		return errors.Wrap(err, "case_6")
	}

	// Wait for sequencer/subs to be ready
	time.Sleep(50 * time.Millisecond)

	// Retrieve current entity state
	ent, err = cs.GetState(ent.ID, 50)
	if err != nil {
		return errors.Wrap(err, "case_6")
	}
	// Force move entity on sector border with tool
	// Starter is unique: 01CF001HTBA3CDR1ERJ6RF183A (1024, 1024, 1024)
	if err := ts.EntityMove(ent.ID, geometry.Position{
		SectorID: ent.Position.SectorID,
		Coord:    geometry.Vec3{X: 1024, Y: 0, Z: 0},
	}); err != nil {
		return errors.Wrap(err, "case_6")
	}
	// Wait for move to be effective
	time.Sleep(50 * time.Millisecond)

	// Neighbour is 01CKQQPVZN5KQC8XC9Q9NK8YXQ relative at (1024, 0, 0)
	/*
		       1024
		+---+---+ 1024
		|   |   |
		|   |   |
		0---+---+
		|   |
		|   01CKQQPVZN5KQC8XC9Q9NK8YXQ
		|
		01CF001HTBA3CDR1ERJ6RF183A
	*/
	newPos := geometry.Position{
		SectorID: gulid.MustParse("01CKQQPVZN5KQC8XC9Q9NK8YXQ"),
		Coord:    geometry.Vec3{X: 10, Y: 0, Z: 0},
	}
	// Move entity at X:10 in next sector
	if err := cs.Move(tok.ID, ent, newPos); err != nil {
		return errors.Wrap(err, "case_6")
	}
	// Check entity moved at correct position
	_, err = cs.GetStateAt(ent.ID, 500, func(actual entity.E) bool {
		return actual.Position.SectorID.Compare(newPos.SectorID) == 0 &&
			actual.Position.Coord == newPos.Coord
	})
	if err != nil {
		return errors.Wrap(err, "case_6")
	}

	if err := as.DisconnectPC(tok.ID); err != nil {
		return errors.Wrap(err, "case_6")
	}
	if err := as.SignOut(tok.ID, username6); err != nil {
		return errors.Wrap(err, "case_6")
	}
	if err := as.Unsubscribe(username6, password6); err != nil {
		return errors.Wrap(err, "case_6")
	}
	return nil
}
