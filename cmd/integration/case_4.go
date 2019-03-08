package main

import (
	"time"

	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/pkg/errors"
)

const (
	username4 = "test_zaVjEoRv"
	password4 = "test_rYy"

	pcName4 = "test_Jchkj"
	pcType4 = "01CE3J5ASXJSVC405QTES4M221" // mesmerist
)

// Case4 :
// - Subscribe
// - SignIn
// - CreatePC
// - ListPC
// - ConnectPC
// - SignOut
// - Unsubscribe
// Test a signout while being connected
func Case4(s *auth.Service) error {
	if err := s.Subscribe(username4, password4); err != nil {
		return errors.Wrap(err, "case_4")
	}
	tok, err := s.SignIn(username4, password4)
	if err != nil {
		return errors.Wrap(err, "case_4")
	}
	if err := s.CreatePC(tok.ID, pcName4, pcType4); err != nil {
		return errors.Wrap(err, "case_4")
	}
	pcs, err := s.ListPC(tok.ID)
	if err != nil || len(pcs) != 1 {
		return errors.Wrap(err, "case_4")
	}
	if _, err := s.ConnectPC(tok.ID, pcs[0].ID); err != nil {
		return errors.Wrap(err, "case_4")
	}
	// Wait for sequencer/subs to be ready
	time.Sleep(50 * time.Millisecond)
	if err := s.SignOut(tok.ID, username4); err != nil {
		return errors.Wrap(err, "case_4")
	}
	if err := s.Unsubscribe(username4, password4); err != nil {
		return errors.Wrap(err, "case_4")
	}
	return nil
}
