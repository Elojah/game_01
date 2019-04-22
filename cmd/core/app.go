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

type app struct {
	AbilityStore         ability.Store
	AbilityTemplateStore ability.TemplateStore
	AbilityFeedbackStore ability.FeedbackStore

	account.TokenStore

	EntityPermissionService entity.PermissionService
	EntityPermissionStore   entity.PermissionStore
	EntitySpawnStore        entity.SpawnStore
	EntityStore             entity.Store
	EntityTemplateStore     entity.TemplateStore
	EntityInventoryService  entity.InventoryService

	EventQStore         event.QStore
	EventStore          event.Store
	EventTriggerService event.TriggerService

	infra.CoreStore
	infra.QSequencerStore

	ItemStore     item.Store
	ItemLootStore item.LootStore

	sector.EntitiesStore
	SectorStore   sector.Store
	SectorService sector.Service

	id ulid.ID

	subs map[ulid.ID]*infra.Subscription
	seqs map[ulid.ID]*Sequencer

	limit         int
	lootRadius    float64
	consumeRadius float64
}

func (a *app) Dial(c Config) error {
	a.id = c.ID
	a.limit = c.Limit
	a.lootRadius = c.LootRadius
	a.consumeRadius = c.ConsumeRadius
	if err := a.SectorService.Up(c.MoveTolerance); err != nil {
		return err
	}
	go a.Run()
	return nil
}

func (a *app) Run() {
	logger := log.With().Str("core", a.id.String()).Logger()

	a.subs = make(map[ulid.ID]*infra.Subscription)
	a.subs[a.id] = a.SubscribeSequencer(a.id)
	go func(sub *infra.Subscription) {
		for msg := range sub.Channel() {
			go a.Sequencer(msg)
		}
	}(a.subs[a.id])

	a.seqs = make(map[ulid.ID]*Sequencer)

	if err := a.SetCore(infra.Core{ID: a.id}); err != nil {
		logger.Error().Err(err).Msg("failed to set core")
		return
	}
}

func (a *app) Close() error {
	var result *multierror.Error

	for _, s := range a.subs {
		if err := s.Unsubscribe(); err != nil {
			result = multierror.Append(result, err)
		}
	}
	for _, s := range a.seqs {
		if err := s.Close(); err != nil {
			result = multierror.Append(result, err)
		}
	}
	return result.ErrorOrNil()
}

func (a *app) Sequencer(msg *infra.Message) {
	logger := log.With().Str("core", a.id.String()).Logger()

	var sequencer infra.Sequencer
	if err := sequencer.Unmarshal([]byte(msg.Payload)); err != nil {
		logger.Error().Err(err).Msg("failed to unmarshal sequencer")
		return
	}
	logger = logger.With().Str("sequencer", sequencer.ID.String()).Logger()

	switch sequencer.Action {

	case infra.Open:
		a.seqs[sequencer.ID] = NewSequencer(sequencer.ID, a.limit, a.Apply)
		a.seqs[sequencer.ID].EventStore = a.EventStore
		a.seqs[sequencer.ID].EventTriggerService = a.EventTriggerService
		a.seqs[sequencer.ID].EntityStore = a.EntityStore
		a.seqs[sequencer.ID].Run()

		a.subs[sequencer.ID] = a.EventQStore.SubscribeEvent(sequencer.ID)

		go func(seq *Sequencer, sub *infra.Subscription) {
			for msg := range sub.Channel() {
				seq.Handler(msg)
			}
		}(a.seqs[sequencer.ID], a.subs[sequencer.ID])
		logger.Info().Msg("sequencer up")

	case infra.Close:
		seq, ok := a.seqs[sequencer.ID]
		if !ok {
			logger.Error().Msg("sequencer not found")
			return
		}
		if err := seq.Close(); err != nil {
			logger.Error().Err(err).Msg("failed to close sequencer")
			return
		}
		sub, ok := a.subs[sequencer.ID]
		if !ok {
			logger.Error().Str("subscription", sequencer.ID.String()).Msg("subscription not found")
			return
		}
		if err := sub.Unsubscribe(); err != nil {
			logger.Error().Err(err).Str("subscription", sequencer.ID.String()).Msg("failed to unsubscribe")
			return
		}
		delete(a.seqs, sequencer.ID)
		delete(a.subs, sequencer.ID)
		logger.Info().Msg("sequencer down")
	}
}

func (a *app) Apply(id ulid.ID, e event.E) {
	logger := log.With().
		Str("core", a.id.String()).
		Str("sequencer", id.String()).
		Str("event", e.ID.String()).
		Uint64("ts", id.Time()).
		Str("action", e.Action.String()).
		Logger()

	var err error
	switch e.Action.GetValue().(type) {
	case *event.MoveTarget:
		err = a.MoveTarget(id, e)
	case *event.CastSource:
		err = a.CastSource(id, e)
	case *event.PerformSource:
		err = a.PerformSource(id, e)
	case *event.PerformTarget:
		err = a.PerformTarget(id, e)
	case *event.PerformFeedback:
		err = a.PerformFeedback(id, e)
	case *event.LootSource:
		err = a.LootSource(id, e)
	case *event.LootTarget:
		err = a.LootTarget(id, e)
	case *event.LootFeedback:
		err = a.LootFeedback(id, e)
	case *event.ConsumeSource:
		err = a.ConsumeSource(id, e)
	case *event.ConsumeTarget:
		err = a.ConsumeTarget(id, e)
	case *event.ConsumeFeedback:
		err = a.ConsumeFeedback(id, e)
	case *event.Spawn:
		err = a.Spawn(id, e)
	default:
		logger.Error().Msg("unrecognized action")
	}
	if err != nil {
		if gerrors.IsGameLogicError(err) {
			if err := a.EventTriggerService.Cancel(e); err != nil {
				logger.Error().Err(err).Msg("cancel event")
			}
		}
		logger.Error().Err(err).Msg("apply event")
	}
	logger.Info().Msg("applied")
}
