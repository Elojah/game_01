package account

import "github.com/elojah/game_01/pkg/ulid"

// A A represents an user account.
type A struct {
	ID       ulid.ID
	Username string
	Password string `json:"-"`
}

// Subset is the subset to retrieve an account.
type Subset struct {
	Username string
	Password string
}

// Mapper wraps account interactions.
type Mapper interface {
	SetAccount(A) error
	GetAccount(Subset) (A, error)
}
