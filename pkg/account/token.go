package account

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// TokenStore is the service gate for Token resource.
type TokenStore interface {
	UpsertToken(Token) error
	FetchToken(ulid.ID) (Token, error)
	RemoveToken(ulid.ID) error
}

// TokenHCStore is the service gate for Token health check.
type TokenHCStore interface {
	UpsertTokenHC(ulid.ID, uint64) error
	ListTokenHC(uint64) ([]ulid.ID, error)
}
