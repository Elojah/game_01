package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	gulid "github.com/elojah/game_01/pkg/ulid"
	"github.com/pkg/errors"
)

// AccountService wraps account helpers.
type AccountService struct {
	auth string
}

// Create a new account through HTTPS.
func (s *AccountService) Create(username string, password string) error {
	raw, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return errors.Wrap(err, "create account")
	}
	resp, err := http.Post(s.auth+"/subscribe", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return errors.Wrap(err, "create account")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "create account")
	}

	return nil
}

// SignIn sign to an already created account through HTTPS and returns corresponding token.
func (s *AccountService) SignIn(username string, password string) (account.Token, error) {
	raw, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return account.Token{}, errors.Wrap(err, "sign in")
	}
	resp, err := http.Post(s.auth+"/signin", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return account.Token{}, errors.Wrap(err, "sign in")
	}
	if resp.StatusCode != http.StatusOK {
		return account.Token{}, errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "sign in")
	}

	defer resp.Body.Close()
	var tok account.Token
	if err := json.NewDecoder(resp.Body).Decode(&tok); err != nil {
		return tok, err
	}
	return tok, nil
}

// CreatePC creates a new PC for token.
func (s *AccountService) CreatePC(tokenID gulid.ID, pcName string, pcType string) error {
	raw, err := json.Marshal(map[string]string{
		"token": tokenID.String(),
		"name":  pcName,
		"type":  pcType,
	})
	if err != nil {
		return errors.Wrap(err, "create pc")
	}
	resp, err := http.Post(s.auth+"/pc/create", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return errors.Wrap(err, "create pc")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "create pc")
	}
	return nil
}

// ListPC list all pcs corresponding to token.
func (s *AccountService) ListPC(tokenID gulid.ID) ([]entity.PC, error) {
	raw, err := json.Marshal(map[string]string{
		"token": tokenID.String(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "list pc")
	}
	resp, err := http.Post(s.auth+"/pc/list", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return nil, errors.Wrap(err, "list pc")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "list pc")
	}

	var pcs []entity.PC
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&pcs); err != nil {
		return nil, errors.Wrap(err, "list pc")
	}
	return pcs, nil
}
