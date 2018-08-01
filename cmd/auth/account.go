package main

import (
	"errors"
	"strings"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/ulid"
)

// Account represents the payload to create a new account via subscribe.
type Account struct {
	Username string
	Password string
}

// Domain transforms a dto account into a domain account.
func (a Account) Domain() account.A {
	return account.A{
		Username: a.Username,
		Password: a.Password,
	}
}

// Check check if username and passwords are valid.
func (a Account) Check() error {
	lu := len(a.Username)
	lp := len(a.Password)
	if lu < 4 || lu > 25 ||
		lp < 7 || lp > 50 ||
		strings.IndexFunc(a.Username, func(r rune) bool {
			return r < 'A' || r > 'z'
		}) != -1 ||
		strings.IndexFunc(a.Password, func(r rune) bool {
			return r < 'A' || r > 'z'
		}) != -1 {
		return errors.New("invalid account")
	}
	return nil
}

// SignoutAccount represents the payload to disconnect an account via /signout.
type SignoutAccount struct {
	Username string
	Token    ulid.ID
}
