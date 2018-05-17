package game

import (
	"net"
)

// Token represents a user connection. Creation is made by secure https only.
type Token struct {
	ID      ID           `json:"ID"`
	IP      *net.UDPAddr `json:"-"`
	Account ID           `json:"-"`
}

// TokenSubset represents a subset of Token resources.
type TokenSubset struct {
	IDs []ID
}

// TokenService is the service gate for Token resource.
type TokenService interface {
	SetToken(Token) error
	GetToken(ID) (Token, error)
}
