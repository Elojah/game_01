package game

type Token struct {
	ID          ID
	Permissions map[ID]Right

	IP      *net.UDPAddr
	Account ID
}

type PermissionSubset struct {
	IDs []ID
}

type TokenService interface {
	GetToken(ID) (Token, error)
	AddTokenPermission(ID, Right) error
	UpdateTokenPermission(PermissionSubset, Right) error
	DeleteTokenPermission(PermissionSubset) error
}
