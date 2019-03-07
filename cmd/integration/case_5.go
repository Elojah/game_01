package main

import (
	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/elojah/game_01/cmd/integration/client"
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/pkg/errors"
)

const (
	username5 = "test_gzArksMUjske"
	password5 = "test_apMwqzFnAPhg"

	pcName5 = "test_emh"
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
	if err := cs.MoveSameSector(tok.ID, ent, geometry.Vec3{X: 52.0021, Y: 83.7427, Z: 57.2037}); err != nil {
		return errors.Wrap(err, "case_5")
	}
	if err := as.DisconnectPC(tok.ID); err != nil {
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
