package sector

import "github.com/elojah/game_01/pkg/ulid"

// EntitiesMapper set or get sector entities. Can also add or remove individual entity to sector.
type EntitiesMapper interface {
	GetEntities(EntitiesSubset) (Entities, error)
	AddEntityToSector(ulid.ID, ulid.ID) error
	RemoveEntityToSector(ulid.ID, ulid.ID) error
}

// EntitiesSubset retrieves one Entities per sector ulid.ID only.
type EntitiesSubset struct {
	SectorID ulid.ID
}
