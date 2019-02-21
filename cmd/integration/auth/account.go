package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/account"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// Subscribe a new account through HTTPS.
func (s *Service) Subscribe(username string, password string) error {
	raw, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return errors.Wrap(err, "create account")
	}
	resp, err := http.Post(s.url+"/subscribe", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return errors.Wrap(err, "create account")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "create account")
	}

	return nil
}

// SignIn sign to an already created account through HTTPS and returns corresponding token.
func (s *Service) SignIn(username string, password string) (account.Token, error) {
	raw, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return account.Token{}, errors.Wrap(err, "sign in")
	}
	resp, err := http.Post(s.url+"/signin", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return account.Token{}, errors.Wrap(err, "sign in")
	}
	if resp.StatusCode != http.StatusOK {
		return account.Token{}, errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "sign in")
	}

	defer resp.Body.Close()
	var tok account.Token
	if err := json.NewDecoder(resp.Body).Decode(&tok); err != nil {
		return tok, errors.Wrap(err, "sign in")
	}
	return tok, nil
}

// SignOut a connected account.
func (s *Service) SignOut(tokenID gulid.ID, username string) error {
	raw, err := json.Marshal(map[string]string{
		"token":    tokenID.String(),
		"username": username,
	})
	if err != nil {
		return errors.Wrap(err, "sign out")
	}
	resp, err := http.Post(s.url+"/signout", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return errors.Wrap(err, "sign out")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "sign out")
	}
	return nil
}

// Unsubscribe an existing account.
func (s *Service) Unsubscribe(username string, password string) error {
	raw, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return errors.Wrap(err, "unsubscribe")
	}
	resp, err := http.Post(s.url+"/unsubscribe", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return errors.Wrap(err, "unsubscribe")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "unsubscribe")
	}
	return nil
}
