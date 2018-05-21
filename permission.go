package game

// Permission represents a links between 2 objects (token/identities/etc.).
type Permission struct {
	ID     ID
	Source string
	Target string
	Value  int
}

// PermissionMapper defines Permission operations.
type PermissionMapper interface {
	SetPermission(Permission) error
	GetPermission(PermissionSubset) (Permission, error)
}

// PermissionSubset is the subset to retrieve a Permission.
type PermissionSubset struct {
	Source string
	Target string
}
