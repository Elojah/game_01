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
	GetRandomCore() (Core, error)
	SetCore(Core) error
	DelCore(ulid.ID) error
}
