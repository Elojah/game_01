package main

import (
	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/pkg/errors"
)

const (
	username2 = "test_bwzFEZBgxu"
	password2 = "test_NYYovlaKoFZUVR"

	pcName2 = "test_hBBJFsDkrDyUom"
	pcType2 = "01CE3J5ASXJSVC405QTES4M221" // mesmerist
)

// Case2 :
// - Subscribe
// - SignIn
// - CreatePC
// - SignOut
// - Unsubscribe
func Case2(s *auth.Service) error {
	if err := s.Subscribe(username2, password2); err != nil {
		return errors.Wrap(err, "case_2")
	}
	tok, err := s.SignIn(username2, password2)
	if err != nil {
		return errors.Wrap(err, "case_2")
	}
	if err := s.CreatePC(tok.ID, pcName2, pcType2); err != nil {
		return errors.Wrap(err, "case_2")
	}
	if err := s.SignOut(tok.ID, username2); err != nil {
		return errors.Wrap(err, "case_2")
	}
	if err := s.Unsubscribe(username2, password2); err != nil {
		return errors.Wrap(err, "case_2")
	}
	return nil
}
