package account

import (
	"net"

	game "github.com/elojah/game_01"
)

// Token represents a user connection. Creation is made by secure https only.
type Token struct {
	ID      game.ID      `json:"ID"`
	IP      *net.UDPAddr `json:"-"`
	Account game.ID      `json:"-"`
	Ping    uint64       `json:"-"`
}

// TokenMapper is the service gate for Token resource.
type TokenMapper interface {
	SetToken(Token) error
	GetToken(game.ID) (Token, error)
}

// TokenHCMapper is the service gate for Token health check.
type TokenHCMapper interface {
	SetTokenHC(game.ID, int64) error
	ListTokenHC(TokenHCSubset) ([]game.ID, error)
}

// TokenHCSubset retrieves token healthchecks based on last tick.
type TokenHCSubset struct {
	MaxTS int64
}
