package ability

import gulid "github.com/elojah/game_01/pkg/ulid"

// StarterStore contains basic operations for ability starter object.
type StarterStore interface {
	UpsertStarter(Starter) error
	FetchStarter(gulid.ID) (Starter, error)
}
