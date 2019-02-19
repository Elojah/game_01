package ability

import gulid "github.com/elojah/game_01/pkg/ulid"

// StarterStore interface starter abilities for an entity template.
type StarterStore interface {
	SetStarter(Starter) error
	GetStarter(gulid.ID) (Starter, error)
}