package main

import (
	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/pkg/errors"
)

const (
	username3 = "test_gzArksMUjske"
	password3 = "test_apMwqzFnAPhg"

	pcName3 = "test_emh"
	pcType3 = "01CE3J5ASXJSVC405QTES4M221" // mesmerist
)

// Case3 :
// - Subscribe
// - SignIn
// - CreatePC
// - ListPC
// - ConnectPC
// - DisconnectPC
// - SignOut
// - Unsubscribe
func Case3(s *auth.Service) error {
	if err := s.Subscribe(username3, password3); err != nil {
		return errors.Wrap(err, "case_3")
	}
	tok, err := s.SignIn(username3, password3)
	if err != nil {
		return errors.Wrap(err, "case_3")
	}
	if err := s.CreatePC(tok.ID, pcName3, pcType3); err != nil {
		return errors.Wrap(err, "case_3")
	}
	pcs, err := s.ListPC(tok.ID)
	if err != nil || len(pcs) != 1 {
		return errors.Wrap(err, "case_3")
	}
	if _, err := s.ConnectPC(tok.ID, pcs[0].ID); err != nil {
		return errors.Wrap(err, "case_3")
	}
	if err := s.DisconnectPC(tok.ID); err != nil {
		return errors.Wrap(err, "case_3")
	}
	if err := s.SignOut(tok.ID, username3); err != nil {
		return errors.Wrap(err, "case_3")
	}
	if err := s.Unsubscribe(username3, password3); err != nil {
		return errors.Wrap(err, "case_3")
	}
	return nil
}
