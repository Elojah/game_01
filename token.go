package game

import (
	"net"
)

// Token represents a user connection. Creation is made by secure https only.
type Token struct {
	ID          ID
	Permissions map[ID]Right

	IP      *net.UDPAddr
	Account ID
}

// TokenSubset represents a subset of Token resources.
type TokenSubset struct {
	IDs []ID
}

// TokenService is the service gate for Token resource.
type TokenService interface {
	GetToken(ID) (Token, error)
	// AddTokenPermission(TokenSubset, Permissions) error
	// UpdateTokenPermission(TokenSubset, PermissionSubset, PermissionPatch) error
	// DeleteTokenPermission(TokenSubset, PermissionSubset) error
}
