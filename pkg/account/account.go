package account

import (
	"net"

	"github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

// Address registered to send data back.
type Address struct {
	IP   string
	Port uint64
}

// Check if address is valid.
func (ad Address) Check() error {
	ip := net.ParseIP(ad.IP)
	if ip.To4() == nil {
		return errors.ErrInvalidIP{IP: ad.IP}
	}
	if !(ad.Port > 1 && ad.Port < 99999) {
		return errors.ErrInvalidPort{Port: ad.Port}
	}
	return nil
}

// Store contains basic operations for account A.
type Store interface {
	Upsert(A) error
	Fetch(string) (A, error)
	Remove(string) error
}

// App contains account stores and applications.
type App interface {
	Store
	TokenStore
	TokenHCStore

	CreateToken(A, Address) (Token, error)
	FetchTokenFromAddr(ulid.ID, string) (Token, error)
	DisconnectToken(ulid.ID) error
}
