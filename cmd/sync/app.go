package main

import (
	nats "github.com/nats-io/go-nats"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/storage"
	"github.com/elojah/mux"
)

type app struct {
	account.TokenMapper

	EntityMapper entity.Mapper

	event.QMapper
	event.QRecurrerMapper
	event.SubscriptionMapper

	sector.EntitiesMapper
	SectorMapper sector.Mapper

	*mux.M

	id game.ID

	sub *event.Subscription

	tickRate  uint32
	recurrers map[game.ID]*Recurrer
}

func (a *app) Dial(c Config) error {
	a.tickRate = c.TickRate
	a.id = c.ID
	return nil
}

func (a *app) Start() {
	logger := log.With().Str("sync", a.id.String()).Logger()

	sub, err := a.SetSubscription(a.id.String(), a.AddRecurrer)
	if err != nil {
		logger.Error().Err(err).Msg("failed to sub")
		return
	}
	a.sub = sub
	a.recurrers = make(map[game.ID]*Recurrer)
}

func (a *app) Close() {
	a.sub.Unsubscribe()
	for _, r := range a.recurrers {
		r.Close()
	}
}

func (a *app) AddRecurrer(msg *nats.Msg) {
	logger := log.With().Str("sync", a.id.String()).Logger()

	var recurrerS storage.Recurrer
	if _, err := recurrerS.Unmarshal(msg.Data); err != nil {
		logger.Error().Err(err).Msg("failed to unmarshal recurrer")
		return
	}
	recurrer := recurrerS.Domain()

	if recurrer.Action == event.CloseRec {
		a.recurrers[recurrer.ID].Close()
		delete(a.recurrers, recurrer.ID)
		return
	}

	token, err := a.GetToken(recurrer.TokenID)
	if err != nil {
		logger.Error().Err(err).
			Str("id", recurrer.TokenID.String()).
			Str("recurrer_id", recurrer.ID.String()).
			Msg("failed to retrieve token")
		return
	}

	rec := NewRecurrer(recurrer, a.tickRate, func(e entity.E) {
		raw, err := storage.NewEntity(e).Marshal(nil)
		if err != nil {
			logger.Error().Err(err).Msg("failed to retrieve marshal entity")
			return
		}
		logger.Info().Str("entity", e.ID.String()).Str("ip", token.IP.String()).Msg("send entity to ip")
		a.Send(raw, token.IP)
	})
	rec.EntityMapper = a.EntityMapper
	rec.EntitiesMapper = a.EntitiesMapper
	rec.SectorMapper = a.SectorMapper

	go rec.Start()
	a.recurrers[recurrer.ID] = rec
	logger.Info().
		Str("recurrer", recurrer.ID.String()).
		Str("entity", recurrer.EntityID.String()).
		Str("ip", token.IP.String()).
		Msg("synchronizing")
}
