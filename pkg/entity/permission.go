package entity

import gulid "github.com/elojah/game_01/pkg/ulid"

// Permission represents a links between 2 objects (token/identities/etc.).
type Permission struct {
	ID     gulid.ID
	Source string
	Target string
	Value  int
}

// PermissionStore contains basic operations for entity permission object.
type PermissionStore interface {
	UpsertPermission(Permission) error
	FetchPermission(string, string) (Permission, error)
	ListPermission(string) ([]Permission, error)
	RemovePermission(string, string) error
}
