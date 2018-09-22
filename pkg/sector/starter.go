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
	GetRandomStarter() (Starter, error)
	SetStarter(Starter) error
	DelStarter(ulid.ID) error
}
