package main

import (
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/storage"
	"github.com/elojah/game_01/pkg/ulid"
)

type app struct {
	ability.FeedbackMapper
	AbilityTemplateMapper ability.TemplateMapper
	AbilityMapper         ability.Mapper

	account.TokenMapper

	EntityTemplateMapper entity.TemplateMapper
	EntityMapper         entity.Mapper

	event.QListenerMapper
	event.QMapper
	EventMapper event.Mapper

	infra.CoreMapper

	entity.PermissionMapper

	sector.EntitiesMapper
	SectorMapper sector.Mapper

	id ulid.ID

	subs map[ulid.ID]*event.Subscription
	seqs map[ulid.ID]*Sequencer

	limit         int
	moveTolerance float64
}

func (a *app) Dial(c Config) error {
	a.id = c.ID
	a.limit = c.Limit
	a.moveTolerance = c.MoveTolerance
	go a.Start()
	return nil
}

func (a *app) Start() {
	logger := log.With().Str("core", a.id.String()).Logger()

	a.subs = make(map[ulid.ID]*event.Subscription)
	a.subs[a.id] = a.SubscribeListener(a.id)
	go func(sub *event.Subscription) {
		for msg := range sub.Channel() {
			go a.AddListener(msg)
		}
	}(a.subs[a.id])

	a.seqs = make(map[ulid.ID]*Sequencer)

	if err := a.SetCore(infra.Core{ID: a.id}); err != nil {
		logger.Error().Err(err).Msg("failed to set core")
		return
	}
}

func (a *app) Close() {
	for _, s := range a.subs {
		s.Unsubscribe()
	}
	for _, s := range a.seqs {
		s.Close()
	}
}

func (a *app) AddListener(msg *event.Message) {
	logger := log.With().Str("core", a.id.String()).Logger()

	var listenerS storage.Listener
	if _, err := listenerS.Unmarshal([]byte(msg.Payload)); err != nil {
		logger.Error().Err(err).Msg("failed to unmarshal listener")
		return
	}
	listener := listenerS.Domain()
	id := listener.ID.String()
	logger = logger.With().Str("listener", id).Uint8("action", uint8(listener.Action)).Logger()

	switch listener.Action {
	case event.Open:
		a.seqs[listener.ID] = NewSequencer(listener.ID, a.limit, a.EventMapper, a.Apply)
		a.seqs[listener.ID].Start()

		a.subs[listener.ID] = a.SubscribeEvent(listener.ID)

		go func(seq *Sequencer, sub *event.Subscription) {
			for msg := range sub.Channel() {
				seq.Handler(msg)
			}
		}(a.seqs[listener.ID], a.subs[listener.ID])

		logger.Info().Msg("listening")
	case event.Close:
		seq, ok := a.seqs[listener.ID]
		if !ok {
			logger.Error().Msg("listener not found")
			return
		}
		seq.Close()
		sub, ok := a.subs[listener.ID]
		if !ok {
			logger.Error().Msg("subscription not found")
			return
		}
		if err := sub.Unsubscribe(); err != nil {
			logger.Error().Err(err).Msg("failed to unsubscribe")
			return
		}
		delete(a.seqs, listener.ID)
		delete(a.subs, listener.ID)
	}
}

func (a *app) Apply(id ulid.ID, e event.E) {
	logger := log.With().
		Str("core", a.id.String()).
		Str("listener", id.String()).
		Int64("ts", e.TS.UnixNano()).
		Str("type", event.ActionString(e.Action)).
		Logger()

	switch e.Action.(type) {
	case event.Move:
		if err := a.Move(id, e); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case event.Cast:
		if err := a.Cast(id, e); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	default:
		logger.Error().Msg("unrecognized action")
	}
	logger.Info().Msg("event applied")
}
