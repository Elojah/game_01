package storage

import (
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/ulid"
)

// Domain converts a storage user into a domain user.
func (a *Account) Domain(username string) (account.A, error) {
	return account.A{
		Username: username,
		ID:       ulid.ID(a.ID),
		Password: a.Password,
	}, nil
}

// NewAccount converts a domain account into a storage account.
func NewAccount(a account.A) *Account {
	return &Account{
		ID:       [16]byte(a.ID),
		Password: a.Password,
	}
}
