package game

import (
	"net"
)

// Token represents a user connection. Creation is made by secure https only.
type Token struct {
	ID      ID           `json:"ID"`
	IP      *net.UDPAddr `json:"-"`
	Account ID           `json:"-"`
	Ping    uint64       `json:"-"`
}

// TokenMapper is the service gate for Token resource.
type TokenMapper interface {
	SetToken(Token) error
	GetToken(ID) (Token, error)
}

// TokenHCMapper is the service gate for Token health check.
type TokenHCMapper interface {
	SetTokenHC(ID, int64) error
	ListTokenHC(TokenHCSubset) ([]ID, error)
}

// TokenHCSubset retrieves token healthchecks based on last tick.
type TokenHCSubset struct {
	MaxTS int64
}
