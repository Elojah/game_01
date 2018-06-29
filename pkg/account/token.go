package account

import (
	"net"

	"github.com/elojah/game_01/pkg/ulid"
)

// Token represents a user connection. Creation is made by secure https only.
type Token struct {
	ID       ulid.ID      `json:"ID"`
	IP       *net.UDPAddr `json:"-"`
	Account  ulid.ID      `json:"-"`
	Ping     uint64       `json:"-"`
	CorePool ulid.ID      `json:"-"`
	SyncPool ulid.ID      `json:"-"`
}

// TokenMapper is the service gate for Token resource.
type TokenMapper interface {
	SetToken(Token) error
	GetToken(ulid.ID) (Token, error)
}

// TokenHCMapper is the service gate for Token health check.
type TokenHCMapper interface {
	SetTokenHC(ulid.ID, int64) error
	ListTokenHC(TokenHCSubset) ([]ulid.ID, error)
}

// TokenHCSubset retrieves token healthchecks based on last tick.
type TokenHCSubset struct {
	MaxTS int64
}
