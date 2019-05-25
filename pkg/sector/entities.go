package sector

import "github.com/elojah/game_01/pkg/ulid"

// Store contains basic operations for entities sector object.
type EntitiesStore interface {
	FetchEntities(ulid.ID) (Entities, error)
	AddEntityToSector(ulid.ID, ulid.ID) error
	RemoveEntityFromSector(ulid.ID, ulid.ID) error
}
