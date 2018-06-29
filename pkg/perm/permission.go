package perm

import "github.com/elojah/game_01/pkg/ulid"

// P represents a links between 2 objects (token/identities/etc.).
type P struct {
	ID     ulid.ID
	Source string
	Target string
	Value  int
}

// Mapper defines P operations.
type Mapper interface {
	SetPermission(P) error
	GetPermission(Subset) (P, error)
	DelPermission(Subset) error
}

// Subset is the subset to retrieve a P.
type Subset struct {
	Source string
	Target string
}
