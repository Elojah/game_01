package entity

import "github.com/elojah/game_01/pkg/ulid"

// Permission represents a links between 2 objects (token/identities/etc.).
type Permission struct {
	ID     ulid.ID
	Source string
	Target string
	Value  int
}

// PermissionMapper defines Permission operations.
type PermissionMapper interface {
	SetPermission(Permission) error
	GetPermission(PermissionSubset) (Permission, error)
	ListPermission(PermissionSubset) ([]Permission, error)
	DelPermission(PermissionSubset) error
}

// PermissionSubset is the subset to retrieve a Permission.
type PermissionSubset struct {
	Source string
	Target string
}
