package sector

import (
	game "github.com/elojah/game_01"
)

// Entities represents a entity sector in a world.
type Entities struct {
	SectorID  game.ID
	EntityIDs []game.ID
}

// EntitiesMapper set or get sector entities. Can also add or remove individual entity to sector.
type EntitiesMapper interface {
	GetEntities(EntitiesSubset) (Entities, error)
	AddEntityToSector(game.ID, game.ID) error
	RemoveEntityToSector(game.ID, game.ID) error
}

// EntitiesSubset retrieves one Entities per sector game.ID only.
type EntitiesSubset struct {
	SectorID game.ID
}
