package game

type Permissions map[ID]Right

type PermissionSubset struct {
	IDs []ID
}

type PermissionPatch struct {
	Right Right
}
