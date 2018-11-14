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
	EntityStore entity.Store
	SectorStore sector.Store
	sector.EntitiesStore

	logger   zerolog.Logger
	id       gulid.ID
	entityID gulid.ID

	ticker   *time.Ticker
	callback func(entity.E)
}

// NewRecurrer returns a new recurrer which sends entity data associated to id to addr, tick times per second.
func NewRecurrer(rec infra.Recurrer, tick uint32, callback func(entity.E)) *Recurrer {
	return &Recurrer{
		logger:   log.With().Str("recurrer", rec.TokenID.String()).Logger(),
		id:       rec.TokenID,
		entityID: rec.EntityID,
		callback: callback,

		ticker: time.NewTicker(time.Second / time.Duration(tick)),
	}
}

// Close close the tick sender.
func (r *Recurrer) Close() error {
	r.logger.Info().Msg("close recurrer")
	r.ticker.Stop()
	return nil
}

// Run starts to read the ticker and send entities.
func (r *Recurrer) Run() {
	for t := range r.ticker.C {
		en, err := r.EntityStore.GetEntity(r.entityID, ulid.Timestamp(t))
		if err != nil {
			r.logger.Error().Err(err).Msg("failed to retrieve entity")
			continue
		}
		sector, err := r.SectorStore.GetSector(en.Position.SectorID)
		if err != nil {
			r.logger.Error().Err(err).Msg("failed to retrieve current sector")
			continue
		}
		go r.sendSector(sector.ID, t)
		for _, id := range sector.Exposed {
			go r.sendSector(id, t)
		}
	}
}

func (r *Recurrer) sendSector(sectorID gulid.ID, t time.Time) {
	se, err := r.GetEntities(sectorID)
	if err != nil {
		r.logger.Error().Err(err).Str("id", sectorID.String()).Msg("failed to retrieve sector")
		return
	}
	for _, entityID := range se.EntityIDs {
		go r.sendEntity(entityID, t)
	}
}

func (r *Recurrer) sendEntity(entityID gulid.ID, t time.Time) {
	// TODO Use t-token ping instead of t.
	e, err := r.EntityStore.GetEntity(entityID, ulid.Timestamp(t))
	if err != nil {
		r.logger.Error().Err(err).Str("id", entityID.String()).Msg("failed to retrieve entity")
		return
	}
	r.callback(e)
}
