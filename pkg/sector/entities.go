package sector

import "github.com/elojah/game_01/pkg/ulid"

// EntitiesStore set or get sector entities. Can also add or remove individual entity to sector.
type EntitiesStore interface {
	GetEntities(ulid.ID) (Entities, error)
	AddEntityToSector(ulid.ID, ulid.ID) error
	RemoveEntityFromSector(ulid.ID, ulid.ID) error
}
