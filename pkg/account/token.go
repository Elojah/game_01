package account

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// TokenStore is the service gate for Token resource.
type TokenStore interface {
	SetToken(Token) error
	GetToken(TokenSubset) (Token, error)
	DelToken(TokenSubset) error
}

// TokenSubset retrieves a token per ID.
type TokenSubset struct {
	ID ulid.ID
}

// TokenHCStore is the service gate for Token health check.
type TokenHCStore interface {
	SetTokenHC(ulid.ID, int64) error
	ListTokenHC(TokenHCSubset) ([]ulid.ID, error)
}

// TokenHCSubset retrieves token healthchecks based on last tick.
type TokenHCSubset struct {
	MaxTS int64
}

// TokenService represents token usecases.
type TokenService interface {
	New(A, string) (Token, error)
	Access(ulid.ID, string) (Token, error)
	Disconnect(ulid.ID) error
}
