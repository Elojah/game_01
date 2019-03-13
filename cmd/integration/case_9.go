package main

import (
	"time"

	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/pkg/errors"
)

const (
	username9 = "test_case9"
	password9 = "test_case9" // nolint: gosec

	pcName9 = "test_case9"
	pcType9 = "01CE3J5ASXJSVC405QTES4M221" // mesmerist
)

// Case4 :
// - Subscribe
// - SignIn
// - CreatePC
// - ListPC
// - ConnectPC
// - DelPC
// - SignOut
// - Unsubscribe
// Test a del pc while being connected
func Case9(s *auth.Service) error {
	if err := s.Subscribe(username9, password9); err != nil {
		return errors.Wrap(err, "case_9")
	}
	tok, err := s.SignIn(username9, password9)
	if err != nil {
		return errors.Wrap(err, "case_9")
	}
	if err := s.CreatePC(tok.ID, pcName9, pcType9); err != nil {
		return errors.Wrap(err, "case_9")
	}
	pcs, err := s.ListPC(tok.ID)
	if err != nil || len(pcs) != 1 {
		return errors.Wrap(err, "case_9")
	}
	if _, err := s.ConnectPC(tok.ID, pcs[0].ID); err != nil {
		return errors.Wrap(err, "case_9")
	}
	// Wait for sequencer/subs to be ready
	time.Sleep(50 * time.Millisecond)
	if err := s.DelPC(tok.ID, pcs[0].ID); err != nil {
		return errors.Wrap(err, "case_9")
	}
	if err := s.SignOut(tok.ID, username9); err != nil {
		return errors.Wrap(err, "case_9")
	}
	if err := s.Unsubscribe(username9, password9); err != nil {
		return errors.Wrap(err, "case_9")
	}
	return nil
}
