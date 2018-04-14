package game

type Token struct {
	ID          ID
	Permissions map[ID]Right

	IP      *net.UDPAddr
	Account ID
}

type TokenSubset struct {
	IDs []ID
}

type PermissionSubset struct {
	IDs []ID
}

type TokenService interface {
	GetToken(ID) (Token, error)
	AddTokenPermission(TokenSubset, PermissionSubset, Right) error
	UpdateTokenPermission(TokenSubset, PermissionSubset, Right) error
	DeleteTokenPermission(TokenSubset, PermissionSubset) error
}
