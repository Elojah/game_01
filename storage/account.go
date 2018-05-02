package storage

import (
	"github.com/elojah/game_01"
)

// Domain converts a storage user into a domain user.
func (u *Account) Domain() (game.Account, error) {
	return game.Account{
		Username: u.Username,
		Password: u.Password,
	}, nil
}
