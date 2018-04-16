package game

// Permissions represents permissions access as stored in database as a map id->right.
type Permissions map[ID]Right

// PermissionSubset is a subset of permissions for update/delete operations.
type PermissionSubset struct {
	IDs []ID
}

// PermissionPatch is a patch of permissions for update operation.
type PermissionPatch struct {
	Right Right
}
