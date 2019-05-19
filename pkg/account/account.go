package account

import "github.com/elojah/game_01/pkg/ulid"

// Store contains basic operations fo account A.
type Store interface {
	InsertAccount(A) error
	FetchAccount(string) (A, error)
	RemoveAccount(string) error
}

// App contains account stores and applications.
type App interface {
	Store
	TokenStore
	TokenHCStore

	CreateToken(A, string) (Token, error)
	FetchTokenFromAddr(ulid.ID, string) (Token, error)
	DisconnectToken(ulid.ID) error
}
