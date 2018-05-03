package storage

import (
	"github.com/elojah/game_01"
)

// Domain converts a storage user into a domain user.
func (u *Account) Domain(username string) (game.Account, error) {
	return game.Account{
		Username: username,
		ID:       game.ID(u.ID),
		Password: u.Password,
	}, nil
}

// NewAccount converts a domain account into a storage account.
func NewAccount(account game.Account) *Account {
	return &Account{
		ID:       [16]byte(account.ID),
		Password: account.Password,
	}
}
