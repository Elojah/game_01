package main

import (
	"os"
	"sync"
	"testing"
	"time"

	"github.com/oklog/ulid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/elojah/game_01/pkg/ability"
	entitymocks "github.com/elojah/game_01/pkg/entity/mocks"
	"github.com/elojah/game_01/pkg/event"
	eventmocks "github.com/elojah/game_01/pkg/event/mocks"
	"github.com/elojah/game_01/pkg/infra"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

func TestSequencer(t *testing.T) {

	now := ulid.Now()
	cid := gulid.NewID().String()
	eset := []event.E{
		event.E{
			ID: gulid.NewTimeID(now),
			Action: event.Action{
				CastSource: &event.CastSource{
					AbilityID: gulid.NewID(),
					Targets: map[string]ability.Targets{
						cid: ability.Targets{
							Entities: []gulid.ID{gulid.NewID(), gulid.NewID(), gulid.NewID()},
						},
					},
				},
			},
		},
		event.E{
			ID: gulid.NewTimeID(now + 1),
			Action: event.Action{
				MoveTarget: &event.MoveTarget{SourceID: gulid.NewID()},
			},
		},
		event.E{
			ID: gulid.NewTimeID(now + 2),
			Action: event.Action{
				MoveTarget: &event.MoveTarget{SourceID: gulid.NewID()},
			},
		},
		event.E{
			ID: gulid.NewTimeID(now + 3),
			Action: event.Action{
				MoveTarget: &event.MoveTarget{SourceID: gulid.NewID()},
			},
		},
	}

	t.Run("simple", func(t *testing.T) {

		seqID := gulid.NewID()
		entityStore := entitymocks.NewStore()
		eventStore := eventmocks.NewStore()
		eventTriggerService := eventmocks.NewTriggerService()
		eventStore.ListEventFunc = func(key gulid.ID, min gulid.ID) ([]event.E, error) {
			assert.Equal(t, seqID.String(), key.String())
			switch eventStore.ListEventCount {
			case 0:
				assert.Equal(t, eset[0].ID.String(), min.String())
			}
			return []event.E{eset[0]}, nil
		}

		var wg sync.WaitGroup
		wg.Add(1)
		seq := NewSequencer(seqID, 32,
			func(id gulid.ID, e event.E) {
				assert.True(t, eset[0].Equal(e))
				wg.Done()
			},
		)
		seq.EventStore = eventStore
		seq.EventTriggerService = eventTriggerService
		seq.EntityStore = entityStore
		seq.logger = zerolog.Nop()
		seq.Run()

		raw, err := eset[0].Marshal()
		assert.NoError(t, err)
		msg := &infra.Message{Payload: string(raw)}
		seq.Handler(msg)
		wg.Wait()
		seq.Close()
	})

	t.Run("two", func(t *testing.T) {

		seqID := gulid.NewID()
		entityStore := entitymocks.NewStore()
		eventStore := eventmocks.NewStore()
		eventTriggerService := eventmocks.NewTriggerService()
		eventStore.ListEventFunc = func(key gulid.ID, min gulid.ID) ([]event.E, error) {
			assert.Equal(t, seqID.String(), key.String())
			switch min.String() {
			case eset[0].ID.String():
				return []event.E{eset[0]}, nil
			case eset[1].ID.String():
				return []event.E{eset[1]}, nil
			}
			return nil, nil
		}

		var wg sync.WaitGroup
		wg.Add(2)
		seq := NewSequencer(seqID, 32,
			func(id gulid.ID, e event.E) {
				wg.Done()
			},
		)
		seq.EventStore = eventStore
		seq.EventTriggerService = eventTriggerService
		seq.EntityStore = entityStore
		seq.logger = zerolog.Nop()
		seq.Run()

		raw0, err := eset[0].Marshal()
		assert.NoError(t, err)
		msg0 := &infra.Message{Payload: string(raw0)}

		raw1, err := eset[1].Marshal()
		assert.NoError(t, err)
		msg1 := &infra.Message{Payload: string(raw1)}

		seq.Handler(msg0)
		seq.Handler(msg1)

		wg.Wait()

		seq.Close()
	})

	t.Run("cancel", func(t *testing.T) {

		seqID := gulid.NewID()
		entityStore := entitymocks.NewStore()
		eventStore := eventmocks.NewStore()
		eventTriggerService := eventmocks.NewTriggerService()
		eventStore.ListEventFunc = func(key gulid.ID, min gulid.ID) ([]event.E, error) {
			assert.Equal(t, seqID.String(), key.String())
			switch min.String() {
			case eset[0].ID.String():
				return []event.E{eset[0], eset[1]}, nil
			case eset[2].ID.String():
				n := 42
				es := make([]event.E, n)
				for i := 0; i < n; i++ {
					es[i] = eset[2]
				}
				return append(es, eset[3]), nil
			}
			return nil, nil
		}

		var wg sync.WaitGroup
		wg.Add(1)
		seq := NewSequencer(seqID, 32,
			func(id gulid.ID, e event.E) {
				assert.False(t, e.Equal(eset[3]))
				if e.Equal(eset[1]) {
					wg.Done()
				}
			},
		)
		seq.EventStore = eventStore
		seq.EventTriggerService = eventTriggerService
		seq.EntityStore = entityStore
		seq.logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
		seq.logger = zerolog.Nop()
		seq.Run()

		raw2, err := eset[2].Marshal()
		assert.NoError(t, err)
		msg2 := &infra.Message{Payload: string(raw2)}

		raw0, err := eset[0].Marshal()
		assert.NoError(t, err)
		msg0 := &infra.Message{Payload: string(raw0)}

		seq.Handler(msg2)
		seq.Handler(msg0)

		wg.Wait()

		seq.Close()
	})

	t.Run("interrupt", func(t *testing.T) {

		seqID := gulid.NewID()
		entityStore := entitymocks.NewStore()
		eventStore := eventmocks.NewStore()
		eventTriggerService := eventmocks.NewTriggerService()
		eventStore.ListEventFunc = func(key gulid.ID, min gulid.ID) ([]event.E, error) {
			assert.Equal(t, seqID.String(), key.String())
			switch min.String() {
			case eset[0].ID.String():
				return []event.E{eset[0], eset[1]}, nil
			case eset[2].ID.String():
				n := 4242
				es := make([]event.E, n)
				for i := 0; i < n; i++ {
					es[i] = eset[2]
				}
				return append(es, eset[3]), nil
			}
			return nil, nil
		}

		var wg sync.WaitGroup
		wg.Add(1)
		seq := NewSequencer(seqID, 1,
			func(id gulid.ID, e event.E) {
				if e.Equal(eset[2]) {
					time.Sleep(5 * time.Millisecond)
				}
				assert.False(t, e.Equal(eset[3]))
				if e.Equal(eset[1]) {
					wg.Done()
				}
			},
		)
		seq.EventStore = eventStore
		seq.EventTriggerService = eventTriggerService
		seq.EntityStore = entityStore
		seq.logger = zerolog.Nop()
		seq.Run()

		raw2, err := eset[2].Marshal()
		assert.NoError(t, err)
		msg2 := &infra.Message{Payload: string(raw2)}

		raw0, err := eset[0].Marshal()
		assert.NoError(t, err)
		msg0 := &infra.Message{Payload: string(raw0)}

		seq.Handler(msg2)
		time.Sleep(10 * time.Millisecond)
		seq.Handler(msg0)

		wg.Wait()

		seq.Close()
	})

}
