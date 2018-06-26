package storage

import (
	"github.com/elojah/game_01"
	"github.com/elojah/game_01/pkg/sector"
)

// Domain converts a storage sector entities into a domain sector entities.
func (se *SectorEntities) Domain() sector.Entities {
	entityIDs := make([]game.ID, len(se.EntityIDs))
	for i, entity := range se.EntityIDs {
		entityIDs[i] = game.ID(entity)
	}
	return sector.Entities{
		SectorID:  game.ID(se.SectorID),
		EntityIDs: entityIDs,
	}
}

// NewSectorEntities converts a domain se into a storage se.
func NewSectorEntities(se sector.Entities) *SectorEntities {
	entityIDs := make([][16]byte, len(se.EntityIDs))
	for i, entity := range se.EntityIDs {
		entityIDs[i] = [16]byte(entity)
	}
	return &SectorEntities{
		SectorID: [16]byte(se.SectorID),
	}
}
