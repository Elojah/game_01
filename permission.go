package game

// Permission represents a links between 2 objects (token/identities/etc.).
type Permission struct {
	ID     ID
	Source ID
	Target ID
	Value  Right
}

// PermissionService defines Permission operations.
type PermissionService interface {
	CreatePermission(Permission) error
	GetPermission(PermissionSubset) (Permission, error)
}

// PermissionSubset is the subset to retrieve a Permission.
type PermissionSubset struct {
	Source ID
	Target ID
}
