package entity

import gulid "github.com/elojah/game_01/pkg/ulid"

// Permission represents a links between 2 objects (token/identities/etc.).
type Permission struct {
	ID     gulid.ID
	Source string
	Target string
	Value  int
}

// PermissionStore defines Permission operations.
type PermissionStore interface {
	SetPermission(Permission) error
	GetPermission(string, string) (Permission, error)
	ListPermission(string) ([]Permission, error)
	DelPermission(string, string) error
}

// PermissionService wraps permission usecases for entities.
type PermissionService interface {
	CheckPermission(gulid.ID, gulid.ID) error
}
