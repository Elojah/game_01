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

// TokenService is the service gate for Token resource.
type TokenService interface {
	SetToken(Token) error
	GetToken(ID) (Token, error)
}
