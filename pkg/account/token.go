package account

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// TokenStore is the service gate for Token resource.
type TokenStore interface {
	SetToken(Token) error
	GetToken(ulid.ID) (Token, error)
	DelToken(ulid.ID) error
}

// TokenHCStore is the service gate for Token health check.
type TokenHCStore interface {
	SetTokenHC(ulid.ID, int64) error
	ListTokenHC(int64) ([]ulid.ID, error)
}

// TokenService represents token usecases.
type TokenService interface {
	New(A, string) (Token, error)
	Access(ulid.ID, string) (Token, error)
	Disconnect(ulid.ID) error
}
