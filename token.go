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
