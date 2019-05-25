package main

import (
	"time"

	"github.com/oklog/ulid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/sector"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// Recurrer retrieves entity data associated to pc id and send it at regular ticks.
type Recurrer struct {
	EntityStore         entity.Store
	SectorStore         sector.Store
	SectorEntitiesStore sector.EntitiesStore

	logger   zerolog.Logger
	id       gulid.ID
	entityID gulid.ID

	ticker    *time.Ticker
	batchSize uint32
	callback  func(entity.DTO)
}

// NewRecurrer returns a new recurrer which sends entity data associated to id to addr, tick times per second.
func NewRecurrer(rec infra.Recurrer, tick uint32, batchSize uint32, callback func(entity.DTO)) *Recurrer {
	return &Recurrer{
		logger:   log.With().Str("recurrer", rec.TokenID.String()).Logger(),
		id:       rec.TokenID,
		entityID: rec.EntityID,
		callback: callback,

		ticker:    time.NewTicker(time.Second / time.Duration(tick)),
		batchSize: batchSize,
	}
}

// Close close the tick sender.
func (r *Recurrer) Close() error {
	r.logger.Info().Msg("close")
	r.ticker.Stop()
	return nil
}

// Run starts to read the ticker and send entities.
func (r *Recurrer) Run() {
	for t := range r.ticker.C {
		e, err := r.EntityStore.FetchEntity(r.entityID, ulid.Timestamp(t))
		if err != nil {
			r.logger.Error().Err(err).Msg("run")
			continue
		}
		sector, err := r.SectorStore.FetchSector(e.Position.SectorID)
		if err != nil {
			r.logger.Error().Err(err).Msg("run")
			continue
		}
		r.sendSector(sector.ID, t)
		for _, id := range sector.Exposed {
			r.sendSector(id, t)
		}
	}
}

func (r *Recurrer) sendSector(sectorID gulid.ID, t time.Time) {
	se, err := r.SectorEntitiesStore.FetchEntities(sectorID)
	if err != nil {
		r.logger.Error().Err(err).Str("sector", sectorID.String()).Msg("send sector")
		return
	}

	dto := entity.DTO{Entities: make([]entity.E, r.batchSize)}
	var i uint32
	for _, entityID := range se.EntityIDs {
		var err error
		dto.Entities[i], err = r.EntityStore.FetchEntity(entityID, ulid.Timestamp(t))
		if err != nil {
			// Soft error
			r.logger.Error().Err(err).Str("entity", entityID.String()).Msg("send entity")
			continue
		}
		i++
		// if batch size is complete, send it and reset current counter
		if i == r.batchSize {
			r.callback(dto)
			i = 0
		}
	}
	// if there is still unsend entities, send them
	if i != 0 {
		// Reduce entities size to remove empty or previous data.
		dto.Entities = append(dto.Entities[:i])
		r.callback(dto)
	}
}
