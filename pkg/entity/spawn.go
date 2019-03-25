package entity

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// SpawnStore wraps spawn db common operations.
type SpawnStore interface {
	GetSpawn(ulid.ID) (Spawn, error)
	SetSpawn(Spawn) error
	DelSpawn(ulid.ID) error
}
