package cases

import (
	"time"

	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/pkg/errors"
)

const (
	usernameConnectPC = "test_connect_pc"
	passwordConnectPC = "test_connect_pc" // nolint: gosec

	pcNameConnectPC = "test_connect_pc"
	pcTypeConnectPC = "01CE3J5ASXJSVC405QTES4M221" // mesmerist
)

// ConnectPC :
// - Subscribe
// - SignIn
// - CreatePC
// - ListPC
// - ConnectPC
// - DisconnectPC
// - SignOut
// - Unsubscribe
func ConnectPC(s *auth.Service) error {
	if err := s.Subscribe(usernameConnectPC, passwordConnectPC); err != nil {
		return errors.Wrap(err, "case_connect_pc")
	}
	tok, err := s.SignIn(usernameConnectPC, passwordConnectPC)
	if err != nil {
		return errors.Wrap(err, "case_connect_pc")
	}
	if err := s.CreatePC(tok.ID, pcNameConnectPC, pcTypeConnectPC); err != nil {
		return errors.Wrap(err, "case_connect_pc")
	}
	pcs, err := s.ListPC(tok.ID)
	if err != nil || len(pcs) != 1 {
		return errors.Wrap(err, "case_connect_pc")
	}
	if _, err := s.ConnectPC(tok.ID, pcs[0].ID); err != nil {
		return errors.Wrap(err, "case_connect_pc")
	}
	// Wait for sequencer/subs to be ready
	time.Sleep(50 * time.Millisecond)
	if err := s.DisconnectPC(tok.ID); err != nil {
		return errors.Wrap(err, "case_connect_pc")
	}
	if err := s.SignOut(tok.ID, usernameConnectPC); err != nil {
		return errors.Wrap(err, "case_connect_pc")
	}
	if err := s.Unsubscribe(usernameConnectPC, passwordConnectPC); err != nil {
		return errors.Wrap(err, "case_connect_pc")
	}
	return nil
}
