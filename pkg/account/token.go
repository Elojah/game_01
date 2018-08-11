package account

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// TokenMapper is the service gate for Token resource.
type TokenMapper interface {
	SetToken(Token) error
	GetToken(TokenSubset) (Token, error)
	DelToken(TokenSubset) error
}

// TokenSubset retrieves a token per ID.
type TokenSubset struct {
	ID ulid.ID
}

// TokenHCMapper is the service gate for Token health check.
type TokenHCMapper interface {
	SetTokenHC(ulid.ID, int64) error
	ListTokenHC(TokenHCSubset) ([]ulid.ID, error)
}

// TokenHCSubset retrieves token healthchecks based on last tick.
type TokenHCSubset struct {
	MaxTS int64
}
