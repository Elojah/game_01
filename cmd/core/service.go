package main

import (
	multierror "github.com/hashicorp/go-multierror"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/item"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
)

type service struct {
	ability   ability.App
	account   account.App
	entity    entity.App
	event     event.App
	item      item.App
	sector    sector.App
	sequencer infra.SequencerApp

	id ulid.ID

	subs map[ulid.ID]*infra.Subscription
	seqs map[ulid.ID]*Sequencer

	limit         int
	lootRadius    float64
	consumeRadius float64
}

func (svc *service) Dial(c Config) error {
	svc.id = c.ID
	svc.limit = c.Limit
	svc.lootRadius = c.LootRadius
	svc.consumeRadius = c.ConsumeRadius
	svc.sector.Dial(c.MoveTolerance)
	go svc.Run()
	return nil
}

func (svc *service) Run() {
	logger := log.With().Str("core", svc.id.String()).Logger()

	svc.subs = make(map[ulid.ID]*infra.Subscription)
	svc.subs[svc.id] = svc.sequencer.SubscribeSequencer(svc.id)
	go func(sub *infra.Subscription) {
		for msg := range sub.Channel() {
			go svc.Sequencer(msg)
		}
	}(svc.subs[svc.id])

	svc.seqs = make(map[ulid.ID]*Sequencer)

	if err := svc.sequencer.UpsertCore(infra.Core{ID: svc.id}); err != nil {
		logger.Error().Err(err).Msg("failed to set core")
		return
	}
}

func (svc *service) Close() error {
	var result *multierror.Error

	for _, s := range svc.subs {
		if err := s.Unsubscribe(); err != nil {
			result = multierror.Append(result, err)
		}
	}
	for _, s := range svc.seqs {
		if err := s.Close(); err != nil {
			result = multierror.Append(result, err)
		}
	}
	return result.ErrorOrNil()
}

func (svc *service) Sequencer(msg *infra.Message) {
	logger := log.With().Str("core", svc.id.String()).Logger()

	var sequencer infra.Sequencer
	if err := sequencer.Unmarshal([]byte(msg.Payload)); err != nil {
		logger.Error().Err(err).Msg("failed to unmarshal sequencer")
		return
	}
	logger = logger.With().Str("sequencer", sequencer.ID.String()).Logger()

	switch sequencer.Action {

	case infra.Open:
		svc.seqs[sequencer.ID] = NewSequencer(sequencer.ID, svc.limit, svc.Apply)
		svc.seqs[sequencer.ID].Event = svc.event
		svc.seqs[sequencer.ID].Entity = svc.entity
		svc.seqs[sequencer.ID].Run()

		svc.subs[sequencer.ID] = svc.event.Subscribe(sequencer.ID)

		go func(seq *Sequencer, sub *infra.Subscription) {
			for msg := range sub.Channel() {
				seq.Handler(msg)
			}
		}(svc.seqs[sequencer.ID], svc.subs[sequencer.ID])
		logger.Info().Msg("sequencer up")

	case infra.Close:
		seq, ok := svc.seqs[sequencer.ID]
		if !ok {
			logger.Error().Msg("sequencer not found")
			return
		}
		if err := seq.Close(); err != nil {
			logger.Error().Err(err).Msg("failed to close sequencer")
			return
		}
		sub, ok := svc.subs[sequencer.ID]
		if !ok {
			logger.Error().Str("subscription", sequencer.ID.String()).Msg("subscription not found")
			return
		}
		if err := sub.Unsubscribe(); err != nil {
			logger.Error().Err(err).Str("subscription", sequencer.ID.String()).Msg("failed to unsubscribe")
			return
		}
		delete(svc.seqs, sequencer.ID)
		delete(svc.subs, sequencer.ID)
		logger.Info().Msg("sequencer down")
	}
}

func (svc *service) Apply(id ulid.ID, e event.E) {
	logger := log.With().
		Str("core", svc.id.String()).
		Str("sequencer", id.String()).
		Str("event", e.ID.String()).
		Uint64("ts", e.ID.Time()).
		Str("action", e.Action.String()).
		Logger()

	var err error
	switch e.Action.GetValue().(type) {
	case *event.MoveTarget:
		err = svc.MoveTarget(id, e)
	case *event.CastSource:
		err = svc.CastSource(id, e)
	case *event.PerformSource:
		err = svc.PerformSource(id, e)
	case *event.PerformTarget:
		err = svc.PerformTarget(id, e)
	case *event.PerformFeedback:
		err = svc.PerformFeedback(id, e)
	case *event.LootSource:
		err = svc.LootSource(id, e)
	case *event.LootTarget:
		err = svc.LootTarget(id, e)
	case *event.LootFeedback:
		err = svc.LootFeedback(id, e)
	case *event.ConsumeSource:
		err = svc.ConsumeSource(id, e)
	case *event.ConsumeTarget:
		err = svc.ConsumeTarget(id, e)
	case *event.ConsumeFeedback:
		err = svc.ConsumeFeedback(id, e)
	case *event.Spawn:
		err = svc.Spawn(id, e)
	default:
		logger.Error().Msg("unrecognized action")
	}
	if err != nil {
		if gerrors.IsGameLogicError(err) {
			if err := svc.event.Cancel(e); err != nil {
				logger.Error().Err(err).Msg("cancel event")
			}
		}
		logger.Error().Err(err).Msg("apply event")
	}
	logger.Info().Msg("applied")
}
