package entity

import (
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/ulid"
)

// Store is an interface for E object.
type Store interface {
	SetEntity(E, int64) error
	GetEntity(Subset) (E, error)
	DelEntity(Subset) error
}

// Subset is a subset to retrieve one entity.
type Subset struct {
	ID     ulid.ID
	MinTS  int64
	MaxTS  int64
	Cursor uint64
	Count  int64
}

// Service represents entity usecases.
type Service interface {
	Disconnect(id ulid.ID, tok account.Token) error
}
