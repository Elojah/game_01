package cases

import (
	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/pkg/errors"
)

const (
	usernameDelPC = "test_del_pc"
	passwordDelPC = "test_del_pc" // nolint: gosec

	pcNameDelPC = "test_del_pc"
	pcTypeDelPC = "01CE3J5ASXJSVC405QTES4M221" // mesmerist

	pcSpawnDelPC = "01D6WJF3XF8ADHAGASDR6PW12P"
)

// DelPC :
// - Subscribe
// - SignIn
// - CreatePC
// - ListPC
// - DelPC
// - SignOut
// - Unsubscribe
func DelPC(s *auth.Service) error {
	if err := s.Subscribe(usernameDelPC, passwordDelPC); err != nil {
		return errors.Wrap(err, "case_del_pc")
	}
	tok, err := s.SignIn(usernameDelPC, passwordDelPC)
	if err != nil {
		return errors.Wrap(err, "case_del_pc")
	}
	if err := s.CreatePC(tok.ID, pcNameDelPC, pcTypeDelPC, pcSpawnDelPC); err != nil {
		return errors.Wrap(err, "case_del_pc")
	}
	pcs, err := s.ListPC(tok.ID)
	if err != nil || len(pcs) != 1 {
		return errors.Wrap(err, "case_del_pc")
	}
	if err := s.DelPC(tok.ID, pcs[0].ID); err != nil {
		return errors.Wrap(err, "case_del_pc")
	}
	if err := s.SignOut(tok.ID, usernameDelPC); err != nil {
		return errors.Wrap(err, "case_del_pc")
	}
	if err := s.Unsubscribe(usernameDelPC, passwordDelPC); err != nil {
		return errors.Wrap(err, "case_del_pc")
	}
	return nil
}
