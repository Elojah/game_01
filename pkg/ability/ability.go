package ability

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// Store is the communication interface for abilities.
type Store interface {
	SetAbility(A, ulid.ID) error
	GetAbility(ulid.ID, ulid.ID) (A, error)
	ListAbility(entityID ulid.ID) ([]A, error)
}
