package main

import (
	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/pkg/errors"
)

const (
	username1 = "test_bwzFEZBgxu"
	password1 = "test_NYYovlaKoFZUVR"
)

// Case1 :
// - Subscribe
// - SignIn
// - SignOut
// - Unsubscribe
func Case1(s *auth.Service) error {
	if err := s.Subscribe(username1, password1); err != nil {
		return errors.Wrap(err, "case_1")
	}
	tok, err := s.SignIn(username1, password1)
	if err != nil {
		return errors.Wrap(err, "case_1")
	}
	if err := s.SignOut(tok.ID, username1); err != nil {
		return errors.Wrap(err, "case_1")
	}
	if err := s.Unsubscribe(username1, password1); err != nil {
		return errors.Wrap(err, "case_1")
	}
	return nil
}
