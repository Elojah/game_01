package sector

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// Starter is a starter sector
type Starter struct {
	S
	ID ulid.ID
}

// StarterService interfaces starter data interactions.
type StarterService interface {
	GetStarter(StarterSubset) (Starter, error)
	GetRandomStarter(StarterSubset) (Starter, error)
	SetStarter(Starter) error
}

// StarterSubset retrieves a Starter by ID.
type StarterSubset struct {
	ID ulid.ID
}
