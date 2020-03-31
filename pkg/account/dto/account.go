package dto

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

// Check check if username and passwords are valid.
func (a Account) Check() error {
	lu := len(a.Username)
	lp := len(a.Password)
	if lu < 4 || lu > 25 ||
		lp < 7 || lp > 50 ||
		strings.IndexFunc(a.Username, func(r rune) bool {
			return (r < 'A' || r > 'z') && (r < '0' || r > '9') && (r != '_')
		}) != -1 ||
		strings.IndexFunc(a.Password, func(r rune) bool {
			return (r < 'A' || r > 'z') && (r < '0' || r > '9') && (r != '_')
		}) != -1 {
		return errors.New("invalid account")
	}
	return nil
}

// SignInAccount represents the payload to connect an account via /signin.
type SignInAccount struct {
	Account
	account.Address
}

// Check check if account and address are valid.
func (a SignInAccount) Check() error {
	if err := a.Account.Check(); err != nil {
		return err
	}
	if err := a.Address.Check(); err != nil {
		return err
	}
	return nil
}

// SignoutAccount represents the payload to disconnect an account via /signout.
type SignoutAccount struct {
	Username string
	Token    ulid.ID
}
