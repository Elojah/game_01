package cases

import (
	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/pkg/errors"
)

const (
	usernameCreatePC = "test_create_pc"
	passwordCreatePC = "test_create_pc" // nolint: gosec

	pcNameCreatePC = "test_create_pc"
	pcTypeCreatePC = "01CE3J5ASXJSVC405QTES4M221" // mesmerist

	pcSpawnCreatePC = "01D6WJF3XF8ADHAGASDR6PW12P"
)

// CreatePC :
// - Subscribe
// - SignIn
// - CreatePC
// - SignOut
// - Unsubscribe
func CreatePC(s *auth.Service) error {
	if err := s.Subscribe(usernameCreatePC, passwordCreatePC); err != nil {
		return errors.Wrap(err, "case_create_pc")
	}
	tok, err := s.SignIn(usernameCreatePC, passwordCreatePC)
	if err != nil {
		return errors.Wrap(err, "case_create_pc")
	}
	if err := s.CreatePC(tok.ID, pcNameCreatePC, pcTypeCreatePC, pcSpawnCreatePC); err != nil {
		return errors.Wrap(err, "case_create_pc")
	}
	if err := s.SignOut(tok.ID, usernameCreatePC); err != nil {
		return errors.Wrap(err, "case_create_pc")
	}
	if err := s.Unsubscribe(usernameCreatePC, passwordCreatePC); err != nil {
		return errors.Wrap(err, "case_create_pc")
	}
	return nil
}
