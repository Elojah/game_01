package main

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

// Recurrer retrieves entity data associated to pc id and send it at regular ticks.
type Recurrer struct {
	game.EntityMapper
	game.SectorEntitiesMapper
	game.SectorMapper

	logger zerolog.Logger
	id     game.ID

	ticker   *time.Ticker
	callback func(raw []byte)
}

// NewRecurrer returns a new recurrer which sends entity data associated to id to addr, tick times per second.
func NewRecurrer(id game.ID, tick uint32, callback func(raw []byte)) *Recurrer {
	return &Recurrer{
		logger:   log.With().Str("recurrer", id.String()).Logger(),
		id:       id,
		callback: callback,

		ticker: time.NewTicker(time.Second / time.Duration(tick)),
	}
}

// Close close the tick sender.
func (r *Recurrer) Close() {
	r.ticker.Stop()
}

func (r *Recurrer) tick() {
	for t := range r.ticker.C {
		entity, err := r.GetEntity(game.EntitySubset{Key: r.id.String(), MaxTS: t.UnixNano()})
		if err != nil {
			r.logger.Error().Err(err).Msg("failed to retrieve entity")
			continue
		}
		sector, err := r.GetSector(game.SectorSubset{ID: entity.Position.SectorID})
		if err != nil {
			r.logger.Error().Err(err).Msg("failed to retrieve current sector")
			continue
		}
		go r.sendSector(sector.ID, t)
		for _, exps := range sector.ExitPoints {
			for _, exp := range exps {
				go r.sendSector(exp.SectorID, t)
			}
		}
	}
}

func (r *Recurrer) sendSector(sectorID game.ID, t time.Time) {
	se, err := r.GetSectorEntities(game.SectorEntitiesSubset{SectorID: sectorID})
	if err != nil {
		r.logger.Error().Err(err).Str("id", sectorID.String()).Msg("failed to retrieve sector")
		return
	}
	for _, entityID := range se.EntityIDs {
		go r.sendEntity(entityID, t)
	}
}

func (r *Recurrer) sendEntity(entityID game.ID, t time.Time) {
	entity, err := r.GetEntity(game.EntitySubset{Key: entityID.String(), MaxTS: t.UnixNano()})
	if err != nil {
		r.logger.Error().Err(err).Str("id", entityID.String()).Msg("failed to retrieve entity")
		return
	}
	raw, err := storage.NewEntity(entity).Marshal(nil)
	if err != nil {
		r.logger.Error().Err(err).Msg("failed to retrieve marshal entity")
		return
	}
	go r.callback(raw)
}

// Start init a new goroutine to tick regularly.
func (r *Recurrer) Start() {
	go r.tick()
}
