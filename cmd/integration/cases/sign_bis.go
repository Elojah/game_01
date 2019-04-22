package cases

import (
	"time"

	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/pkg/errors"
)

const (
	usernameSignBis = "test_sign_bis"
	passwordSignBis = "test_sign_bis" // nolint: gosec

	pcNameSignBis = "test_sign_bis"
	pcTypeSignBis = "01CE3J5ASXJSVC405QTES4M221" // mesmerist

	pcSpawnSignBis = "01D6WJF3XF8ADHAGASDR6PW12P"
)

// SignBis :
// - Subscribe
// - SignIn
// - CreatePC
// - ListPC
// - ConnectPC
// - SignOut
// - Unsubscribe
// Test a signout while being connected
func SignBis(s *auth.Service) error {
	if err := s.Subscribe(usernameSignBis, passwordSignBis); err != nil {
		return errors.Wrap(err, "case_sign_bis")
	}
	tok, err := s.SignIn(usernameSignBis, passwordSignBis)
	if err != nil {
		return errors.Wrap(err, "case_sign_bis")
	}
	if err := s.CreatePC(tok.ID, pcNameSignBis, pcTypeSignBis, pcSpawnSignBis); err != nil {
		return errors.Wrap(err, "case_sign_bis")
	}
	pcs, err := s.ListPC(tok.ID)
	if err != nil || len(pcs) != 1 {
		return errors.Wrap(err, "case_sign_bis")
	}
	if _, err := s.ConnectPC(tok.ID, pcs[0].ID); err != nil {
		return errors.Wrap(err, "case_sign_bis")
	}
	// Wait for sequencer/subs to be ready
	time.Sleep(50 * time.Millisecond)
	if err := s.SignOut(tok.ID, usernameSignBis); err != nil {
		return errors.Wrap(err, "case_sign_bis")
	}
	if err := s.Unsubscribe(usernameSignBis, passwordSignBis); err != nil {
		return errors.Wrap(err, "case_sign_bis")
	}
	return nil
}
