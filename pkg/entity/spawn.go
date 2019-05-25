package entity

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// SpawnStore contains basic operations for entity spawn object.
type SpawnStore interface {
	UpsertSpawn(Spawn) error
	FetchSpawn(ulid.ID) (Spawn, error)
	RemoveSpawn(ulid.ID) error
}
