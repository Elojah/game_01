package storage

import (
	"github.com/elojah/game_01"
)

// Domain converts a storage sector entities into a domain sector entities.
func (se *SectorEntities) Domain() game.SectorEntities {
	entityIDs := make([]game.ID, len(se.EntityIDs))
	for i, entity := range se.EntityIDs {
		entityIDs[i] = game.ID(entity)
	}
	return game.SectorEntities{
		SectorID:  game.ID(se.SectorID),
		EntityIDs: entityIDs,
	}
}

// NewSectorEntities converts a domain se into a storage se.
func NewSectorEntities(se game.SectorEntities) *SectorEntities {
	entityIDs := make([][16]byte, len(se.EntityIDs))
	for i, entity := range se.EntityIDs {
		entityIDs[i] = [16]byte(entity)
	}
	return &SectorEntities{
		SectorID: [16]byte(se.SectorID),
	}
}
