package main

import (
	"net"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/mux/client"
)

type service struct {
	account.TokenStore

	EntityStore entity.Store

	event.QStore
	infra.QRecurrerStore

	infra.SyncStore

	sector.EntitiesStore
	SectorStore sector.Store

	*client.C

	id ulid.ID

	sub *infra.Subscription

	port      uint
	tickRate  uint32
	batchSize uint32
	recurrers map[ulid.ID]*Recurrer
}

func (svc *service) Dial(c Config) error {
	svc.port = c.EntityPort
	svc.tickRate = c.TickRate
	svc.batchSize = c.BatchSize
	svc.id = c.ID
	go svc.Run()
	return nil
}

func (svc *service) Run() {
	logger := log.With().Str("sync", svc.id.String()).Logger()

	svc.sub = svc.SubscribeRecurrer(svc.id)
	go func() {
		for msg := range svc.sub.Channel() {
			go svc.Recurrer(msg)
		}
	}()

	svc.recurrers = make(map[ulid.ID]*Recurrer)

	if err := svc.SetSync(infra.Sync{ID: svc.id}); err != nil {
		logger.Error().Err(err).Msg("failed to set sync")
		return
	}
}

func (svc *service) Close() error {
	var result *multierror.Error

	if err := svc.sub.Unsubscribe(); err != nil {
		return err
	}
	for _, r := range svc.recurrers {
		if r != nil {
			if err := r.Close(); err != nil {
				result = multierror.Append(result, err)
			}
		}
	}
	return result.ErrorOrNil()
}

func (svc *service) Recurrer(msg *infra.Message) {
	logger := log.With().Str("sync", svc.id.String()).Logger()

	var r infra.Recurrer
	if err := r.Unmarshal([]byte(msg.Payload)); err != nil {
		logger.Error().Err(err).Msg("failed to unmarshal recurrer")
		return
	}

	logger = logger.With().Str("recurrer", r.TokenID.String()).Logger()

	if r.Action == infra.Close {
		rec := svc.recurrers[r.TokenID]
		if rec != nil {
			if err := rec.Close(); err != nil {
				logger.Error().Err(err).Msg("failed to close recurrer")
				return
			}
		}
		delete(svc.recurrers, r.TokenID)
		return
	}

	tok, err := svc.GetToken(r.TokenID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve token")
		return
	}

	addr, err := net.ResolveUDPAddr("udp", tok.IP)
	if err != nil {
		logger.Error().Err(err).Msg("failed to parse ip")
		return
	}
	addr.Port = int(svc.port)
	logger = logger.With().Str("address", addr.String()).Logger()
	rec := NewRecurrer(r, svc.tickRate, svc.batchSize, func(dto entity.DTO) {
		raw, err := dto.Marshal()
		if err != nil {
			logger.Error().Err(err).Msg("failed to marshal entity")
			return
		}
		svc.Send(raw, addr)
	})
	rec.EntityStore = svc.EntityStore
	rec.SectorEntitiesStore = svc.EntitiesStore
	rec.SectorStore = svc.SectorStore

	go rec.Run()
	svc.recurrers[r.TokenID] = rec
	logger.Info().Msg("recurrer up")
}
