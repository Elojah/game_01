package sector

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// Starter is a starter sector
type Starter struct {
	SectorID ulid.ID
}

// StarterStore interfaces starter data interactions.
type StarterStore interface {
	GetRandomStarter(StarterSubset) (Starter, error)
	SetStarter(Starter) error
	DelStarter(StarterSubset) error
}

// StarterSubset retrieves a Starter by ID.
type StarterSubset struct {
	ID ulid.ID
}
