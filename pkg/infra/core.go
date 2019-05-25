package infra

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// Core represents a core running machine.
type Core struct {
	ID ulid.ID
}

// CoreStore contains basic operations for infra core object.
type CoreStore interface {
	UpsertCore(Core) error
	FetchRandomCore() (Core, error)
	RemoveCore(ulid.ID) error
}
