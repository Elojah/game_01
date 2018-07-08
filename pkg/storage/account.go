package storage

import (
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/ulid"
)

// Domain converts a storage user into a domain user.
func (a *Account) Domain(username string) (account.A, error) {
	return account.A{
		ID:       ulid.ID(a.ID),
		Username: username,
		Password: a.Password,
		Token:    ulid.ID(a.Token),
	}, nil
}

// NewAccount converts a domain account into a storage account.
func NewAccount(a account.A) *Account {
	return &Account{
		ID:       [16]byte(a.ID),
		Username: a.Username,
		Password: a.Password,
		Token:    [16]byte(a.Token),
	}
}
