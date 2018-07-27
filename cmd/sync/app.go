package main

import (
	"net"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/mux/client"
)

type app struct {
	account.TokenMapper

	EntityMapper entity.Mapper

	event.QMapper
	infra.QRecurrerMapper

	infra.SyncMapper

	sector.EntitiesMapper
	SectorMapper sector.Mapper

	*client.C

	id ulid.ID

	sub *infra.Subscription

	port      uint
	tickRate  uint32
	recurrers map[ulid.ID]*Recurrer
}

func (a *app) Dial(c Config) error {
	a.port = c.EntityPort
	a.tickRate = c.TickRate
	a.id = c.ID
	go a.Run()
	return nil
}

func (a *app) Run() {
	logger := log.With().Str("sync", a.id.String()).Logger()

	a.sub = a.SubscribeRecurrer(a.id)
	go func() {
		for msg := range a.sub.Channel() {
			go a.AddRecurrer(msg)
		}
	}()

	a.recurrers = make(map[ulid.ID]*Recurrer)

	if err := a.SetSync(infra.Sync{ID: a.id}); err != nil {
		logger.Error().Err(err).Msg("failed to set sync")
		return
	}
}

func (a *app) Close() {
	a.sub.Unsubscribe()
	for _, r := range a.recurrers {
		if r != nil {
			r.Close()
		}
	}
}

func (a *app) AddRecurrer(msg *infra.Message) {
	logger := log.With().Str("sync", a.id.String()).Logger()

	var recurrer infra.Recurrer
	if _, err := recurrer.Unmarshal([]byte(msg.Payload)); err != nil {
		logger.Error().Err(err).Msg("failed to unmarshal recurrer")
		return
	}

	logger = logger.With().Str("recurrer", recurrer.TokenID.String()).Logger()

	if recurrer.Action == infra.Close {
		rec := a.recurrers[recurrer.TokenID]
		if rec != nil {
			rec.Close()
		}
		delete(a.recurrers, recurrer.TokenID)
		return
	}

	tok, err := a.GetToken(account.TokenSubset{ID: recurrer.TokenID})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve token")
		return
	}

	address, err := net.ResolveUDPAddr("udp", tok.IP)
	if err != nil {
		logger.Error().Err(err).Msg("failed to parse ip")
		return
	}
	address.Port = int(a.port)
	logger = logger.With().Str("address", address.String()).Logger()
	rec := NewRecurrer(recurrer, a.tickRate, func(e entity.E) {
		raw, err := e.Marshal(nil)
		if err != nil {
			logger.Error().Err(err).Msg("failed to marshal entity")
			return
		}
		logger.Info().Str("entity", e.ID.String()).Msg("send entity")
		a.Send(raw, address)
	})
	rec.EntityMapper = a.EntityMapper
	rec.EntitiesMapper = a.EntitiesMapper
	rec.SectorMapper = a.SectorMapper

	go rec.Run()
	a.recurrers[recurrer.TokenID] = rec
	logger.Info().Msg("sync up")
}
