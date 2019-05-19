package infra

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// Sync represents a core running machine.
type Sync struct {
	ID ulid.ID
}

// SyncStore contains basic operations for infra sync object.
type SyncStore interface {
	InsertSync(Sync) error
	FetchRandomSync() (Sync, error)
	RemoveSync(ulid.ID) error
}
