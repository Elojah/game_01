package main

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/sector"
)

// Recurrer retrieves entity data associated to pc id and send it at regular ticks.
type Recurrer struct {
	EntityMapper entity.Mapper
	SectorMapper sector.Mapper
	sector.EntitiesMapper

	logger   zerolog.Logger
	id       game.ID
	entityID game.ID

	ticker   *time.Ticker
	callback func(entity.E)
}

// NewRecurrer returns a new recurrer which sends entity data associated to id to addr, tick times per second.
func NewRecurrer(rec event.Recurrer, tick uint32, callback func(entity.E)) *Recurrer {
	return &Recurrer{
		logger:   log.With().Str("recurrer", rec.ID.String()).Logger(),
		id:       rec.ID,
		entityID: rec.EntityID,
		callback: callback,

		ticker: time.NewTicker(time.Second / time.Duration(tick)),
	}
}

// Close close the tick sender.
func (r *Recurrer) Close() {
	r.ticker.Stop()
}

// Start starts to read the ticker and send entities.
func (r *Recurrer) Start() {
	for t := range r.ticker.C {
		entity, err := r.EntityMapper.GetEntity(entity.Subset{Key: r.entityID.String(), MaxTS: t.UnixNano()})
		if err != nil {
			r.logger.Error().Err(err).Msg("failed to retrieve entity")
			continue
		}
		sector, err := r.SectorMapper.GetSector(sector.Subset{ID: entity.Position.SectorID})
		if err != nil {
			r.logger.Error().Err(err).Msg("failed to retrieve current sector")
			continue
		}
		go r.sendSector(sector.ID, t)
		for id := range sector.Adjacents() {
			go r.sendSector(id, t)
		}
	}
}

func (r *Recurrer) sendSector(sectorID game.ID, t time.Time) {
	se, err := r.GetEntities(sector.EntitiesSubset{SectorID: sectorID})
	if err != nil {
		r.logger.Error().Err(err).Str("id", sectorID.String()).Msg("failed to retrieve sector")
		return
	}
	for _, entityID := range se.EntityIDs {
		go r.sendEntity(entityID, t)
	}
}

func (r *Recurrer) sendEntity(entityID game.ID, t time.Time) {
	// TODO Use token ping instead of now()
	entity, err := r.EntityMapper.GetEntity(entity.Subset{Key: entityID.String(), MaxTS: t.UnixNano()})
	if err != nil {
		r.logger.Error().Err(err).Str("id", entityID.String()).Msg("failed to retrieve entity")
		return
	}
	r.callback(entity)
}
