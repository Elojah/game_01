package ability

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// Store is the communication interface for abilities.
type Store interface {
	SetAbility(A, ulid.ID) error
	GetAbility(Subset) (A, error)
	ListAbility(Subset) ([]A, error)
}

// Subset retrieves per EntityID+ID or list per EntityID.
type Subset struct {
	ID       ulid.ID
	EntityID ulid.ID
}
