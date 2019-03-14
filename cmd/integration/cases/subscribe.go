package cases

import (
	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/pkg/errors"
)

const (
	usernameSubscribe = "test_subscribe"
	passwordSubscribe = "test_subscribe" // nolint: gosec
)

// Subscribe :
// - Subscribe
// - Unsubscribe
func Subscribe(s *auth.Service) error {
	if err := s.Subscribe(usernameSubscribe, passwordSubscribe); err != nil {
		return errors.Wrap(err, "case_subscribe")
	}
	if err := s.Unsubscribe(usernameSubscribe, passwordSubscribe); err != nil {
		return errors.Wrap(err, "case_subscribe")
	}
	return nil
}
