package ability

import (
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// Store is the communication interface for abilities.
type Store interface {
	SetAbility(A, gulid.ID) error
	GetAbility(gulid.ID, gulid.ID) (A, error)
	ListAbility(entityID gulid.ID) ([]A, error)
}

// Service wraps ability helpers.
type Service interface {
	SetStarterAbilities(gulid.ID, gulid.ID) error
	CopyAbilities(gulid.ID, gulid.ID) error
}
