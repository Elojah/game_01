package auth

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/account"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

type Client struct {
	address  string
	username string
	password string

	Token gulid.ID
}

func (c *Client) Close() error {
	return c.SignOut()
}

// Dial initialize a client and assigns a new Token.
func (c *Client) Dial(cfg Config) error {
	c.address = cfg.Address
	c.username = cfg.Username
	c.password = cfg.Password

	// ignore certificate validity
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // nolint: gosec

	if err := c.Subscribe(); err != nil {
		// ignore error
		_ = err
	}
	token, err := c.SignIn()
	if err != nil {
		return err
	}

	c.Token = token.ID

	return nil
}

// Subscribe a new account through HTTPS.
func (c *Client) Subscribe() error {
	raw, err := json.Marshal(map[string]string{
		"username": c.username,
		"password": c.password,
	})
	if err != nil {
		return errors.Wrap(err, "create account")
	}
	resp, err := http.Post(c.address+"/subscribe", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return errors.Wrap(err, "create account")
	}
	_ = resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "create account")
	}
	return nil
}

// SignIn sign to an already created account through HTTPS and returns corresponding token.
func (c *Client) SignIn() (account.Token, error) {
	raw, err := json.Marshal(map[string]string{
		"username": c.username,
		"password": c.password,
	})
	if err != nil {
		return account.Token{}, errors.Wrap(err, "sign in")
	}
	resp, err := http.Post(c.address+"/signin", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return account.Token{}, errors.Wrap(err, "sign in")
	}
	if resp.StatusCode != http.StatusOK {
		// return account.Token{}, errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "sign in")
	}

	defer resp.Body.Close()
	var tok account.Token
	if err := json.NewDecoder(resp.Body).Decode(&tok); err != nil {
		return tok, errors.Wrap(err, "sign in")
	}
	return tok, nil
}

// SignOut a connected account.
func (c *Client) SignOut() error {
	raw, err := json.Marshal(map[string]string{
		"token":    c.Token.String(),
		"username": c.username,
	})
	if err != nil {
		return errors.Wrap(err, "sign out")
	}
	resp, err := http.Post(c.address+"/signout", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return errors.Wrap(err, "sign out")
	}
	_ = resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "sign out")
	}
	return nil
}

// Unsubscribe an existing account.
func (c *Client) Unsubscribe() error {
	raw, err := json.Marshal(map[string]string{
		"username": c.username,
		"password": c.password,
	})
	if err != nil {
		return errors.Wrap(err, "unsubscribe")
	}
	resp, err := http.Post(c.address+"/unsubscribe", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return errors.Wrap(err, "unsubscribe")
	}
	_ = resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "unsubscribe")
	}
	return nil
}
