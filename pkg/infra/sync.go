package infra

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// Sync represents a core running machine.
type Sync struct {
	ID ulid.ID
}

// SyncMapper maps sync data interactions.
type SyncMapper interface {
	GetRandomSync(SyncSubset) (Sync, error)
	SetSync(Sync) error
	DelSync(SyncSubset) error
}

// SyncSubset retrieves a randomly assigned core.
type SyncSubset struct {
	ID ulid.ID
}
