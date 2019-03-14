package cases

import (
	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/pkg/errors"
)

const (
	usernameSign = "test_sign"
	passwordSign = "test_sign" // nolint: gosec
)

// Sign :
// - Subscribe
// - SignIn
// - SignOut
// - Unsubscribe
func Sign(s *auth.Service) error {
	if err := s.Subscribe(usernameSign, passwordSign); err != nil {
		return errors.Wrap(err, "case_sign")
	}
	tok, err := s.SignIn(usernameSign, passwordSign)
	if err != nil {
		return errors.Wrap(err, "case_sign")
	}
	if err := s.SignOut(tok.ID, usernameSign); err != nil {
		return errors.Wrap(err, "case_sign")
	}
	if err := s.Unsubscribe(usernameSign, passwordSign); err != nil {
		return errors.Wrap(err, "case_sign")
	}
	return nil
}
