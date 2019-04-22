package cases

import (
	"time"

	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/pkg/errors"
)

const (
	usernameDelPCBis = "test_del_pc_bis"
	passwordDelPCBis = "test_del_pc_bis" // nolint: gosec

	pcNameDelPCBis = "test_del_pc_bis"
	pcTypeDelPCBis = "01CE3J5ASXJSVC405QTES4M221" // mesmerist

	pcSpawnDelPCBis = "01D6WJF3XF8ADHAGASDR6PW12P"
)

// DelPCBis :
// - Subscribe
// - SignIn
// - CreatePC
// - ListPC
// - ConnectPC
// - DelPCBis
// - SignOut
// - Unsubscribe
// Test a del pc while being connected
func DelPCBis(s *auth.Service) error {
	if err := s.Subscribe(usernameDelPCBis, passwordDelPCBis); err != nil {
		return errors.Wrap(err, "case_del_pc_bis")
	}
	tok, err := s.SignIn(usernameDelPCBis, passwordDelPCBis)
	if err != nil {
		return errors.Wrap(err, "case_del_pc_bis")
	}
	if err := s.CreatePC(tok.ID, pcNameDelPCBis, pcTypeDelPCBis, pcSpawnDelPCBis); err != nil {
		return errors.Wrap(err, "case_del_pc_bis")
	}
	pcs, err := s.ListPC(tok.ID)
	if err != nil || len(pcs) != 1 {
		return errors.Wrap(err, "case_del_pc_bis")
	}
	if _, err := s.ConnectPC(tok.ID, pcs[0].ID); err != nil {
		return errors.Wrap(err, "case_del_pc_bis")
	}
	// Wait for sequencer/subs to be ready
	time.Sleep(50 * time.Millisecond)
	if err := s.DelPC(tok.ID, pcs[0].ID); err != nil {
		return errors.Wrap(err, "case_del_pc_bis")
	}
	if err := s.SignOut(tok.ID, usernameDelPCBis); err != nil {
		return errors.Wrap(err, "case_del_pc_bis")
	}
	if err := s.Unsubscribe(usernameDelPCBis, passwordDelPCBis); err != nil {
		return errors.Wrap(err, "case_del_pc_bis")
	}
	return nil
}
