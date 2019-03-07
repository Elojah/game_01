package main

import (
	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/pkg/errors"
)

const (
	username0 = "test_bwzFEZBgxu"
	password0 = "test_NYYovlaKoFZUVR"
)

// Case0 :
// - Subscribe
// - Unsubscribe
func Case0(s *auth.Service) error {
	if err := s.Subscribe(username0, password0); err != nil {
		return errors.Wrap(err, "case_0")
	}
	if err := s.Unsubscribe(username0, password0); err != nil {
		return errors.Wrap(err, "case_0")
	}
	return nil
}
