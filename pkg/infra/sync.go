package infra

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// Sync represents a core running machine.
type Sync struct {
	ID ulid.ID
}

// SyncStore maps sync data interactions.
type SyncStore interface {
	GetRandomSync() (Sync, error)
	SetSync(Sync) error
	DelSync(ulid.ID) error
}
