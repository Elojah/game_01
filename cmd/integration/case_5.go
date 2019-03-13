package main

import (
	"time"

	"github.com/pkg/errors"

	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/elojah/game_01/cmd/integration/client"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/geometry"
)

const (
	username5 = "test_case5"
	password5 = "test_case5" // nolint: gosec

	pcName5 = "test_case5"
	pcType5 = "01CE3J5ASXJSVC405QTES4M221" // mesmerist
)

// Case5 :
// - Subscribe
// - SignIn
// - CreatePC
// - ListPC
// - ConnectPC
// - MoveSameSector
// - DisconnectPC
// - SignOut
// - Unsubscribe
func Case5(as *auth.Service, cs *client.Service) error {
	if err := as.Subscribe(username5, password5); err != nil {
		return errors.Wrap(err, "case_5")
	}
	tok, err := as.SignIn(username5, password5)
	if err != nil {
		return errors.Wrap(err, "case_5")
	}
	if err := as.CreatePC(tok.ID, pcName5, pcType5); err != nil {
		return errors.Wrap(err, "case_5")
	}
	pcs, err := as.ListPC(tok.ID)
	if err != nil || len(pcs) != 1 {
		return errors.Wrap(err, "case_5")
	}
	ent, err := as.ConnectPC(tok.ID, pcs[0].ID)
	if err != nil {
		return errors.Wrap(err, "case_5")
	}
	// Wait for sequencer/subs to be ready
	time.Sleep(50 * time.Millisecond)
	// Retrieve current entity state
	ent, err = cs.GetState(ent.ID, 50)
	if err != nil {
		return errors.Wrap(err, "case_5")
	}
	// Move entity in 10/10/10 direction
	newCoord, err := cs.MoveSameSector(tok.ID, ent, geometry.Vec3{X: 10, Y: 10, Z: 10})
	if err != nil {
		return errors.Wrap(err, "case_5")
	}
	// Check entity moved at correct position
	_, err = cs.GetStateAt(ent.ID, 500, func(actual entity.E) bool {
		return actual.Position.SectorID.Compare(ent.Position.SectorID) == 0 &&
			actual.Position.Coord == newCoord
	})
	if err != nil {
		return errors.Wrap(err, "case_5")
	}
	if err := as.SignOut(tok.ID, username5); err != nil {
		return errors.Wrap(err, "case_5")
	}
	if err := as.Unsubscribe(username5, password5); err != nil {
		return errors.Wrap(err, "case_5")
	}
	return nil
}
