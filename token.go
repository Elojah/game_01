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

type TokenService interface {
	ListToken(TokenSubset) ([]Token, error)
	AddTokenPermission(TokenSubset, Permissions) error
	UpdateTokenPermission(TokenSubset, PermissionSubset, PermissionPatch) error
	DeleteTokenPermission(TokenSubset, PermissionSubset) error
}
