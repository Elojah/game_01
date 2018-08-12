package sector

import "github.com/elojah/game_01/pkg/ulid"

// EntitiesStore set or get sector entities. Can also add or remove individual entity to sector.
type EntitiesStore interface {
	GetEntities(EntitiesSubset) (Entities, error)
	AddEntityToSector(ulid.ID, ulid.ID) error
	RemoveEntityFromSector(ulid.ID, ulid.ID) error
}

// EntitiesSubset retrieves one Entities per sector ulid.ID only.
type EntitiesSubset struct {
	SectorID ulid.ID
}
