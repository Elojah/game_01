package infra

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// Core represents a core running machine.
type Core struct {
	ID ulid.ID
}

// CoreStore maps core data interactions.
type CoreStore interface {
	GetRandomCore(CoreSubset) (Core, error)
	SetCore(Core) error
	DelCore(CoreSubset) error
}

// CoreSubset retrieves a randomly assigned core.
type CoreSubset struct {
	ID ulid.ID
}
